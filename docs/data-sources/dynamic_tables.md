---
page_title: "snowflake_dynamic_tables Data Source - terraform-provider-snowflake"
subcategory: "Preview"
description: |-
  
---

!> **Caution: Preview Feature** This feature is considered a preview feature in the provider, regardless of the state of the resource in Snowflake. We do not guarantee its stability. It will be reworked and marked as a stable feature in future releases. Breaking changes are expected, even without bumping the major version. To use this feature, add the relevant feature name to `preview_features_enabled` field in the [provider configuration](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs#schema). Please always refer to the [Getting Help](https://github.com/snowflakedb/terraform-provider-snowflake?tab=readme-ov-file#getting-help) section in our Github repo to best determine how to get help for your questions.

# snowflake_dynamic_tables (Data Source)





-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `in` (Block List, Max: 1) IN clause to filter the list of dynamic tables. (see [below for nested schema](#nestedblock--in))
- `like` (Block List, Max: 1) LIKE clause to filter the list of dynamic tables. (see [below for nested schema](#nestedblock--like))
- `limit` (Block List, Max: 1) Optionally limits the maximum number of rows returned, while also enabling “pagination” of the results. Note that the actual number of rows returned might be less than the specified limit (e.g. the number of existing objects is less than the specified limit). (see [below for nested schema](#nestedblock--limit))
- `starts_with` (String) Optionally filters the command output based on the characters that appear at the beginning of the object name. The string is case-sensitive.

### Read-Only

- `id` (String) The ID of this resource.
- `records` (List of Object) The list of dynamic tables. (see [below for nested schema](#nestedatt--records))

<a id="nestedblock--in"></a>
### Nested Schema for `in`

Optional:

- `account` (Boolean) Returns records for the entire account.
- `database` (String) Returns records for the current database in use or for a specified database (db_name).
- `schema` (String) Returns records for the current schema in use or a specified schema (schema_name).


<a id="nestedblock--like"></a>
### Nested Schema for `like`

Required:

- `pattern` (String) Filters the command output by object name. The filter uses case-insensitive pattern matching with support for SQL wildcard characters (% and _).


<a id="nestedblock--limit"></a>
### Nested Schema for `limit`

Optional:

- `from` (String) The optional FROM 'name_string' subclause effectively serves as a “cursor” for the results. This enables fetching the specified number of rows following the first row whose object name matches the specified string
- `rows` (Number) Specifies the maximum number of rows to return.


<a id="nestedatt--records"></a>
### Nested Schema for `records`

Read-Only:

- `automatic_clustering` (Boolean)
- `bytes` (Number)
- `cluster_by` (String)
- `comment` (String)
- `created_on` (String)
- `data_timestamp` (String)
- `database_name` (String)
- `is_clone` (Boolean)
- `is_replica` (Boolean)
- `last_suspended_on` (String)
- `name` (String)
- `owner` (String)
- `refresh_mode` (String)
- `refresh_mode_reason` (String)
- `rows` (Number)
- `scheduling_state` (String)
- `schema_name` (String)
- `target_lag` (String)
- `text` (String)
- `warehouse` (String)
