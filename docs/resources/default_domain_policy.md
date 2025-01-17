---
page_title: "zitadel_default_domain_policy Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing the default domain policy.
---

# zitadel_default_domain_policy (Resource)

Resource representing the default domain policy.

## Example Usage

```terraform
resource zitadel_default_domain_policy domain_policy {
  user_login_must_be_domain                   = false
  validate_org_domains                        = false
  smtp_sender_address_matches_instance_domain = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `smtp_sender_address_matches_instance_domain` (Boolean)
- `user_login_must_be_domain` (Boolean) User login must be domain
- `validate_org_domains` (Boolean) Validate organization domains

### Read-Only

- `id` (String) The ID of this resource.