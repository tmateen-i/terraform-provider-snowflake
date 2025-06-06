// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type TagShowOutputAssert struct {
	*assert.ResourceAssert
}

func TagShowOutput(t *testing.T, name string) *TagShowOutputAssert {
	t.Helper()

	tagAssert := TagShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	tagAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &tagAssert
}

func ImportedTagShowOutput(t *testing.T, id string) *TagShowOutputAssert {
	t.Helper()

	tagAssert := TagShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	tagAssert.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &tagAssert
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (t *TagShowOutputAssert) HasCreatedOn(expected time.Time) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected.String()))
	return t
}

func (t *TagShowOutputAssert) HasName(expected string) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return t
}

func (t *TagShowOutputAssert) HasDatabaseName(expected string) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("database_name", expected))
	return t
}

func (t *TagShowOutputAssert) HasSchemaName(expected string) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return t
}

func (t *TagShowOutputAssert) HasOwner(expected string) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return t
}

func (t *TagShowOutputAssert) HasComment(expected string) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return t
}

func (t *TagShowOutputAssert) HasOwnerRoleType(expected string) *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return t
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (t *TagShowOutputAssert) HasNoCreatedOn() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("created_on"))
	return t
}

func (t *TagShowOutputAssert) HasNoName() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("name"))
	return t
}

func (t *TagShowOutputAssert) HasNoDatabaseName() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("database_name"))
	return t
}

func (t *TagShowOutputAssert) HasNoSchemaName() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("schema_name"))
	return t
}

func (t *TagShowOutputAssert) HasNoOwner() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("owner"))
	return t
}

func (t *TagShowOutputAssert) HasNoComment() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("comment"))
	return t
}

func (t *TagShowOutputAssert) HasNoAllowedValues() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueSet("allowed_values.#", "0"))
	return t
}

func (t *TagShowOutputAssert) HasNoOwnerRoleType() *TagShowOutputAssert {
	t.AddAssertion(assert.ResourceShowOutputValueNotSet("owner_role_type"))
	return t
}
