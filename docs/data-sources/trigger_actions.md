---
page_title: "zitadel_trigger_actions Data Source - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing triggers, when actions get started
---

# zitadel_trigger_actions (Data Source)

Resource representing triggers, when actions get started

## Example Usage

```terraform
data zitadel_trigger_actions trigger_actions {
  org_id       = data.zitadel_org.org.id
  flow_type    = "FLOW_TYPE_EXTERNAL_AUTHENTICATION"
  trigger_type = "TRIGGER_TYPE_POST_AUTHENTICATION"
}

output trigger_actions {
  value = data.zitadel_trigger_actions.trigger_actions
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `flow_type` (String) Type of the flow to which the action triggers belong
- `org_id` (String) ID of the organization
- `trigger_type` (String) Trigger type on when the actions get triggered

### Read-Only

- `action_ids` (Set of String) IDs of the triggered actions
- `id` (String) The ID of this resource.