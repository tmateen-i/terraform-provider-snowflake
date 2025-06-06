// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type ResourceMonitorResourceAssert struct {
	*assert.ResourceAssert
}

func ResourceMonitorResource(t *testing.T, name string) *ResourceMonitorResourceAssert {
	t.Helper()

	return &ResourceMonitorResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedResourceMonitorResource(t *testing.T, id string) *ResourceMonitorResourceAssert {
	t.Helper()

	return &ResourceMonitorResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (r *ResourceMonitorResourceAssert) HasNameString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("name", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasCreditQuotaString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("credit_quota", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasEndTimestampString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("end_timestamp", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasFrequencyString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("frequency", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasFullyQualifiedNameString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNotifyTriggersString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("notify_triggers", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNotifyUsersString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("notify_users", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasStartTimestampString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("start_timestamp", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasSuspendImmediateTriggerString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("suspend_immediate_trigger", expected))
	return r
}

func (r *ResourceMonitorResourceAssert) HasSuspendTriggerString(expected string) *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("suspend_trigger", expected))
	return r
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (r *ResourceMonitorResourceAssert) HasNoName() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("name"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoCreditQuota() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("credit_quota"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoEndTimestamp() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("end_timestamp"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoFrequency() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("frequency"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoFullyQualifiedName() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoStartTimestamp() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("start_timestamp"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoSuspendImmediateTrigger() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("suspend_immediate_trigger"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNoSuspendTrigger() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueNotSet("suspend_trigger"))
	return r
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (r *ResourceMonitorResourceAssert) HasCreditQuotaEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("credit_quota", ""))
	return r
}

func (r *ResourceMonitorResourceAssert) HasEndTimestampEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("end_timestamp", ""))
	return r
}

func (r *ResourceMonitorResourceAssert) HasFrequencyEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("frequency", ""))
	return r
}

func (r *ResourceMonitorResourceAssert) HasFullyQualifiedNameEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNotifyTriggersEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("notify_triggers.#", "0"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasNotifyUsersEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("notify_users.#", "0"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasStartTimestampEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("start_timestamp", ""))
	return r
}

func (r *ResourceMonitorResourceAssert) HasSuspendImmediateTriggerEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("suspend_immediate_trigger", ""))
	return r
}

func (r *ResourceMonitorResourceAssert) HasSuspendTriggerEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValueSet("suspend_trigger", ""))
	return r
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (r *ResourceMonitorResourceAssert) HasNameNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("name"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasCreditQuotaNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("credit_quota"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasEndTimestampNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("end_timestamp"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasFrequencyNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("frequency"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasFullyQualifiedNameNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasStartTimestampNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("start_timestamp"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasSuspendImmediateTriggerNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("suspend_immediate_trigger"))
	return r
}

func (r *ResourceMonitorResourceAssert) HasSuspendTriggerNotEmpty() *ResourceMonitorResourceAssert {
	r.AddAssertion(assert.ValuePresent("suspend_trigger"))
	return r
}
