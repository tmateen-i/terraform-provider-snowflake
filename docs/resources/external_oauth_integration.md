---
page_title: "snowflake_external_oauth_integration Resource - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  Resource used to manage external oauth security integration objects. For more information, check security integrations documentation https://docs.snowflake.com/en/sql-reference/sql/create-security-integration-oauth-external.
---

!> **Note** The provider does not detect external changes on security integration type. In this case, remove the integration of wrong type manually with `terraform destroy` and recreate the resource. It will be addressed in the future.

# snowflake_external_oauth_integration (Resource)

Resource used to manage external oauth security integration objects. For more information, check [security integrations documentation](https://docs.snowflake.com/en/sql-reference/sql/create-security-integration-oauth-external).

## Example Usage

```terraform
# basic resource
resource "snowflake_external_oauth_integration" "test" {
  enabled                                         = true
  external_oauth_issuer                           = "issuer"
  external_oauth_snowflake_user_mapping_attribute = "LOGIN_NAME"
  external_oauth_token_user_mapping_claim         = ["upn"]
  name                                            = "test"
  external_oauth_type                             = "CUSTOM"
}
# resource with all fields set (jws keys url and allowed roles)
resource "snowflake_external_oauth_integration" "test" {
  comment                                         = "comment"
  enabled                                         = true
  external_oauth_allowed_roles_list               = [snowflake_role.one.fully_qualified_name]
  external_oauth_any_role_mode                    = "ENABLE"
  external_oauth_audience_list                    = ["https://example.com"]
  external_oauth_issuer                           = "issuer"
  external_oauth_jws_keys_url                     = ["https://example.com"]
  external_oauth_scope_delimiter                  = ","
  external_oauth_scope_mapping_attribute          = "scope"
  external_oauth_snowflake_user_mapping_attribute = "LOGIN_NAME"
  external_oauth_token_user_mapping_claim         = ["upn"]
  name                                            = "test"
  external_oauth_type                             = "CUSTOM"
}
# resource with all fields set (rsa public keys and blocked roles)
resource "snowflake_external_oauth_integration" "test" {
  comment                                         = "comment"
  enabled                                         = true
  external_oauth_any_role_mode                    = "ENABLE"
  external_oauth_audience_list                    = ["https://example.com"]
  external_oauth_blocked_roles_list               = [snowflake_role.one.fully_qualified_name]
  external_oauth_issuer                           = "issuer"
  external_oauth_rsa_public_key                   = file("key.pem")
  external_oauth_rsa_public_key_2                 = file("key2.pem")
  external_oauth_scope_delimiter                  = ","
  external_oauth_scope_mapping_attribute          = "scope"
  external_oauth_snowflake_user_mapping_attribute = "LOGIN_NAME"
  external_oauth_token_user_mapping_claim         = ["upn"]
  name                                            = "test"
  external_oauth_type                             = "CUSTOM"
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Specifies whether to initiate operation of the integration or suspend it.
- `external_oauth_issuer` (String) Specifies the URL to define the OAuth 2.0 authorization server.
- `external_oauth_snowflake_user_mapping_attribute` (String) Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record. Valid values are (case-insensitive): `LOGIN_NAME` | `EMAIL_ADDRESS`.
- `external_oauth_token_user_mapping_claim` (Set of String) Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record. If removed from the config, the resource is recreated.
- `external_oauth_type` (String) Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server. Valid values are (case-insensitive): `OKTA` | `AZURE` | `PING_FEDERATE` | `CUSTOM`.
- `name` (String) Specifies the name of the External Oath integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.

### Optional

- `comment` (String) Specifies a comment for the OAuth integration.
- `external_oauth_allowed_roles_list` (Set of String) Specifies the list of roles that the client can set as the primary role. For more information about this resource, see [docs](./account_role).
- `external_oauth_any_role_mode` (String) Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token. Valid values are (case-insensitive): `DISABLE` | `ENABLE` | `ENABLE_FOR_PRIVILEGE`.
- `external_oauth_audience_list` (Set of String) Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL
- `external_oauth_blocked_roles_list` (Set of String) Specifies the list of roles that a client cannot set as the primary role. By default, this list includes the ACCOUNTADMIN, ORGADMIN and SECURITYADMIN roles. To remove these privileged roles from the list, use the ALTER ACCOUNT command to set the EXTERNAL_OAUTH_ADD_PRIVILEGED_ROLES_TO_BLOCKED_LIST account parameter to FALSE. For more information about this resource, see [docs](./account_role).
- `external_oauth_jws_keys_url` (Set of String) Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token. The maximum number of URLs that can be specified in the list is 3. If removed from the config, the resource is recreated.
- `external_oauth_rsa_public_key` (String) Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. If removed from the config, the resource is recreated.
- `external_oauth_rsa_public_key_2` (String) Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. Used for key rotation. If removed from the config, the resource is recreated.
- `external_oauth_scope_delimiter` (String) Specifies the scope delimiter in the authorization token.
- `external_oauth_scope_mapping_attribute` (String) Specifies the access token claim to map the access token to an account role. If removed from the config, the resource is recreated.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `describe_output` (List of Object) Outputs the result of `DESCRIBE SECURITY INTEGRATIONS` for the given security integration. (see [below for nested schema](#nestedatt--describe_output))
- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `id` (String) The ID of this resource.
- `related_parameters` (List of Object) Parameters related to this security integration. (see [below for nested schema](#nestedatt--related_parameters))
- `show_output` (List of Object) Outputs the result of `SHOW SECURITY INTEGRATIONS` for the given security integration. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--describe_output"></a>
### Nested Schema for `describe_output`

Read-Only:

- `comment` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--comment))
- `enabled` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--enabled))
- `external_oauth_allowed_roles_list` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_allowed_roles_list))
- `external_oauth_any_role_mode` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_any_role_mode))
- `external_oauth_audience_list` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_audience_list))
- `external_oauth_blocked_roles_list` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_blocked_roles_list))
- `external_oauth_issuer` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_issuer))
- `external_oauth_jws_keys_url` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_jws_keys_url))
- `external_oauth_rsa_public_key` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_rsa_public_key))
- `external_oauth_rsa_public_key_2` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_rsa_public_key_2))
- `external_oauth_scope_delimiter` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_scope_delimiter))
- `external_oauth_snowflake_user_mapping_attribute` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_snowflake_user_mapping_attribute))
- `external_oauth_token_user_mapping_claim` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--external_oauth_token_user_mapping_claim))

<a id="nestedobjatt--describe_output--comment"></a>
### Nested Schema for `describe_output.comment`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--enabled"></a>
### Nested Schema for `describe_output.enabled`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_allowed_roles_list"></a>
### Nested Schema for `describe_output.external_oauth_allowed_roles_list`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_any_role_mode"></a>
### Nested Schema for `describe_output.external_oauth_any_role_mode`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_audience_list"></a>
### Nested Schema for `describe_output.external_oauth_audience_list`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_blocked_roles_list"></a>
### Nested Schema for `describe_output.external_oauth_blocked_roles_list`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_issuer"></a>
### Nested Schema for `describe_output.external_oauth_issuer`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_jws_keys_url"></a>
### Nested Schema for `describe_output.external_oauth_jws_keys_url`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_rsa_public_key"></a>
### Nested Schema for `describe_output.external_oauth_rsa_public_key`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_rsa_public_key_2"></a>
### Nested Schema for `describe_output.external_oauth_rsa_public_key_2`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_scope_delimiter"></a>
### Nested Schema for `describe_output.external_oauth_scope_delimiter`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_snowflake_user_mapping_attribute"></a>
### Nested Schema for `describe_output.external_oauth_snowflake_user_mapping_attribute`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--external_oauth_token_user_mapping_claim"></a>
### Nested Schema for `describe_output.external_oauth_token_user_mapping_claim`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)



<a id="nestedatt--related_parameters"></a>
### Nested Schema for `related_parameters`

Read-Only:

- `external_oauth_add_privileged_roles_to_blocked_list` (List of Object) (see [below for nested schema](#nestedobjatt--related_parameters--external_oauth_add_privileged_roles_to_blocked_list))

<a id="nestedobjatt--related_parameters--external_oauth_add_privileged_roles_to_blocked_list"></a>
### Nested Schema for `related_parameters.external_oauth_add_privileged_roles_to_blocked_list`

Read-Only:

- `default` (String)
- `description` (String)
- `key` (String)
- `level` (String)
- `value` (String)



<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `category` (String)
- `comment` (String)
- `created_on` (String)
- `enabled` (Boolean)
- `integration_type` (String)
- `name` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_external_oauth_integration.example '"<integration_name>"'
```
