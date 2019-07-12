package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/linode/linodego"
)

func resourceLinodeRDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeRDNSCreate,
		Read:   resourceLinodeRDNSRead,
		Delete: resourceLinodeRDNSDelete,
		Exists: resourceLinodeRDNSExists,
		Update: resourceLinodeRDNSUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"address": {
				Type:         schema.TypeString,
				Description:  "The public Linode IPv4 or IPv6 address to operate on.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.SingleIP(),
			},
			"rdns": {
				Type:         schema.TypeString,
				Description:  "The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(3, 254),
			},
		},
	}
}

func resourceLinodeRDNSExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(linodego.Client)

	ipStr := d.Id()

	if len(ipStr) == 0 {
		return false, fmt.Errorf("Error parsing Linode RDNS ID as IP string")
	}

	_, err := client.GetIPAddress(context.Background(), ipStr)

	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			d.SetId("")
			return false, nil
		}

		return false, fmt.Errorf("Error getting Linode RDNS for %s: %s", ipStr, err)
	}
	return true, nil
}

func resourceLinodeRDNSRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	ipStr := d.Id()

	if len(ipStr) == 0 {
		return fmt.Errorf("Error parsing Linode RDNS ID %s as IP string", ipStr)
	}

	ip, err := client.GetIPAddress(context.Background(), ipStr)

	if err != nil {
		return fmt.Errorf("Error finding the specified Linode RDNS: %s", err)
	}

	d.Set("address", d.Id())
	d.Set("rdns", ip.RDNS)

	return nil
}

func resourceLinodeRDNSCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode RDNS")
	}

	var address = d.Get("address").(string)
	var rdns *string
	if rdnsRaw, ok := d.GetOk("rdns"); ok && len(rdnsRaw.(string)) > 0 {
		rdnsStr := rdnsRaw.(string)
		rdns = &rdnsStr
	}
	updateOpts := linodego.IPAddressUpdateOptions{
		RDNS: rdns,
	}
	ip, err := client.UpdateIPAddress(context.Background(), address, updateOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode RDNS: %s", err)
	}
	d.SetId(address)
	d.Set("rdns", ip.RDNS)

	return resourceLinodeRDNSRead(d, meta)
}

func resourceLinodeRDNSUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	ipStr := d.Id()

	if len(ipStr) == 0 {
		return fmt.Errorf("Error parsing Linode RDNS ID %s as IP string", ipStr)
	}

	var rdns *string

	if rdnsRaw, ok := d.GetOk("rdns"); ok && len(rdnsRaw.(string)) > 0 {
		rdnsStr := rdnsRaw.(string)
		rdns = &rdnsStr
	}

	updateOpts := linodego.IPAddressUpdateOptions{
		RDNS: rdns,
	}

	if _, err := client.UpdateIPAddress(context.Background(), d.Id(), updateOpts); err != nil {
		return fmt.Errorf("Error updating Linode RDNS: %s", err)
	}

	return resourceLinodeRDNSRead(d, meta)
}

func resourceLinodeRDNSDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	ipStr := d.Id()

	if len(ipStr) == 0 {
		return fmt.Errorf("Error parsing Linode RDNS ID %s as IP string", ipStr)
	}

	updateOpts := linodego.IPAddressUpdateOptions{
		RDNS: nil,
	}

	if _, err := client.UpdateIPAddress(context.Background(), d.Id(), updateOpts); err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error deleting Linode RDNS: %s", err)
	}

	d.SetId("")

	return nil
}
