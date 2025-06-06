// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type ProcedureJavaResourceAssert struct {
	*assert.ResourceAssert
}

func ProcedureJavaResource(t *testing.T, name string) *ProcedureJavaResourceAssert {
	t.Helper()

	return &ProcedureJavaResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedProcedureJavaResource(t *testing.T, id string) *ProcedureJavaResourceAssert {
	t.Helper()

	return &ProcedureJavaResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (p *ProcedureJavaResourceAssert) HasDatabaseString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("database", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasSchemaString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("schema", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNameString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("name", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasArgumentsString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("arguments", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasCommentString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("comment", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasEnableConsoleOutputString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("enable_console_output", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasExecuteAsString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("execute_as", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasExternalAccessIntegrationsString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("external_access_integrations", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasFullyQualifiedNameString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasHandlerString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("handler", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasImportsString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("imports", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasIsSecureString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("is_secure", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasLogLevelString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("log_level", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasMetricLevelString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("metric_level", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNullInputBehaviorString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("null_input_behavior", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasPackagesString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("packages", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasProcedureDefinitionString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_definition", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasProcedureLanguageString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_language", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasReturnTypeString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("return_type", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasRuntimeVersionString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("runtime_version", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasSecretsString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("secrets", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasSnowparkPackageString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("snowpark_package", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasTargetPathString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("target_path", expected))
	return p
}

func (p *ProcedureJavaResourceAssert) HasTraceLevelString(expected string) *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("trace_level", expected))
	return p
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (p *ProcedureJavaResourceAssert) HasNoDatabase() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("database"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoSchema() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("schema"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoName() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("name"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoComment() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("comment"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoEnableConsoleOutput() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("enable_console_output"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoExecuteAs() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("execute_as"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoFullyQualifiedName() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoHandler() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("handler"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoIsSecure() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("is_secure"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoLogLevel() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("log_level"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoMetricLevel() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("metric_level"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoNullInputBehavior() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("null_input_behavior"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoProcedureDefinition() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("procedure_definition"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoProcedureLanguage() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("procedure_language"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoReturnType() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("return_type"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoRuntimeVersion() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("runtime_version"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoSnowparkPackage() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("snowpark_package"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNoTraceLevel() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueNotSet("trace_level"))
	return p
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (p *ProcedureJavaResourceAssert) HasArgumentsEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("arguments.#", "0"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasCommentEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("comment", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasEnableConsoleOutputEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("enable_console_output", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasExecuteAsEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("execute_as", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasExternalAccessIntegrationsEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("external_access_integrations.#", "0"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasFullyQualifiedNameEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasImportsEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("imports.#", "0"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasIsSecureEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("is_secure", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasLogLevelEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("log_level", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasMetricLevelEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("metric_level", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNullInputBehaviorEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("null_input_behavior", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasPackagesEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("packages.#", "0"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasProcedureDefinitionEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_definition", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasProcedureLanguageEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("procedure_language", ""))
	return p
}

func (p *ProcedureJavaResourceAssert) HasSecretsEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("secrets.#", "0"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasTargetPathEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("target_path.#", "0"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasTraceLevelEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValueSet("trace_level", ""))
	return p
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (p *ProcedureJavaResourceAssert) HasDatabaseNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("database"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasSchemaNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("schema"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNameNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("name"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasCommentNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("comment"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasEnableConsoleOutputNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("enable_console_output"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasExecuteAsNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("execute_as"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasFullyQualifiedNameNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasHandlerNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("handler"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasIsSecureNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("is_secure"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasLogLevelNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("log_level"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasMetricLevelNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("metric_level"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasNullInputBehaviorNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("null_input_behavior"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasProcedureDefinitionNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("procedure_definition"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasProcedureLanguageNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("procedure_language"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasReturnTypeNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("return_type"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasRuntimeVersionNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("runtime_version"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasSnowparkPackageNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("snowpark_package"))
	return p
}

func (p *ProcedureJavaResourceAssert) HasTraceLevelNotEmpty() *ProcedureJavaResourceAssert {
	p.AddAssertion(assert.ValuePresent("trace_level"))
	return p
}
