---
page_title: "zitadel_application_api Data Source - terraform-provider-zitadel"
subcategory: ""
description: |-
  Datasource representing an API application belonging to a project, with all configuration possibilities.
---

# zitadel_application_api (Data Source)

Datasource representing an API application belonging to a project, with all configuration possibilities.

## Example Usage

```terraform
data zitadel_application_api api_application {
  org_id     = data.zitadel_org.org.id
  project_id = data.zitadel_project.project.id
  app_id     = "177073625566806019"
}

output api_application {
  value = data.zitadel_application_api.api_application
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID of this resource.
- `org_id` (String) orgID of the application
- `project_id` (String) ID of the project

### Read-Only

- `auth_method_type` (String) Auth method type
- `id` (String) The ID of this resource.
- `name` (String) Name of the application