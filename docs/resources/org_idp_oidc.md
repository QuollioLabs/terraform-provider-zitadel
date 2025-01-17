---
page_title: "zitadel_org_idp_oidc Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing a OIDC IDP of the organization.
---

# zitadel_org_idp_oidc (Resource)

Resource representing a OIDC IDP of the organization.

## Example Usage

```terraform
resource zitadel_org_idp_oidc oidc_idp {
  org_id               = zitadel_org.org.id
  name                 = "oidcidp"
  styling_type         = "STYLING_TYPE_UNSPECIFIED"
  client_id            = "google"
  client_secret        = "google_secret"
  issuer               = "https://google.com"
  scopes               = ["openid", "profile", "email"]
  display_name_mapping = "OIDC_MAPPING_FIELD_PREFERRED_USERNAME"
  username_mapping     = "OIDC_MAPPING_FIELD_PREFERRED_USERNAME"
  auto_register        = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auto_register` (Boolean) auto register for users from this idp
- `client_id` (String, Sensitive) client id generated by the identity provider
- `client_secret` (String, Sensitive) client secret generated by the identity provider
- `display_name_mapping` (String) definition which field is mapped to the display name of the user
- `issuer` (String) the oidc issuer of the identity provider
- `name` (String) Name of the IDP
- `org_id` (String) ID of the organization
- `scopes` (Set of String) the scopes requested by ZITADEL during the request on the identity provider
- `styling_type` (String) Some identity providers specify the styling of the button to their login, supported values: STYLING_TYPE_UNSPECIFIED, STYLING_TYPE_GOOGLE
- `username_mapping` (String) definition which field is mapped to the email of the user

### Read-Only

- `id` (String) The ID of this resource.