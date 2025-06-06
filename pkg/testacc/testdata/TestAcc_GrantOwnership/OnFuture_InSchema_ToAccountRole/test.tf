resource "snowflake_account_role" "test" {
  name = var.account_role_name
}

resource "snowflake_schema" "test" {
  name     = var.schema_name
  database = var.database_name
}

resource "snowflake_grant_ownership" "test" {
  account_role_name = snowflake_account_role.test.name
  on {
    future {
      object_type_plural = "TABLES"
      in_schema          = "\"${var.database_name}\".\"${snowflake_schema.test.name}\""
    }
  }
}
