---
page_title: "snowflake_grants Data Source - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  
---

# snowflake_grants (Data Source)



## Example Usage

```terraform
##################################
### SHOW GRANTS ON ...
##################################

# account
data "snowflake_grants" "example_on_account" {
  grants_on {
    account = true
  }
}

# account object (e.g. database)
data "snowflake_grants" "example_on_account_object" {
  grants_on {
    object_name = "some_database"
    object_type = "DATABASE"
  }
}

# database object (e.g. schema)
data "snowflake_grants" "example_on_database_object" {
  grants_on {
    object_name = "\"some_database\".\"some_schema\""
    object_type = "SCHEMA"
  }
}

# schema object (e.g. table)
data "snowflake_grants" "example_on_schema_object" {
  grants_on {
    object_name = "\"some_database\".\"some_schema\".\"some_table\""
    object_type = "TABLE"
  }
}

##################################
### SHOW GRANTS TO ...
##################################

# application
data "snowflake_grants" "example_to_application" {
  grants_to {
    application = "some_application"
  }
}

# application role
data "snowflake_grants" "example_to_application_role" {
  grants_to {
    application_role = "\"some_application\".\"some_application_role\""
  }
}

# account role
data "snowflake_grants" "example_to_role" {
  grants_to {
    account_role = "some_role"
  }
}

# database role
data "snowflake_grants" "example_to_database_role" {
  grants_to {
    database_role = "\"some_database\".\"some_database_role\""
  }
}

# share
data "snowflake_grants" "example_to_share" {
  grants_to {
    share {
      share_name = "some_share"
    }
  }
}

# user
data "snowflake_grants" "example_to_user" {
  grants_to {
    user = "some_user"
  }
}

##################################
### SHOW GRANTS OF ...
##################################

# application role
data "snowflake_grants" "example_of_application_role" {
  grants_of {
    application_role = "\"some_application\".\"some_application_role\""
  }
}

# database role
data "snowflake_grants" "example_of_database_role" {
  grants_of {
    database_role = "\"some_database\".\"some_database_role\""
  }
}

# account role
data "snowflake_grants" "example_of_role" {
  grants_of {
    account_role = "some_role"
  }
}

# share
data "snowflake_grants" "example_of_share" {
  grants_of {
    share = "some_share"
  }
}

##################################
### SHOW FUTURE GRANTS IN ...
##################################

# database
data "snowflake_grants" "example_future_in_database" {
  future_grants_in {
    database = "some_database"
  }
}

# schema
data "snowflake_grants" "example_future_in_schema" {
  future_grants_in {
    schema = "\"some_database\".\"some_schema\""
  }
}

##################################
### SHOW FUTURE GRANTS TO ...
##################################

# account role
data "snowflake_grants" "example_future_to_role" {
  future_grants_to {
    account_role = "some_role"
  }
}

# database role
data "snowflake_grants" "example_future_to_database_role" {
  future_grants_to {
    database_role = "\"some_database\".\"some_database_role\""
  }
}
```

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `future_grants_in` (Block List, Max: 1) Lists all privileges on new (i.e. future) objects. (see [below for nested schema](#nestedblock--future_grants_in))
- `future_grants_to` (Block List, Max: 1) Lists all privileges granted to the object on new (i.e. future) objects. (see [below for nested schema](#nestedblock--future_grants_to))
- `grants_of` (Block List, Max: 1) Lists all objects to which the given object has been granted. (see [below for nested schema](#nestedblock--grants_of))
- `grants_on` (Block List, Max: 1) Lists all privileges that have been granted on an object or on an account. (see [below for nested schema](#nestedblock--grants_on))
- `grants_to` (Block List, Max: 1) Lists all privileges granted to the object. (see [below for nested schema](#nestedblock--grants_to))

### Read-Only

- `grants` (List of Object) The list of grants (see [below for nested schema](#nestedatt--grants))
- `id` (String) The ID of this resource.

<a id="nestedblock--future_grants_in"></a>
### Nested Schema for `future_grants_in`

Optional:

- `database` (String) Lists all privileges on new (i.e. future) objects of a specified type in the database granted to a role.
- `schema` (String) Lists all privileges on new (i.e. future) objects of a specified type in the schema granted to a role. Schema must be a fully qualified name ("&lt;db_name&gt;"."&lt;schema_name&gt;").


<a id="nestedblock--future_grants_to"></a>
### Nested Schema for `future_grants_to`

Optional:

- `account_role` (String) Lists all privileges on new (i.e. future) objects of a specified type in a database or schema granted to the account role.
- `database_role` (String) Lists all privileges on new (i.e. future) objects granted to the database role. Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").


<a id="nestedblock--grants_of"></a>
### Nested Schema for `grants_of`

Optional:

- `account_role` (String) Lists all users and roles to which the account role has been granted.
- `application_role` (String) Lists all the users and roles to which the application role has been granted. Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
- `database_role` (String) Lists all users and roles to which the database role has been granted. Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
- `share` (String) Lists all the accounts for the share and indicates the accounts that are using the share.


<a id="nestedblock--grants_on"></a>
### Nested Schema for `grants_on`

Optional:

- `account` (Boolean) Object hierarchy to list privileges on. The only valid value is: ACCOUNT. Setting this attribute lists all the account-level (i.e. global) privileges that have been granted to roles.
- `object_name` (String) Name of object to list privileges on.
- `object_type` (String) Type of object to list privileges on.


<a id="nestedblock--grants_to"></a>
### Nested Schema for `grants_to`

Optional:

- `account_role` (String) Lists all privileges and roles granted to the role.
- `application` (String) Lists all the privileges and roles granted to the application.
- `application_role` (String) Lists all the privileges and roles granted to the application role. Must be a fully qualified name ("&lt;app_name&gt;"."&lt;app_role_name&gt;").
- `database_role` (String) Lists all privileges and roles granted to the database role. Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
- `share` (Block List, Max: 1) Lists all the privileges granted to the share. (see [below for nested schema](#nestedblock--grants_to--share))
- `user` (String) Lists all the roles granted to the user. Note that the PUBLIC role, which is automatically available to every user, is not listed.

<a id="nestedblock--grants_to--share"></a>
### Nested Schema for `grants_to.share`

Required:

- `share_name` (String) Lists all of the privileges and roles granted to the specified share.



<a id="nestedatt--grants"></a>
### Nested Schema for `grants`

Read-Only:

- `created_on` (String)
- `grant_option` (Boolean)
- `granted_by` (String)
- `granted_on` (String)
- `granted_to` (String)
- `grantee_name` (String)
- `name` (String)
- `privilege` (String)
