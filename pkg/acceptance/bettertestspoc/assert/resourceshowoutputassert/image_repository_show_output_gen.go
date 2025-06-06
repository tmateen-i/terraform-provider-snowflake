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

type ImageRepositoryShowOutputAssert struct {
	*assert.ResourceAssert
}

func ImageRepositoryShowOutput(t *testing.T, name string) *ImageRepositoryShowOutputAssert {
	t.Helper()

	i := ImageRepositoryShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	i.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &i
}

func ImportedImageRepositoryShowOutput(t *testing.T, id string) *ImageRepositoryShowOutputAssert {
	t.Helper()

	i := ImageRepositoryShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	i.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &i
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (i *ImageRepositoryShowOutputAssert) HasCreatedOn(expected time.Time) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected.String()))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasName(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasDatabaseName(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("database_name", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasSchemaName(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasRepositoryUrl(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("repository_url", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasOwner(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasOwnerRoleType(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasComment(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return i
}

func (i *ImageRepositoryShowOutputAssert) HasPrivatelinkRepositoryUrl(expected string) *ImageRepositoryShowOutputAssert {
	i.AddAssertion(assert.ResourceShowOutputValueSet("privatelink_repository_url", expected))
	return i
}
