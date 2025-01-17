package org_idp_oidc

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetDatasource() *schema.Resource {
	return &schema.Resource{
		Description: "Datasource representing a OIDC IDP of the organization.",
		Schema: map[string]*schema.Schema{
			idpIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of this resource.",
			},
			orgIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the organization",
			},
			nameVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the IDP",
			},
			stylingTypeVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Some identity providers specify the styling of the button to their login",
			},
			clientIDVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "client id generated by the identity provider",
				Sensitive:   true,
			},
			clientSecretVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "client secret generated by the identity provider",
				Sensitive:   true,
			},
			issuerVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the oidc issuer of the identity provider",
			},
			scopesVar: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "the scopes requested by ZITADEL during the request on the identity provider",
			},
			displayNameMappingVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "definition which field is mapped to the display name of the user",
			},
			usernameMappingVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "definition which field is mapped to the email of the user",
			},
			autoRegisterVar: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "auto register for users from this idp",
			},
		},
		ReadContext: read,
		Importer:    &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}
