// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type ExternalOauthSecurityIntegrationResourceAssert struct {
	*assert.ResourceAssert
}

func ExternalOauthSecurityIntegrationResource(t *testing.T, name string) *ExternalOauthSecurityIntegrationResourceAssert {
	t.Helper()

	return &ExternalOauthSecurityIntegrationResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedExternalOauthSecurityIntegrationResource(t *testing.T, id string) *ExternalOauthSecurityIntegrationResourceAssert {
	t.Helper()

	return &ExternalOauthSecurityIntegrationResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNameString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("name", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasCommentString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("comment", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasEnabledString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("enabled", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAllowedRolesListString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_allowed_roles_list", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAnyRoleModeString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_any_role_mode", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAudienceListString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_audience_list", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthBlockedRolesListString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_blocked_roles_list", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthIssuerString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_issuer", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthJwsKeysUrlString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_jws_keys_url", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthRsaPublicKeyString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_rsa_public_key", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthRsaPublicKey2String(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_rsa_public_key_2", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthScopeDelimiterString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_scope_delimiter", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthScopeMappingAttributeString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_scope_mapping_attribute", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthSnowflakeUserMappingAttributeString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_snowflake_user_mapping_attribute", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthTokenUserMappingClaimString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_token_user_mapping_claim", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthTypeString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_type", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasFullyQualifiedNameString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasRelatedParametersString(expected string) *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("related_parameters", expected))
	return e
}

///////////////////////////////
// Attribute no value checks //
///////////////////////////////

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoName() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("name"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoComment() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("comment"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoEnabled() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("enabled"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthAnyRoleMode() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_any_role_mode"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthIssuer() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_issuer"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthRsaPublicKey() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_rsa_public_key"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthRsaPublicKey2() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_rsa_public_key_2"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthScopeDelimiter() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_scope_delimiter"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthScopeMappingAttribute() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_scope_mapping_attribute"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthSnowflakeUserMappingAttribute() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_snowflake_user_mapping_attribute"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoExternalOauthType() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("external_oauth_type"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNoFullyQualifiedName() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return e
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasCommentEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("comment", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAllowedRolesListEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_allowed_roles_list.#", "0"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAnyRoleModeEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_any_role_mode", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAudienceListEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_audience_list.#", "0"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthBlockedRolesListEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_blocked_roles_list.#", "0"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthJwsKeysUrlEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_jws_keys_url.#", "0"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthRsaPublicKeyEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_rsa_public_key", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthRsaPublicKey2Empty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_rsa_public_key_2", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthScopeDelimiterEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_scope_delimiter", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthScopeMappingAttributeEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("external_oauth_scope_mapping_attribute", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasFullyQualifiedNameEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("fully_qualified_name", ""))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasRelatedParametersEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValueSet("related_parameters.#", "0"))
	return e
}

///////////////////////////////
// Attribute presence checks //
///////////////////////////////

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasNameNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("name"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasCommentNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("comment"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasEnabledNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("enabled"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthAnyRoleModeNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_any_role_mode"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthIssuerNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_issuer"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthRsaPublicKeyNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_rsa_public_key"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthRsaPublicKey2NotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_rsa_public_key_2"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthScopeDelimiterNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_scope_delimiter"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthScopeMappingAttributeNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_scope_mapping_attribute"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthSnowflakeUserMappingAttributeNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_snowflake_user_mapping_attribute"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasExternalOauthTypeNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("external_oauth_type"))
	return e
}

func (e *ExternalOauthSecurityIntegrationResourceAssert) HasFullyQualifiedNameNotEmpty() *ExternalOauthSecurityIntegrationResourceAssert {
	e.AddAssertion(assert.ValuePresent("fully_qualified_name"))
	return e
}
