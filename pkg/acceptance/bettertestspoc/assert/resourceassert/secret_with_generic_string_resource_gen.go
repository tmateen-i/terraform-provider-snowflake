// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type SecretWithGenericStringResourceAssert struct {
	*assert.ResourceAssert
}

func SecretWithGenericStringResource(t *testing.T, name string) *SecretWithGenericStringResourceAssert {
	t.Helper()

	return &SecretWithGenericStringResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedSecretWithGenericStringResource(t *testing.T, id string) *SecretWithGenericStringResourceAssert {
	t.Helper()

	return &SecretWithGenericStringResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (s *SecretWithGenericStringResourceAssert) HasDatabaseString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("database", expected))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSchemaString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("schema", expected))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNameString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("name", expected))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasCommentString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("comment", expected))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasFullyQualifiedNameString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSecretStringString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("secret_string", expected))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSecretTypeString(expected string) *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("secret_type", expected))
	return s
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (s *SecretWithGenericStringResourceAssert) HasNoDatabase() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("database"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNoSchema() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("schema"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNoName() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("name"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNoComment() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("comment"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNoFullyQualifiedName() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNoSecretString() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("secret_string"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNoSecretType() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueNotSet("secret_type"))
	return s
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (s *SecretWithGenericStringResourceAssert) HasCommentEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("comment", ""))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasFullyQualifiedNameEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSecretTypeEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValueSet("secret_type", ""))
	return s
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (s *SecretWithGenericStringResourceAssert) HasDatabaseNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("database"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSchemaNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("schema"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasNameNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("name"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasCommentNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("comment"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasFullyQualifiedNameNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSecretStringNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("secret_string"))
	return s
}

func (s *SecretWithGenericStringResourceAssert) HasSecretTypeNotEmpty() *SecretWithGenericStringResourceAssert {
	s.AddAssertion(assert.ValuePresent("secret_type"))
	return s
}
