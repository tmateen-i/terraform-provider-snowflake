---
page_title: "snowflake_function_javascript Resource - terraform-provider-snowflake"
subcategory: "Preview"
description: |-
  Resource used to manage javascript function objects. For more information, check function documentation https://docs.snowflake.com/en/sql-reference/sql/create-function.
---

!> **Caution: Preview Feature** This feature is considered a preview feature in the provider, regardless of the state of the resource in Snowflake. We do not guarantee its stability. It will be reworked and marked as a stable feature in future releases. Breaking changes are expected, even without bumping the major version. To use this feature, add the relevant feature name to `preview_features_enabled` field in the [provider configuration](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs#schema). Please always refer to the [Getting Help](https://github.com/snowflakedb/terraform-provider-snowflake?tab=readme-ov-file#getting-help) section in our Github repo to best determine how to get help for your questions.

!> **Sensitive values** This resource's `function_definition` and `show_output.arguments_raw` fields are not marked as sensitive in the provider. Ensure that no personal data, sensitive data, export-controlled data, or other regulated data is entered as metadata when using the provider. If you use one of these fields, they may be present in logs, so ensure that the provider logs are properly restricted. For more information, see [Sensitive values limitations](../#sensitive-values-limitations) and [Metadata fields in Snowflake](https://docs.snowflake.com/en/sql-reference/metadata).

-> **Note** External changes to `is_secure`, `return_results_behavior`, and `null_input_behavior` are not currently supported. They will be handled in the following versions of the provider which may still affect this resource.

-> **Note** `COPY GRANTS` and `OR REPLACE` are not currently supported.

-> **Note** `RETURN... [[ NOT ] NULL]` is not currently supported. It will be improved in the following versions of the provider which may still affect this resource.

-> **Note** Use of return type `TABLE` is currently limited. It will be improved in the following versions of the provider which may still affect this resource.

-> **Note** Snowflake is not returning full data type information for arguments which may lead to unexpected plan outputs. Diff suppression for such cases will be improved.

-> **Note** Snowflake is not returning the default values for arguments so argument's `arg_default_value` external changes cannot be tracked.

-> **Note** Limit the use of special characters (`.`, `'`, `/`, `"`, `(`, `)`, `[`, `]`, `{`, `}`, ` `) in argument names, stage ids, and secret ids. It's best to limit to only alphanumeric and underscores. There is a lot of parsing of SHOW/DESCRIBE outputs involved and using special characters may limit the possibility to achieve the correct results.

~> **Required warehouse** This resource may require active warehouse. Please, make sure you have either set a DEFAULT_WAREHOUSE for the user, or specified a warehouse in the provider configuration.

# snowflake_function_javascript (Resource)

Resource used to manage javascript function objects. For more information, check [function documentation](https://docs.snowflake.com/en/sql-reference/sql/create-function).

## Example Usage

```terraform
# Minimal
resource "snowflake_function_javascript" "minimal" {
  database = snowflake_database.test.name
  schema   = snowflake_schema.test.name
  name     = "my_javascript_function"
  arguments {
    arg_data_type = "VARIANT"
    arg_name      = "x"
  }
  return_type         = "VARIANT"
  function_definition = <<EOT
  if (x == 0) {
    return 1;
  } else {
    return 2;
  }
EOT
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database` (String) The database in which to create the function. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
- `function_definition` (String) Defines the handler code executed when the UDF is called. Wrapping `$$` signs are added by the provider automatically; do not include them. The `function_definition` value must be JavaScript source code. For more information, see [Introduction to JavaScript UDFs](https://docs.snowflake.com/en/developer-guide/udf/javascript/udf-javascript-introduction). To mitigate permadiff on this field, the provider replaces blank characters with a space. This can lead to false positives in cases where a change in case or run of whitespace is semantically significant.
- `name` (String) The name of the function; the identifier does not need to be unique for the schema in which the function is created because UDFs are identified and resolved by the combination of the name and argument types. Check the [docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#all-languages). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
- `return_type` (String) Specifies the results returned by the UDF, which determines the UDF type. Use `<result_data_type>` to create a scalar UDF that returns a single value with the specified data type. Use `TABLE (col_name col_data_type, ...)` to creates a table UDF that returns tabular results with the specified table column(s) and column type(s). For the details, consult the [docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#all-languages).
- `schema` (String) The schema in which to create the function. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.

### Optional

- `arguments` (Block List) List of the arguments for the function. Consult the [docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#all-languages) for more details. (see [below for nested schema](#nestedblock--arguments))
- `comment` (String) (Default: `user-defined function`) Specifies a comment for the function.
- `enable_console_output` (Boolean) Enable stdout/stderr fast path logging for anonymous stored procs. This is a public parameter (similar to LOG_LEVEL). For more information, check [ENABLE_CONSOLE_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-console-output).
- `is_secure` (String) (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies that the function is secure. By design, the Snowflake's `SHOW FUNCTIONS` command does not provide information about secure functions (consult [function docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#id1) and [Protecting Sensitive Information with Secure UDFs and Stored Procedures](https://docs.snowflake.com/en/developer-guide/secure-udf-procedure)) which is essential to manage/import function with Terraform. Use the role owning the function while managing secure functions. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
- `log_level` (String) LOG_LEVEL to use when filtering events For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
- `metric_level` (String) METRIC_LEVEL value to control whether to emit metrics to Event Table For more information, check [METRIC_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#metric-level).
- `null_input_behavior` (String) Specifies the behavior of the function when called with null inputs. Valid values are (case-insensitive): `CALLED ON NULL INPUT` | `RETURNS NULL ON NULL INPUT`.
- `return_results_behavior` (String) Specifies the behavior of the function when returning results. Valid values are (case-insensitive): `VOLATILE` | `IMMUTABLE`.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `trace_level` (String) Trace level value to use when generating/filtering trace events For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).

### Read-Only

- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `function_language` (String) Specifies language for the user. Used to detect external changes.
- `id` (String) The ID of this resource.
- `parameters` (List of Object) Outputs the result of `SHOW PARAMETERS IN FUNCTION` for the given function. (see [below for nested schema](#nestedatt--parameters))
- `show_output` (List of Object) Outputs the result of `SHOW FUNCTION` for the given function. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedblock--arguments"></a>
### Nested Schema for `arguments`

Required:

- `arg_data_type` (String) The argument type.
- `arg_name` (String) The argument name. The provider wraps it in double quotes by default, so be aware of that while referencing the argument in the function definition.

Optional:

- `arg_default_value` (String) Optional default value for the argument. For text values use single quotes. Numeric values can be unquoted. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Read-Only:

- `enable_console_output` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--enable_console_output))
- `log_level` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--log_level))
- `metric_level` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--metric_level))
- `trace_level` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--trace_level))

<a id="nestedobjatt--parameters--enable_console_output"></a>
### Nested Schema for `parameters.enable_console_output`

Read-Only:

- `default` (String)
- `description` (String)
- `key` (String)
- `level` (String)
- `value` (String)


<a id="nestedobjatt--parameters--log_level"></a>
### Nested Schema for `parameters.log_level`

Read-Only:

- `default` (String)
- `description` (String)
- `key` (String)
- `level` (String)
- `value` (String)


<a id="nestedobjatt--parameters--metric_level"></a>
### Nested Schema for `parameters.metric_level`

Read-Only:

- `default` (String)
- `description` (String)
- `key` (String)
- `level` (String)
- `value` (String)


<a id="nestedobjatt--parameters--trace_level"></a>
### Nested Schema for `parameters.trace_level`

Read-Only:

- `default` (String)
- `description` (String)
- `key` (String)
- `level` (String)
- `value` (String)



<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `arguments_raw` (String)
- `catalog_name` (String)
- `created_on` (String)
- `description` (String)
- `external_access_integrations` (String)
- `is_aggregate` (Boolean)
- `is_ansi` (Boolean)
- `is_builtin` (Boolean)
- `is_data_metric` (Boolean)
- `is_external_function` (Boolean)
- `is_memoizable` (Boolean)
- `is_secure` (Boolean)
- `is_table_function` (Boolean)
- `language` (String)
- `max_num_arguments` (Number)
- `min_num_arguments` (Number)
- `name` (String)
- `schema_name` (String)
- `secrets` (String)
- `valid_for_clustering` (Boolean)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_function_javascript.example '"<database_name>"."<schema_name>"."<function_name>"(varchar, varchar, varchar)'
```

Note: Snowflake is not returning all information needed to populate the state correctly after import (e.g. data types with attributes like NUMBER(32, 10) are returned as NUMBER, default values for arguments are not returned at all).
Also, `ALTER` for functions is very limited so most of the attributes on this resource are marked as force new. Because of that, in multiple situations plan won't be empty after importing and manual state operations may be required.
