// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type TagResourceAssert struct {
	*assert.ResourceAssert
}

func TagResource(t *testing.T, name string) *TagResourceAssert {
	t.Helper()

	return &TagResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedTagResource(t *testing.T, id string) *TagResourceAssert {
	t.Helper()

	return &TagResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (t *TagResourceAssert) HasDatabaseString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("database", expected))
	return t
}

func (t *TagResourceAssert) HasSchemaString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("schema", expected))
	return t
}

func (t *TagResourceAssert) HasNameString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("name", expected))
	return t
}

func (t *TagResourceAssert) HasAllowedValuesString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("allowed_values", expected))
	return t
}

func (t *TagResourceAssert) HasCommentString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("comment", expected))
	return t
}

func (t *TagResourceAssert) HasFullyQualifiedNameString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return t
}

func (t *TagResourceAssert) HasMaskingPoliciesString(expected string) *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("masking_policies", expected))
	return t
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (t *TagResourceAssert) HasNoDatabase() *TagResourceAssert {
	t.AddAssertion(assert.ValueNotSet("database"))
	return t
}

func (t *TagResourceAssert) HasNoSchema() *TagResourceAssert {
	t.AddAssertion(assert.ValueNotSet("schema"))
	return t
}

func (t *TagResourceAssert) HasNoName() *TagResourceAssert {
	t.AddAssertion(assert.ValueNotSet("name"))
	return t
}

func (t *TagResourceAssert) HasNoComment() *TagResourceAssert {
	t.AddAssertion(assert.ValueNotSet("comment"))
	return t
}

func (t *TagResourceAssert) HasNoFullyQualifiedName() *TagResourceAssert {
	t.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return t
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (t *TagResourceAssert) HasAllowedValuesEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("allowed_values.#", "0"))
	return t
}

func (t *TagResourceAssert) HasCommentEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("comment", ""))
	return t
}

func (t *TagResourceAssert) HasFullyQualifiedNameEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return t
}

func (t *TagResourceAssert) HasMaskingPoliciesEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValueSet("masking_policies.#", "0"))
	return t
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (t *TagResourceAssert) HasDatabaseNotEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValuePresent("database"))
	return t
}

func (t *TagResourceAssert) HasSchemaNotEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValuePresent("schema"))
	return t
}

func (t *TagResourceAssert) HasNameNotEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValuePresent("name"))
	return t
}

func (t *TagResourceAssert) HasCommentNotEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValuePresent("comment"))
	return t
}

func (t *TagResourceAssert) HasFullyQualifiedNameNotEmpty() *TagResourceAssert {
	t.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return t
}
