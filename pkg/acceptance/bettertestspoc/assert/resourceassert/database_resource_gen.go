// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type DatabaseResourceAssert struct {
	*assert.ResourceAssert
}

func DatabaseResource(t *testing.T, name string) *DatabaseResourceAssert {
	t.Helper()

	return &DatabaseResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedDatabaseResource(t *testing.T, id string) *DatabaseResourceAssert {
	t.Helper()

	return &DatabaseResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (d *DatabaseResourceAssert) HasNameString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("name", expected))
	return d
}

func (d *DatabaseResourceAssert) HasCatalogString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("catalog", expected))
	return d
}

func (d *DatabaseResourceAssert) HasCommentString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("comment", expected))
	return d
}

func (d *DatabaseResourceAssert) HasDataRetentionTimeInDaysString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("data_retention_time_in_days", expected))
	return d
}

func (d *DatabaseResourceAssert) HasDefaultDdlCollationString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("default_ddl_collation", expected))
	return d
}

func (d *DatabaseResourceAssert) HasDropPublicSchemaOnCreationString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("drop_public_schema_on_creation", expected))
	return d
}

func (d *DatabaseResourceAssert) HasEnableConsoleOutputString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("enable_console_output", expected))
	return d
}

func (d *DatabaseResourceAssert) HasExternalVolumeString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("external_volume", expected))
	return d
}

func (d *DatabaseResourceAssert) HasFullyQualifiedNameString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return d
}

func (d *DatabaseResourceAssert) HasIsTransientString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("is_transient", expected))
	return d
}

func (d *DatabaseResourceAssert) HasLogLevelString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("log_level", expected))
	return d
}

func (d *DatabaseResourceAssert) HasMaxDataExtensionTimeInDaysString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("max_data_extension_time_in_days", expected))
	return d
}

func (d *DatabaseResourceAssert) HasQuotedIdentifiersIgnoreCaseString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("quoted_identifiers_ignore_case", expected))
	return d
}

func (d *DatabaseResourceAssert) HasReplaceInvalidCharactersString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("replace_invalid_characters", expected))
	return d
}

func (d *DatabaseResourceAssert) HasReplicationString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("replication", expected))
	return d
}

func (d *DatabaseResourceAssert) HasStorageSerializationPolicyString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("storage_serialization_policy", expected))
	return d
}

func (d *DatabaseResourceAssert) HasSuspendTaskAfterNumFailuresString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("suspend_task_after_num_failures", expected))
	return d
}

func (d *DatabaseResourceAssert) HasTaskAutoRetryAttemptsString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("task_auto_retry_attempts", expected))
	return d
}

func (d *DatabaseResourceAssert) HasTraceLevelString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("trace_level", expected))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskManagedInitialWarehouseSizeString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("user_task_managed_initial_warehouse_size", expected))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskMinimumTriggerIntervalInSecondsString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("user_task_minimum_trigger_interval_in_seconds", expected))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskTimeoutMsString(expected string) *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("user_task_timeout_ms", expected))
	return d
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (d *DatabaseResourceAssert) HasNoName() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("name"))
	return d
}

func (d *DatabaseResourceAssert) HasNoCatalog() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("catalog"))
	return d
}

func (d *DatabaseResourceAssert) HasNoComment() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("comment"))
	return d
}

func (d *DatabaseResourceAssert) HasNoDataRetentionTimeInDays() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("data_retention_time_in_days"))
	return d
}

func (d *DatabaseResourceAssert) HasNoDefaultDdlCollation() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("default_ddl_collation"))
	return d
}

func (d *DatabaseResourceAssert) HasNoDropPublicSchemaOnCreation() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("drop_public_schema_on_creation"))
	return d
}

func (d *DatabaseResourceAssert) HasNoEnableConsoleOutput() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("enable_console_output"))
	return d
}

func (d *DatabaseResourceAssert) HasNoExternalVolume() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("external_volume"))
	return d
}

func (d *DatabaseResourceAssert) HasNoFullyQualifiedName() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return d
}

func (d *DatabaseResourceAssert) HasNoIsTransient() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("is_transient"))
	return d
}

func (d *DatabaseResourceAssert) HasNoLogLevel() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("log_level"))
	return d
}

func (d *DatabaseResourceAssert) HasNoMaxDataExtensionTimeInDays() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("max_data_extension_time_in_days"))
	return d
}

func (d *DatabaseResourceAssert) HasNoQuotedIdentifiersIgnoreCase() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("quoted_identifiers_ignore_case"))
	return d
}

func (d *DatabaseResourceAssert) HasNoReplaceInvalidCharacters() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("replace_invalid_characters"))
	return d
}

func (d *DatabaseResourceAssert) HasNoStorageSerializationPolicy() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("storage_serialization_policy"))
	return d
}

func (d *DatabaseResourceAssert) HasNoSuspendTaskAfterNumFailures() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("suspend_task_after_num_failures"))
	return d
}

func (d *DatabaseResourceAssert) HasNoTaskAutoRetryAttempts() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("task_auto_retry_attempts"))
	return d
}

func (d *DatabaseResourceAssert) HasNoTraceLevel() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("trace_level"))
	return d
}

func (d *DatabaseResourceAssert) HasNoUserTaskManagedInitialWarehouseSize() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("user_task_managed_initial_warehouse_size"))
	return d
}

func (d *DatabaseResourceAssert) HasNoUserTaskMinimumTriggerIntervalInSeconds() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("user_task_minimum_trigger_interval_in_seconds"))
	return d
}

func (d *DatabaseResourceAssert) HasNoUserTaskTimeoutMs() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueNotSet("user_task_timeout_ms"))
	return d
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (d *DatabaseResourceAssert) HasCatalogEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("catalog", ""))
	return d
}

func (d *DatabaseResourceAssert) HasCommentEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("comment", ""))
	return d
}

func (d *DatabaseResourceAssert) HasDataRetentionTimeInDaysEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("data_retention_time_in_days", ""))
	return d
}

func (d *DatabaseResourceAssert) HasDefaultDdlCollationEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("default_ddl_collation", ""))
	return d
}

func (d *DatabaseResourceAssert) HasDropPublicSchemaOnCreationEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("drop_public_schema_on_creation", ""))
	return d
}

func (d *DatabaseResourceAssert) HasEnableConsoleOutputEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("enable_console_output", ""))
	return d
}

func (d *DatabaseResourceAssert) HasExternalVolumeEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("external_volume", ""))
	return d
}

func (d *DatabaseResourceAssert) HasFullyQualifiedNameEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return d
}

func (d *DatabaseResourceAssert) HasIsTransientEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("is_transient", ""))
	return d
}

func (d *DatabaseResourceAssert) HasLogLevelEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("log_level", ""))
	return d
}

func (d *DatabaseResourceAssert) HasMaxDataExtensionTimeInDaysEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("max_data_extension_time_in_days", ""))
	return d
}

func (d *DatabaseResourceAssert) HasQuotedIdentifiersIgnoreCaseEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("quoted_identifiers_ignore_case", ""))
	return d
}

func (d *DatabaseResourceAssert) HasReplaceInvalidCharactersEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("replace_invalid_characters", ""))
	return d
}

func (d *DatabaseResourceAssert) HasReplicationEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("replication.#", "0"))
	return d
}

func (d *DatabaseResourceAssert) HasStorageSerializationPolicyEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("storage_serialization_policy", ""))
	return d
}

func (d *DatabaseResourceAssert) HasSuspendTaskAfterNumFailuresEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("suspend_task_after_num_failures", ""))
	return d
}

func (d *DatabaseResourceAssert) HasTaskAutoRetryAttemptsEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("task_auto_retry_attempts", ""))
	return d
}

func (d *DatabaseResourceAssert) HasTraceLevelEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("trace_level", ""))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskManagedInitialWarehouseSizeEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("user_task_managed_initial_warehouse_size", ""))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskMinimumTriggerIntervalInSecondsEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("user_task_minimum_trigger_interval_in_seconds", ""))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskTimeoutMsEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValueSet("user_task_timeout_ms", ""))
	return d
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (d *DatabaseResourceAssert) HasNameNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("name"))
	return d
}

func (d *DatabaseResourceAssert) HasCatalogNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("catalog"))
	return d
}

func (d *DatabaseResourceAssert) HasCommentNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("comment"))
	return d
}

func (d *DatabaseResourceAssert) HasDataRetentionTimeInDaysNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("data_retention_time_in_days"))
	return d
}

func (d *DatabaseResourceAssert) HasDefaultDdlCollationNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("default_ddl_collation"))
	return d
}

func (d *DatabaseResourceAssert) HasDropPublicSchemaOnCreationNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("drop_public_schema_on_creation"))
	return d
}

func (d *DatabaseResourceAssert) HasEnableConsoleOutputNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("enable_console_output"))
	return d
}

func (d *DatabaseResourceAssert) HasExternalVolumeNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("external_volume"))
	return d
}

func (d *DatabaseResourceAssert) HasFullyQualifiedNameNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return d
}

func (d *DatabaseResourceAssert) HasIsTransientNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("is_transient"))
	return d
}

func (d *DatabaseResourceAssert) HasLogLevelNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("log_level"))
	return d
}

func (d *DatabaseResourceAssert) HasMaxDataExtensionTimeInDaysNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("max_data_extension_time_in_days"))
	return d
}

func (d *DatabaseResourceAssert) HasQuotedIdentifiersIgnoreCaseNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("quoted_identifiers_ignore_case"))
	return d
}

func (d *DatabaseResourceAssert) HasReplaceInvalidCharactersNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("replace_invalid_characters"))
	return d
}

func (d *DatabaseResourceAssert) HasStorageSerializationPolicyNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("storage_serialization_policy"))
	return d
}

func (d *DatabaseResourceAssert) HasSuspendTaskAfterNumFailuresNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("suspend_task_after_num_failures"))
	return d
}

func (d *DatabaseResourceAssert) HasTaskAutoRetryAttemptsNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("task_auto_retry_attempts"))
	return d
}

func (d *DatabaseResourceAssert) HasTraceLevelNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("trace_level"))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskManagedInitialWarehouseSizeNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("user_task_managed_initial_warehouse_size"))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskMinimumTriggerIntervalInSecondsNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("user_task_minimum_trigger_interval_in_seconds"))
	return d
}

func (d *DatabaseResourceAssert) HasUserTaskTimeoutMsNotEmpty() *DatabaseResourceAssert {
	d.AddAssertion(assert.ValuePresent("user_task_timeout_ms"))
	return d
}
