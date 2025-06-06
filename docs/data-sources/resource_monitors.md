---
page_title: "snowflake_resource_monitors Data Source - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  Data source used to get details of filtered resource monitors. Filtering is aligned with the current possibilities for SHOW RESOURCE MONITORS https://docs.snowflake.com/en/sql-reference/sql/show-resource-monitors query (like is supported). The results of SHOW is encapsulated in show_output collection.
---

# snowflake_resource_monitors (Data Source)

Data source used to get details of filtered resource monitors. Filtering is aligned with the current possibilities for [SHOW RESOURCE MONITORS](https://docs.snowflake.com/en/sql-reference/sql/show-resource-monitors) query (`like` is supported). The results of SHOW is encapsulated in show_output collection.

## Example Usage

```terraform
# Simple usage
data "snowflake_resource_monitors" "simple" {
}

output "simple_output" {
  value = data.snowflake_resource_monitors.simple.resource_monitors
}

# Filtering (like)
data "snowflake_resource_monitors" "like" {
  like = "resource-monitor-name"
}

output "like_output" {
  value = data.snowflake_resource_monitors.like.resource_monitors
}

# Ensure the number of resource monitors is equal to at least one element (with the use of postcondition)
data "snowflake_resource_monitors" "assert_with_postcondition" {
  like = "resource-monitor-name-%"
  lifecycle {
    postcondition {
      condition     = length(self.resource_monitors) > 0
      error_message = "there should be at least one resource monitor"
    }
  }
}

# Ensure the number of resource monitors is equal to exactly one element (with the use of check block)
check "resource_monitor_check" {
  data "snowflake_resource_monitors" "assert_with_check_block" {
    like = "resource-monitor-name"
  }

  assert {
    condition     = length(data.snowflake_resource_monitors.assert_with_check_block.resource_monitors) == 1
    error_message = "Resource monitors filtered by '${data.snowflake_resource_monitors.assert_with_check_block.like}' returned ${length(data.snowflake_resource_monitors.assert_with_check_block.resource_monitors)} resource monitors where one was expected"
  }
}
```

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `like` (String) Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).

### Read-Only

- `id` (String) The ID of this resource.
- `resource_monitors` (List of Object) Holds the aggregated output of all resource monitor details queries. (see [below for nested schema](#nestedatt--resource_monitors))

<a id="nestedatt--resource_monitors"></a>
### Nested Schema for `resource_monitors`

Read-Only:

- `show_output` (List of Object) (see [below for nested schema](#nestedobjatt--resource_monitors--show_output))

<a id="nestedobjatt--resource_monitors--show_output"></a>
### Nested Schema for `resource_monitors.show_output`

Read-Only:

- `comment` (String)
- `created_on` (String)
- `credit_quota` (Number)
- `end_time` (String)
- `frequency` (String)
- `level` (String)
- `name` (String)
- `owner` (String)
- `remaining_credits` (Number)
- `start_time` (String)
- `suspend_at` (Number)
- `suspend_immediate_at` (Number)
- `used_credits` (Number)
