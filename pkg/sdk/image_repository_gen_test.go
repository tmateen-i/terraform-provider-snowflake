package sdk

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/helpers/random"
)

func TestImageRepositories_Create(t *testing.T) {
	id := randomSchemaObjectIdentifier()
	// Minimal valid CreateImageRepositoryOptions
	defaultOpts := func() *CreateImageRepositoryOptions {
		return &CreateImageRepositoryOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *CreateImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})
	t.Run("validation: invalid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = emptySchemaObjectIdentifier
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: conflicting fields for [opts.IfNotExists opts.OrReplace]", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.OrReplace = Bool(true)
		assertOptsInvalidJoinedErrors(t, opts, errOneOf("CreateImageRepositoryOptions", "IfNotExists", "OrReplace"))
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, "CREATE IMAGE REPOSITORY %s", id.FullyQualifiedName())
	})

	t.Run("all options", func(t *testing.T) {
		comment := random.Comment()
		tagId := NewAccountObjectIdentifier("tag1")
		opts := defaultOpts()
		opts.IfNotExists = Bool(true)
		opts.Comment = &comment
		opts.Tag = []TagAssociation{
			{
				Name:  tagId,
				Value: "value1",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, "CREATE IMAGE REPOSITORY IF NOT EXISTS %s COMMENT = '%s' TAG (%s = 'value1')", id.FullyQualifiedName(), comment, tagId.FullyQualifiedName())
	})
}

func TestImageRepositories_Alter(t *testing.T) {
	id := randomSchemaObjectIdentifier()
	// Minimal valid AlterImageRepositoryOptions
	defaultOpts := func() *AlterImageRepositoryOptions {
		return &AlterImageRepositoryOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *AlterImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: invalid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = emptySchemaObjectIdentifier
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("validation: exactly one field from [opts.Set opts.Unset opts.SetTags opts.UnsetTags] should be present", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsInvalidJoinedErrors(t, opts, errExactlyOneOf("AlterImageRepositoryOptions", "Set", "SetTags", "UnsetTags"))
	})

	t.Run("set: all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &ImageRepositorySet{
			Comment: &StringAllowEmpty{
				Value: "test",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, "ALTER IMAGE REPOSITORY %s SET COMMENT = 'test'", id.FullyQualifiedName())
	})

	t.Run("set: empty comment", func(t *testing.T) {
		opts := defaultOpts()
		opts.Set = &ImageRepositorySet{
			Comment: &StringAllowEmpty{
				Value: "",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, "ALTER IMAGE REPOSITORY %s SET COMMENT = ''", id.FullyQualifiedName())
	})

	t.Run("set tags", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfExists = Bool(true)
		opts.SetTags = []TagAssociation{
			{
				Name:  NewAccountObjectIdentifier("tag1"),
				Value: "value1",
			},
			{
				Name:  NewAccountObjectIdentifier("tag2"),
				Value: "value2",
			},
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER IMAGE REPOSITORY IF EXISTS %s SET TAG "tag1" = 'value1', "tag2" = 'value2'`, id.FullyQualifiedName())
	})

	t.Run("unset tags", func(t *testing.T) {
		opts := defaultOpts()
		opts.UnsetTags = []ObjectIdentifier{
			NewAccountObjectIdentifier("tag1"),
			NewAccountObjectIdentifier("tag2"),
		}
		assertOptsValidAndSQLEquals(t, opts, `ALTER IMAGE REPOSITORY %s UNSET TAG "tag1", "tag2"`, id.FullyQualifiedName())
	})
}

func TestImageRepositories_Drop(t *testing.T) {
	id := randomSchemaObjectIdentifier()
	// Minimal valid DropImageRepositoryOptions
	defaultOpts := func() *DropImageRepositoryOptions {
		return &DropImageRepositoryOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *DropImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})
	t.Run("validation: invalid identifier for [opts.name]", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = emptySchemaObjectIdentifier
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.IfExists = Bool(true)
		assertOptsValidAndSQLEquals(t, opts, "DROP IMAGE REPOSITORY IF EXISTS %s", id.FullyQualifiedName())
	})
}

func TestImageRepositories_Show(t *testing.T) {
	// Minimal valid ShowImageRepositoryOptions
	defaultOpts := func() *ShowImageRepositoryOptions {
		return &ShowImageRepositoryOptions{}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		var opts *ShowImageRepositoryOptions = nil
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("basic", func(t *testing.T) {
		opts := defaultOpts()
		assertOptsValidAndSQLEquals(t, opts, "SHOW IMAGE REPOSITORIES")
	})

	t.Run("like", func(t *testing.T) {
		opts := defaultOpts()
		opts.Like = &Like{
			Pattern: String("pattern"),
		}
		assertOptsValidAndSQLEquals(t, opts, "SHOW IMAGE REPOSITORIES LIKE 'pattern'")
	})

	t.Run("in", func(t *testing.T) {
		schemaId := randomDatabaseObjectIdentifier()
		opts := defaultOpts()
		opts.In = &In{
			Schema: schemaId,
		}
		assertOptsValidAndSQLEquals(t, opts, "SHOW IMAGE REPOSITORIES IN SCHEMA %s", schemaId.FullyQualifiedName())
	})
}
