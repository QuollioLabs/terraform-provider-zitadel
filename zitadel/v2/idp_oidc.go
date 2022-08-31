package v2

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/idp"
	management2 "github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management"
)

const (
	idpOrgIDVar           = "org_id"
	idpNameVar            = "name"
	idpStylingTypeVar     = "styling_type"
	idpClientIDVar        = "client_id"
	idpClientSecretVar    = "client_secret"
	idpIssuerVar          = "issuer"
	idpScopesVar          = "scopes"
	idpDisplayNameMapping = "display_name_mapping"
	idpUsernameMapping    = "username_mapping"
	idpAutoRegister       = "auto_register"
)

func GetOrgOIDCIDP() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a OIDC IDP of the organization.",
		Schema: map[string]*schema.Schema{
			idpOrgIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the organization",
				ForceNew:    true,
			},
			idpNameVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the IDP",
			},
			idpStylingTypeVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Some identity providers specify the styling of the button to their login",
			},
			idpClientIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "client id generated by the identity provider",
			},
			idpClientSecretVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "client secret generated by the identity provider",
			},
			idpIssuerVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "the oidc issuer of the identity provider",
			},
			idpScopesVar: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "the scopes requested by ZITADEL during the request on the identity provider",
			},
			idpDisplayNameMapping: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "definition which field is mapped to the display name of the user",
			},
			idpUsernameMapping: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "definition which field is mapped to the email of the user",
			},
			idpAutoRegister: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "auto register for users from this idp",
			},
		},
		ReadContext:   readOrgOIDCIDP,
		UpdateContext: updateOrgOIDCIDP,
		CreateContext: createOrgOIDCIDP,
		DeleteContext: deleteOrgIDP,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func deleteOrgIDP(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started delete")

	clientinfo, ok := m.(*ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	client, err := getManagementClient(clientinfo, d.Get(idpOrgIDVar).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.RemoveOrgIDP(ctx, &management2.RemoveOrgIDPRequest{
		IdpId: d.Id(),
	})
	if err != nil {
		return diag.Errorf("failed to delete oidc idp: %v", err)
	}
	return nil
}

func createOrgOIDCIDP(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started create")

	clientinfo, ok := m.(*ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	client, err := getManagementClient(clientinfo, d.Get(idpOrgIDVar).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	scopes := make([]string, 0)
	scopesSet := d.Get(idpScopesVar).(*schema.Set)
	for _, scope := range scopesSet.List() {
		scopes = append(scopes, scope.(string))
	}

	stylingType := d.Get(idpStylingTypeVar)
	displayNameMapping := d.Get(idpDisplayNameMapping).(string)
	usernameMapping := d.Get(idpUsernameMapping).(string)
	resp, err := client.AddOrgOIDCIDP(ctx, &management2.AddOrgOIDCIDPRequest{
		Name:               d.Get(idpNameVar).(string),
		StylingType:        idp.IDPStylingType(idp.IDPStylingType_value[stylingType.(string)]),
		ClientId:           d.Get(idpClientIDVar).(string),
		ClientSecret:       d.Get(idpClientSecretVar).(string),
		Issuer:             d.Get(idpIssuerVar).(string),
		Scopes:             scopes,
		DisplayNameMapping: idp.OIDCMappingField(idp.OIDCMappingField_value[displayNameMapping]),
		UsernameMapping:    idp.OIDCMappingField(idp.OIDCMappingField_value[usernameMapping]),
		AutoRegister:       d.Get(idpAutoRegister).(bool),
	})
	if err != nil {
		return diag.Errorf("failed to create oidc idp: %v", err)
	}
	d.SetId(resp.GetIdpId())

	return nil
}

func updateOrgOIDCIDP(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started update")

	clientinfo, ok := m.(*ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	client, err := getManagementClient(clientinfo, d.Get(idpOrgIDVar).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := client.GetOrgIDPByID(ctx, &management2.GetOrgIDPByIDRequest{Id: d.Id()})
	if err != nil {
		return diag.Errorf("failed to read oidc idp: %v", err)
	}

	idpID := d.Id()
	name := d.Get(idpNameVar).(string)
	stylingType := d.Get(idpStylingTypeVar).(string)
	autoRegister := d.Get(idpAutoRegister).(bool)
	changed := false
	if resp.GetIdp().GetName() != name ||
		resp.GetIdp().GetStylingType().String() != stylingType ||
		resp.GetIdp().GetAutoRegister() != autoRegister {
		changed = true
		_, err := client.UpdateOrgIDP(ctx, &management2.UpdateOrgIDPRequest{
			IdpId:        idpID,
			Name:         name,
			StylingType:  idp.IDPStylingType(idp.IDPStylingType_value[stylingType]),
			AutoRegister: autoRegister,
		})
		if err != nil {
			return diag.Errorf("failed to update oidc idp: %v", err)
		}
	}

	oidc := resp.GetIdp().GetOidcConfig()
	clientID := d.Get(idpClientIDVar).(string)
	clientSecret := d.Get(idpClientSecretVar).(string)
	issuer := d.Get(idpIssuerVar).(string)
	scopesSet := d.Get(idpScopesVar).(*schema.Set)
	displayNameMapping := d.Get(idpDisplayNameMapping).(string)
	usernameMapping := d.Get(idpUsernameMapping).(string)

	scopes := make([]string, 0)
	for _, scope := range scopesSet.List() {
		scopes = append(scopes, scope.(string))
	}

	//either nothing changed on the IDP or something besides the secret changed
	if (oidc.GetClientId() != clientID ||
		oidc.GetIssuer() != issuer ||
		!reflect.DeepEqual(oidc.GetScopes(), scopes) ||
		oidc.GetDisplayNameMapping().String() != displayNameMapping ||
		oidc.GetUsernameMapping().String() != usernameMapping) ||
		!changed {

		_, err = client.UpdateOrgIDPOIDCConfig(ctx, &management2.UpdateOrgIDPOIDCConfigRequest{
			IdpId:              idpID,
			ClientId:           clientID,
			ClientSecret:       clientSecret,
			Issuer:             issuer,
			Scopes:             scopes,
			DisplayNameMapping: idp.OIDCMappingField(idp.OIDCMappingField_value[displayNameMapping]),
			UsernameMapping:    idp.OIDCMappingField(idp.OIDCMappingField_value[usernameMapping]),
		})
		if err != nil {
			return diag.Errorf("failed to update oidc idp config: %v", err)
		}
	}
	d.SetId(idpID)
	return nil
}

func readOrgOIDCIDP(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started read")

	clientinfo, ok := m.(*ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	client, err := getManagementClient(clientinfo, d.Get(idpOrgIDVar).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := client.GetOrgIDPByID(ctx, &management2.GetOrgIDPByIDRequest{Id: d.Id()})
	if err != nil {
		d.SetId("")
		return nil
		//return diag.Errorf("failed to read oidc idp: %v", err)
	}

	idp := resp.GetIdp()
	oidc := idp.GetOidcConfig()
	set := map[string]interface{}{
		idpOrgIDVar:           idp.GetDetails().GetResourceOwner(),
		idpNameVar:            idp.GetName(),
		idpStylingTypeVar:     idp.GetStylingType().String(),
		idpClientIDVar:        oidc.GetClientId(),
		idpClientSecretVar:    d.Get(idpClientSecretVar).(string),
		idpIssuerVar:          oidc.GetIssuer(),
		idpScopesVar:          oidc.GetScopes(),
		idpDisplayNameMapping: oidc.GetDisplayNameMapping().String(),
		idpUsernameMapping:    oidc.GetUsernameMapping().String(),
		idpAutoRegister:       idp.GetAutoRegister(),
	}
	for k, v := range set {
		if err := d.Set(k, v); err != nil {
			return diag.Errorf("failed to set %s of oidc idp: %v", k, err)
		}
	}
	d.SetId(idp.Id)

	return nil
}
