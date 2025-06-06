// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type StreamOnExternalTableResourceAssert struct {
	*assert.ResourceAssert
}

func StreamOnExternalTableResource(t *testing.T, name string) *StreamOnExternalTableResourceAssert {
	t.Helper()

	return &StreamOnExternalTableResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedStreamOnExternalTableResource(t *testing.T, id string) *StreamOnExternalTableResourceAssert {
	t.Helper()

	return &StreamOnExternalTableResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (s *StreamOnExternalTableResourceAssert) HasDatabaseString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("database", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasSchemaString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("schema", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNameString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("name", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasAtString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("at", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasBeforeString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("before", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasCommentString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("comment", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasCopyGrantsString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("copy_grants", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasExternalTableString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("external_table", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasFullyQualifiedNameString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasInsertOnlyString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("insert_only", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasStaleString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("stale", expected))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasStreamTypeString(expected string) *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("stream_type", expected))
	return s
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (s *StreamOnExternalTableResourceAssert) HasNoDatabase() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("database"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoSchema() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("schema"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoName() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("name"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoComment() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("comment"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoCopyGrants() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("copy_grants"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoExternalTable() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("external_table"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoFullyQualifiedName() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoInsertOnly() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("insert_only"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoStale() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("stale"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNoStreamType() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("stream_type"))
	return s
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (s *StreamOnExternalTableResourceAssert) HasAtEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("at.#", "0"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasBeforeEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("before.#", "0"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasCommentEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("comment", ""))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasCopyGrantsEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("copy_grants", ""))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasFullyQualifiedNameEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasInsertOnlyEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("insert_only", ""))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasStaleEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("stale", ""))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasStreamTypeEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValueSet("stream_type", ""))
	return s
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (s *StreamOnExternalTableResourceAssert) HasDatabaseNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("database"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasSchemaNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("schema"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasNameNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("name"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasCommentNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("comment"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasCopyGrantsNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("copy_grants"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasExternalTableNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("external_table"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasFullyQualifiedNameNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasInsertOnlyNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("insert_only"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasStaleNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("stale"))
	return s
}

func (s *StreamOnExternalTableResourceAssert) HasStreamTypeNotEmpty() *StreamOnExternalTableResourceAssert {
	s.AddAssertion(assert.ValuePresent("stream_type"))
	return s
}
