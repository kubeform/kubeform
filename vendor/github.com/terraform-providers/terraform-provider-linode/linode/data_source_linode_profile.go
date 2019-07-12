package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeProfileRead,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The profile email address. This address will be used for communication with Linode as necessary.",
			},
			"timezone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The profile's preferred timezone. This is not used by the API, and is for the benefit of clients only. All times the API returns are in UTC.",
			},
			"email_notifications": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, email notifications will be sent about account activity. If false, when false business-critical communications may still be sent through email.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The username for logging in to Linode services.",
			},
			"ip_whitelist_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, logins for the user will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.",
			},
			"lish_auth_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The methods of authentication allowed when connecting via Lish. 'keys_only' is the most secure with the intent to use Lish, and 'disabled' is recommended for users that will not use Lish at all.",
			},
			"authorized_keys": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "The list of SSH Keys authorized to use Lish for this user. This value is ignored if lish_auth_method is 'disabled'.",
			},
			"two_factor_auth": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, logins from untrusted computers will require Two Factor Authentication.",
			},
			"restricted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, the user has restrictions on what can be accessed on the Account.",
			},
			"referrals": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Credit Card information associated with this Account.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total": {
							Type:        schema.TypeInt,
							Description: "The number of users who have signed up with the referral code.",
							Computed:    true,
						},
						"credit": {
							Type:        schema.TypeFloat,
							Description: "The amount of account credit in US Dollars issued to the account through the referral program.",
							Computed:    true,
						},
						"completed": {
							Type:        schema.TypeInt,
							Description: "The number of completed signups with the referral code.",
							Computed:    true,
						},
						"pending": {
							Type:        schema.TypeInt,
							Description: "The number of pending signups for the referral code. To receive credit the signups must be completed.",
							Computed:    true,
						},
						"code": {
							Type:        schema.TypeString,
							Description: "The Profile referral code.  If new accounts use this when signing up for Linode, referring account will receive credit.",
							Computed:    true,
						},
						"url": {
							Type:        schema.TypeString,
							Description: "The referral URL.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLinodeProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	profile, err := client.GetProfile(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting profile: %s", err)
	}

	d.SetId(fmt.Sprintf("%d", profile.UID))
	if err := d.Set("referrals", flattenProfileReferrals(profile.Referrals)); err != nil {
		return fmt.Errorf("Error parsing profile referrals: %s", err)
	}

	d.Set("email", profile.Email)
	d.Set("timezone", profile.Timezone)
	d.Set("email_notifications", profile.EmailNotifications)
	d.Set("username", profile.Username)
	d.Set("ip_whitelist_enabled", profile.IPWhitelistEnabled)
	d.Set("lish_auth_method", profile.LishAuthMethod)
	if err := d.Set("authorized_keys", profile.AuthorizedKeys); err != nil {
		return fmt.Errorf("Error parsing authorized_keys: %s", err)
	}
	d.Set("two_factor_auth", profile.TwoFactorAuth)
	d.Set("restricted", profile.Restricted)

	return nil
}
