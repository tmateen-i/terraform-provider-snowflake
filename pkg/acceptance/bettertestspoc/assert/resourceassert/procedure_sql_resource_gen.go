// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type ProcedureSqlResourceAssert struct {
	*assert.ResourceAssert
}

func ProcedureSqlResource(t *testing.T, name string) *ProcedureSqlResourceAssert {
	t.Helper()

	return &ProcedureSqlResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedProcedureSqlResource(t *testing.T, id string) *ProcedureSqlResourceAssert {
	t.Helper()

	return &ProcedureSqlResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (p *ProcedureSqlResourceAssert) HasDatabaseString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("database", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasSchemaString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("schema", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNameString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("name", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasArgumentsString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("arguments", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasCommentString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("comment", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasEnableConsoleOutputString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("enable_console_output", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasExecuteAsString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("execute_as", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasFullyQualifiedNameString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasIsSecureString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("is_secure", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasLogLevelString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("log_level", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasMetricLevelString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("metric_level", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNullInputBehaviorString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("null_input_behavior", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasProcedureDefinitionString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_definition", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasProcedureLanguageString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_language", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasReturnTypeString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("return_type", expected))
	return p
}

func (p *ProcedureSqlResourceAssert) HasTraceLevelString(expected string) *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("trace_level", expected))
	return p
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (p *ProcedureSqlResourceAssert) HasNoDatabase() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("database"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoSchema() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("schema"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoName() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("name"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoComment() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("comment"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoEnableConsoleOutput() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("enable_console_output"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoExecuteAs() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("execute_as"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoFullyQualifiedName() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoIsSecure() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("is_secure"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoLogLevel() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("log_level"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoMetricLevel() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("metric_level"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoNullInputBehavior() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("null_input_behavior"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoProcedureDefinition() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("procedure_definition"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoProcedureLanguage() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("procedure_language"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoReturnType() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("return_type"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNoTraceLevel() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueNotSet("trace_level"))
	return p
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (p *ProcedureSqlResourceAssert) HasArgumentsEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("arguments.#", "0"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasCommentEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("comment", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasEnableConsoleOutputEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("enable_console_output", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasExecuteAsEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("execute_as", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasFullyQualifiedNameEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasIsSecureEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("is_secure", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasLogLevelEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("log_level", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasMetricLevelEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("metric_level", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNullInputBehaviorEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("null_input_behavior", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasProcedureLanguageEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_language", ""))
	return p
}

func (p *ProcedureSqlResourceAssert) HasTraceLevelEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValueSet("trace_level", ""))
	return p
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (p *ProcedureSqlResourceAssert) HasDatabaseNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("database"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasSchemaNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("schema"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNameNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("name"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasCommentNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("comment"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasEnableConsoleOutputNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("enable_console_output"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasExecuteAsNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("execute_as"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasFullyQualifiedNameNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasIsSecureNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("is_secure"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasLogLevelNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("log_level"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasMetricLevelNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("metric_level"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasNullInputBehaviorNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("null_input_behavior"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasProcedureDefinitionNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("procedure_definition"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasProcedureLanguageNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("procedure_language"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasReturnTypeNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("return_type"))
	return p
}

func (p *ProcedureSqlResourceAssert) HasTraceLevelNotEmpty() *ProcedureSqlResourceAssert {
	p.AddAssertion(assert.ValuePresent("trace_level"))
	return p
}
