package digitalocean

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/digitalocean/godo"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDigitalOceanDatabaseUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceDigitalOceanDatabaseUserCreate,
		Read:   resourceDigitalOceanDatabaseUserRead,
		Update: resourceDigitalOceanDatabaseUserUpdate,
		Delete: resourceDigitalOceanDatabaseUserDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDigitalOceanDatabaseUserImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"cluster_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"mysql_auth_plugin": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					godo.SQLAuthPluginNative,
					godo.SQLAuthPluginCachingSHA2,
				}, false),
				// Prevent diffs when default is used and not specificed in the config.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old == godo.SQLAuthPluginCachingSHA2 && new == ""
				},
			},

			// Computed Properties
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceDigitalOceanDatabaseUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CombinedConfig).godoClient()
	clusterID := d.Get("cluster_id").(string)

	opts := &godo.DatabaseCreateUserRequest{
		Name: d.Get("name").(string),
	}

	if v, ok := d.GetOk("mysql_auth_plugin"); ok {
		opts.MySQLSettings = &godo.DatabaseMySQLUserSettings{
			AuthPlugin: v.(string),
		}
	}

	log.Printf("[DEBUG] Database User create configuration: %#v", opts)
	user, _, err := client.Databases.CreateUser(context.Background(), clusterID, opts)
	if err != nil {
		return fmt.Errorf("Error creating Database User: %s", err)
	}

	d.SetId(makeDatabaseUserID(clusterID, user.Name))
	log.Printf("[INFO] Database User Name: %s", user.Name)

	return resourceDigitalOceanDatabaseUserRead(d, meta)
}

func resourceDigitalOceanDatabaseUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CombinedConfig).godoClient()
	clusterID := d.Get("cluster_id").(string)
	name := d.Get("name").(string)

	// Check if the database user still exists
	user, resp, err := client.Databases.GetUser(context.Background(), clusterID, name)
	if err != nil {
		// If the database user is somehow already destroyed, mark as
		// successfully gone
		if resp != nil && resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving Database User: %s", err)
	}

	d.Set("role", user.Role)
	d.Set("password", user.Password)
	if user.MySQLSettings != nil {
		d.Set("mysql_auth_plugin", user.MySQLSettings.AuthPlugin)
	}

	return nil
}

func resourceDigitalOceanDatabaseUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CombinedConfig).godoClient()

	if d.HasChange("mysql_auth_plugin") {
		authReq := &godo.DatabaseResetUserAuthRequest{}
		if d.Get("mysql_auth_plugin").(string) != "" {
			authReq.MySQLSettings = &godo.DatabaseMySQLUserSettings{
				AuthPlugin: d.Get("mysql_auth_plugin").(string),
			}
		} else {
			// If blank, restore default value.
			authReq.MySQLSettings = &godo.DatabaseMySQLUserSettings{
				AuthPlugin: godo.SQLAuthPluginCachingSHA2,
			}
		}

		_, _, err := client.Databases.ResetUserAuth(context.Background(), d.Get("cluster_id").(string), d.Get("name").(string), authReq)
		if err != nil {
			return fmt.Errorf("Error updating mysql_auth_plugin for DatabaseUser: %s", err)
		}
	}

	return resourceDigitalOceanDatabaseUserRead(d, meta)
}

func resourceDigitalOceanDatabaseUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CombinedConfig).godoClient()
	clusterID := d.Get("cluster_id").(string)
	name := d.Get("name").(string)

	log.Printf("[INFO] Deleting Database User: %s", d.Id())
	_, err := client.Databases.DeleteUser(context.Background(), clusterID, name)
	if err != nil {
		return fmt.Errorf("Error deleting Database User: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceDigitalOceanDatabaseUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if strings.Contains(d.Id(), ",") {
		s := strings.Split(d.Id(), ",")
		d.SetId(makeDatabaseUserID(s[0], s[1]))
		d.Set("cluster_id", s[0])
		d.Set("name", s[1])
	}

	return []*schema.ResourceData{d}, nil
}

func makeDatabaseUserID(clusterID string, name string) string {
	return fmt.Sprintf("%s/user/%s", clusterID, name)
}
