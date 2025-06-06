//go:build exclude

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func main() {
	filename := os.Getenv("GOFILE")
	fmt.Printf("Running generator on %s with args %#v\n", filename, os.Args[1:])

	astFile := parseSourceFile(filename)

	gen := setUpGenerator(astFile)
	fmt.Printf("Generator set up for package \"%s\" with output name \"%s\".\n", gen.outputPackage, gen.outputName)

	gen.addFilePreamble()
	gen.addImports()
	gen.addConstructorsAndBuilderMethods()

	src, errSrcFormat := format.Source(gen.buffer.Bytes())
	if errSrcFormat != nil {
		log.Panicln(errSrcFormat)
	}
	if err := os.WriteFile(gen.outputName, src, 0o600); err != nil {
		log.Panicln(err)
	}
}

func parseSourceFile(filename string) *ast.File {
	astFile, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return astFile
}

type Generator struct {
	astFile       *ast.File
	buffer        bytes.Buffer
	outputPackage string
	outputName    string
}

func setUpGenerator(astFile *ast.File) *Generator {
	wd, errWd := os.Getwd()
	if errWd != nil {
		log.Panicln(errWd)
	}

	file := os.Getenv("GOFILE")
	fileWithoutSuffix, _ := strings.CutSuffix(file, ".go")
	fileWithoutSuffix, _ = strings.CutSuffix(fileWithoutSuffix, "_gen")
	baseName := fmt.Sprintf("%s_builders_gen.go", fileWithoutSuffix)
	outputName := filepath.Join(wd, baseName)

	return &Generator{
		astFile:       astFile,
		buffer:        bytes.Buffer{},
		outputPackage: os.Getenv("GOPACKAGE"),
		outputName:    outputName,
	}
}

func (gen *Generator) printf(format string, args ...any) {
	printf(&gen.buffer, format, args...)
}

func printf(w io.Writer, format string, args ...any) {
	_, err := fmt.Fprintf(w, format, args...)
	if err != nil {
		log.Panicln(err)
	}
}

func (gen *Generator) addFilePreamble() {
	gen.printf("// Code generated by dto builder generator; DO NOT EDIT.\n")
	gen.printf("\n")
	gen.printf("package %s\n", gen.outputPackage)
	gen.printf("\n")
}

func (gen *Generator) addImports() {
	imports := make([]string, 0, len(gen.astFile.Imports))
	for _, importNode := range gen.astFile.Imports {
		imports = append(imports, types.ExprString(importNode.Path))
	}
	gen.generateImports(imports...)
}

func (gen *Generator) generateImports(imports ...string) {
	gen.printf("import (\n")
	for _, i := range imports {
		gen.printf("%s\n", i)
	}
	gen.printf(")\n")
	gen.printf("\n")
}

func (gen *Generator) addConstructorsAndBuilderMethods() {
	for _, node := range gen.astFile.Decls {
		genDecl, ok := node.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			fields := extractFieldDefs(structType.Fields.List)
			def := newStructDef(typeSpec, fields)
			gen.generateConstructor(def)
			gen.generateBuilderMethods(def)
		}
	}
}

func extractFieldDefs(fieldList []*ast.Field) []*fieldDef {
	var fields []*fieldDef
	for _, field := range fieldList {
		for _, name := range field.Names {
			fs := newFieldDef(name, field)
			fields = append(fields, fs)
		}
	}
	return fields
}

type structDef struct {
	name   string
	fields []*fieldDef
}

func newStructDef(ts *ast.TypeSpec, fields []*fieldDef) *structDef {
	return &structDef{
		name:   ts.Name.Name,
		fields: fields,
	}
}

type fieldDef struct {
	name       string
	typeString string
	isRequired bool
}

func (fs *fieldDef) String() string {
	return fmt.Sprintf("Field: name=%s type=%s is required=%t", fs.name, fs.typeString, fs.isRequired)
}

func newFieldDef(name *ast.Ident, field *ast.Field) *fieldDef {
	return &fieldDef{
		name:       name.Name,
		typeString: types.ExprString(field.Type),
		isRequired: strings.TrimSpace(field.Comment.Text()) == "required",
	}
}

func (gen *Generator) generateConstructor(d *structDef) {
	gen.printf("func New%s(", d.name)
	var requiredFields []*fieldDef
	for _, field := range d.fields {
		if field.isRequired {
			requiredFields = append(requiredFields, field)
		}
	}
	if len(requiredFields) != 0 {
		gen.printf("\n")
		for _, field := range requiredFields {
			gen.printf("%s %s,\n", field.name, field.typeString)
		}
	}

	gen.printf(") *%s {\n", d.name)

	var returnStatement string
	if len(requiredFields) != 0 {
		gen.printf("s := %s{}\n", d.name)
		for _, field := range requiredFields {
			gen.printf("s.%s = %s\n", field.name, field.name)
		}
		returnStatement = "&s"
	} else {
		returnStatement = fmt.Sprintf("&%s{}", d.name)
	}

	gen.printf("return %s\n", returnStatement)
	gen.printf("}\n\n")
}

func (gen *Generator) generateBuilderMethods(d *structDef) {
	var optionalFields []*fieldDef
	for _, field := range d.fields {
		if !field.isRequired {
			optionalFields = append(optionalFields, field)
		}
	}

	for _, field := range optionalFields {
		gen.printf("func (s *%s) With%s(%s %s) *%s {\n", d.name, toTitle(field.name), field.name, strings.TrimLeft(field.typeString, "*"), d.name)

		switch {
		case strings.HasPrefix(field.typeString, "*"):
			// If the target field is a pointer, assign the address of input field because right now we always pass them by value
			gen.printf("s.%s = &%s\n", field.name, field.name)
		default:
			gen.printf("s.%s = %s\n", field.name, field.name)
		}
		gen.printf("return s\n")
		gen.printf("}\n\n")
	}
}

func toTitle(s string) string {
	firstLetter, _ := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(firstLetter)) + s[1:]
}
