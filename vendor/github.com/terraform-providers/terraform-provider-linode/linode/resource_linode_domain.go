package linode

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/linode/linodego"
)

func resourceLinodeDomain() *schema.Resource {
	validDomainSeconds := domainSecondsValidator()

	return &schema.Resource{
		Create: resourceLinodeDomainCreate,
		Read:   resourceLinodeDomainRead,
		Update: resourceLinodeDomainUpdate,
		Delete: resourceLinodeDomainDelete,
		Exists: resourceLinodeDomainExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Description: "The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.",
				Required:    true,
			},
			"type": {
				Type:         schema.TypeString,
				Description:  "If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).",
				InputDefault: "master",
				ValidateFunc: validation.StringInSlice([]string{"master", "slave"}, false),
				Required:     true,
				ForceNew:     true,
			},
			"group": {
				Type:         schema.TypeString,
				Description:  "The group this Domain belongs to. This is for display purposes only.",
				ValidateFunc: validation.StringLenBetween(0, 50),
				Optional:     true,
			},
			"status": {
				Type:         schema.TypeString,
				Description:  "Used to control whether this Domain is currently being rendered.",
				Optional:     true,
				Computed:     true,
				InputDefault: "active",
			},
			"description": {
				Type:         schema.TypeString,
				Description:  "A description for this Domain. This is for display purposes only.",
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"master_ips": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The IP addresses representing the master DNS for this Domain.",
				Optional:    true,
			},
			"axfr_ips": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.",
				Optional:    true,
			},
			"ttl_sec": {
				Type:         schema.TypeInt,
				Description:  "'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.",
				ValidateFunc: validDomainSeconds,
				Optional:     true,
			},
			"retry_sec": {
				Type:         schema.TypeInt,
				Description:  "The interval, in seconds, at which a failed refresh should be retried. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.",
				ValidateFunc: validDomainSeconds,
				Optional:     true,
			},
			"expire_sec": {
				Type:         schema.TypeInt,
				Description:  "The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 0, 00, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.",
				ValidateFunc: validDomainSeconds,
				Optional:     true,
			},
			"refresh_sec": {
				Type:         schema.TypeInt,
				Description:  "The amount of time in seconds before this Domain should be refreshed. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.",
				ValidateFunc: validDomainSeconds,
				Optional:     true,
			},
			"soa_email": {
				Type:        schema.TypeString,
				Description: "Start of Authority email address. This is required for master Domains.",
				Optional:    true,
			},
			"tags": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "An array of tags applied to this object. Tags are for organizational purposes only.",
			},
		},
	}
}

// IntInSlice returns a SchemaValidateFunc which tests if the provided value
// is of type int and matches the value of an element in the valid slice
func intInSlice(valid []int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(int)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be int", k))
			return
		}

		for _, n := range valid {
			if v == n {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %d", k, valid, v))
		return
	}
}

func domainSecondsValidator() schema.SchemaValidateFunc {
	validSeconds := []int{0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, 2419200}
	return intInSlice(validSeconds)
}

func resourceLinodeDomainExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return false, fmt.Errorf("Error parsing Linode Domain ID %s as int: %s", d.Id(), err)
	}

	_, err = client.GetDomain(context.Background(), int(id))
	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			return false, nil
		}

		return false, fmt.Errorf("Error getting Linode Domain ID %s: %s", d.Id(), err)
	}
	return true, nil
}

func resourceLinodeDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Domain ID %s as int: %s", d.Id(), err)
	}

	domain, err := client.GetDomain(context.Background(), int(id))

	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			log.Printf("[WARN] removing Linode Domain ID %q from state because it no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding the specified Linode Domain: %s", err)
	}

	d.Set("domain", domain.Domain)
	d.Set("type", domain.Type)
	d.Set("group", domain.Group)
	d.Set("status", domain.Status)
	d.Set("description", domain.Description)
	if err = d.Set("master_ips", domain.MasterIPs); err != nil {
		return fmt.Errorf("Error setting master_ips: %s", err)
	}
	if len(domain.AXfrIPs) > 0 {
		if err = d.Set("afxr_ips", domain.AXfrIPs); err != nil {
			return fmt.Errorf("Error setting axfr_ips: %s", err)
		}
	}
	d.Set("ttl_sec", domain.TTLSec)
	d.Set("retry_sec", domain.RetrySec)
	d.Set("expire_sec", domain.ExpireSec)
	d.Set("refresh_sec", domain.RefreshSec)
	d.Set("soa_email", domain.SOAEmail)
	d.Set("tags", domain.Tags)

	return nil
}

func resourceLinodeDomainCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode Domain")
	}

	createOpts := linodego.DomainCreateOptions{
		Domain:      d.Get("domain").(string),
		Type:        linodego.DomainType(d.Get("type").(string)),
		Group:       d.Get("group").(string),
		Description: d.Get("description").(string),
		SOAEmail:    d.Get("soa_email").(string),
		RetrySec:    d.Get("retry_sec").(int),
		ExpireSec:   d.Get("expire_sec").(int),
		RefreshSec:  d.Get("refresh_sec").(int),
		TTLSec:      d.Get("ttl_sec").(int),
	}

	if tagsRaw, tagsOk := d.GetOk("tags"); tagsOk {
		for _, tag := range tagsRaw.(*schema.Set).List() {
			createOpts.Tags = append(createOpts.Tags, tag.(string))
		}
	}

	if v, ok := d.GetOk("master_ips"); ok {
		var masterIPS []string
		for _, ip := range v.([]interface{}) {
			masterIPS = append(masterIPS, ip.(string))
		}

		createOpts.MasterIPs = masterIPS
	}

	if v, ok := d.GetOk("axfr_ips"); ok {
		var AXfrIPs []string
		for _, ip := range v.([]interface{}) {
			AXfrIPs = append(AXfrIPs, ip.(string))
		}

		createOpts.AXfrIPs = AXfrIPs
	}

	domain, err := client.CreateDomain(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode Domain: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", domain.ID))

	return resourceLinodeDomainRead(d, meta)
}

func resourceLinodeDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Domain id %s as int: %s", d.Id(), err)
	}

	updateOpts := linodego.DomainUpdateOptions{
		Domain:      d.Get("domain").(string),
		Status:      linodego.DomainStatus(d.Get("status").(string)),
		Group:       d.Get("group").(string),
		Description: d.Get("description").(string),
		SOAEmail:    d.Get("soa_email").(string),
		RetrySec:    d.Get("retry_sec").(int),
		ExpireSec:   d.Get("expire_sec").(int),
		RefreshSec:  d.Get("refresh_sec").(int),
		TTLSec:      d.Get("ttl_sec").(int),
	}

	if v, ok := d.GetOk("master_ips"); ok {
		var masterIPS []string
		for _, ip := range v.([]interface{}) {
			masterIPS = append(masterIPS, ip.(string))
		}

		updateOpts.MasterIPs = masterIPS
	}

	if v, ok := d.GetOk("axfr_ips"); ok {
		var AXfrIPs []string
		for _, ip := range v.([]interface{}) {
			AXfrIPs = append(AXfrIPs, ip.(string))
		}

		updateOpts.AXfrIPs = AXfrIPs
	}

	if d.HasChange("tags") {
		tags := []string{}
		for _, tag := range d.Get("tags").(*schema.Set).List() {
			tags = append(tags, tag.(string))
		}

		updateOpts.Tags = tags
	}

	_, err = client.UpdateDomain(context.Background(), int(id), updateOpts)
	if err != nil {
		return fmt.Errorf("Error updating Linode Domain %d: %s", id, err)
	}
	return resourceLinodeDomainRead(d, meta)
}

func resourceLinodeDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Domain id %s as int", d.Id())
	}
	err = client.DeleteDomain(context.Background(), int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode Domain %d: %s", id, err)
	}
	d.SetId("")

	return nil
}
