---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zitadel_domain_policy Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  
---

# zitadel_domain_policy (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String) Id for the organization
- `smtp_sender_address_matches_instance_domain` (Boolean)
- `user_login_must_be_domain` (Boolean) User login must be domain
- `validate_org_domains` (Boolean) Validate organization domains

### Read-Only

- `id` (String) The ID of this resource.
- `is_default` (Boolean) Is this policy the default

