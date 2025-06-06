// Code generated by assertions generator; DO NOT EDIT.

package objectassert

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type SecretAssert struct {
	*assert.SnowflakeObjectAssert[sdk.Secret, sdk.SchemaObjectIdentifier]
}

func Secret(t *testing.T, id sdk.SchemaObjectIdentifier) *SecretAssert {
	t.Helper()
	return &SecretAssert{
		assert.NewSnowflakeObjectAssertWithTestClientObjectProvider(sdk.ObjectTypeSecret, id, func(testClient *helpers.TestClient) assert.ObjectProvider[sdk.Secret, sdk.SchemaObjectIdentifier] {
			return testClient.Secret.Show
		}),
	}
}

func SecretFromObject(t *testing.T, secret *sdk.Secret) *SecretAssert {
	t.Helper()
	return &SecretAssert{
		assert.NewSnowflakeObjectAssertWithObject(sdk.ObjectTypeSecret, secret.ID(), secret),
	}
}

func (s *SecretAssert) HasCreatedOn(expected time.Time) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.CreatedOn != expected {
			return fmt.Errorf("expected created on: %v; got: %v", expected, o.CreatedOn)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasName(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.Name != expected {
			return fmt.Errorf("expected name: %v; got: %v", expected, o.Name)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasSchemaName(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.SchemaName != expected {
			return fmt.Errorf("expected schema name: %v; got: %v", expected, o.SchemaName)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasDatabaseName(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.DatabaseName != expected {
			return fmt.Errorf("expected database name: %v; got: %v", expected, o.DatabaseName)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasOwner(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.Owner != expected {
			return fmt.Errorf("expected owner: %v; got: %v", expected, o.Owner)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasComment(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.Comment == nil {
			return fmt.Errorf("expected comment to have value; got: nil")
		}
		if *o.Comment != expected {
			return fmt.Errorf("expected comment: %v; got: %v", expected, *o.Comment)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasSecretType(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.SecretType != expected {
			return fmt.Errorf("expected secret type: %v; got: %v", expected, o.SecretType)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasOauthScopes(expected ...string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		mapped := collections.Map(o.OauthScopes, func(item string) any { return item })
		mappedExpected := collections.Map(expected, func(item string) any { return item })
		if !slices.Equal(mapped, mappedExpected) {
			return fmt.Errorf("expected oauth scopes: %v; got: %v", expected, o.OauthScopes)
		}
		return nil
	})
	return s
}

func (s *SecretAssert) HasOwnerRoleType(expected string) *SecretAssert {
	s.AddAssertion(func(t *testing.T, o *sdk.Secret) error {
		t.Helper()
		if o.OwnerRoleType != expected {
			return fmt.Errorf("expected owner role type: %v; got: %v", expected, o.OwnerRoleType)
		}
		return nil
	})
	return s
}
