package linode

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/linode/linodego"
)

func resourceLinodeDomainRecord() *schema.Resource {
	validDomainSeconds := domainSecondsValidator()

	return &schema.Resource{
		Create: resourceLinodeDomainRecordCreate,
		Read:   resourceLinodeDomainRecordRead,
		Update: resourceLinodeDomainRecordUpdate,
		Delete: resourceLinodeDomainRecordDelete,
		Exists: resourceLinodeDomainRecordExists,
		Importer: &schema.ResourceImporter{
			State: resourceLinodeDomainRecordImport,
		},
		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the Domain to access.",
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Type:         schema.TypeString,
				Description:  "The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(0, 100),
			},
			"record_type": {
				Type:         schema.TypeString,
				Description:  "The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "NS", "MX", "CNAME", "TXT", "SRV", "PTR", "CAA"}, false),
			},
			"ttl_sec": {
				Type:         schema.TypeInt,
				Description:  "'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.",
				ValidateFunc: validDomainSeconds,
				Optional:     true,
			},
			"target": {
				Type:        schema.TypeString,
				Description: "The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.",
				Required:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				Description:  "The priority of the target host. Lower values are preferred.",
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 255),
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "The protocol this Record's service communicates with. Only valid for SRV records.",
				Optional:    true,
			},
			"service": {
				Type:        schema.TypeString,
				Description: "The service this Record identified. Only valid for SRV records.",
				Optional:    true,
			},
			"tag": {
				Type:        schema.TypeString,
				Description: "The tag portion of a CAA record. It is invalid to set this on other record types.",
				Optional:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "The port this Record points to.",
				Optional:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "The relative weight of this Record. Higher values are preferred.",
				Optional:    true,
			},
		},
	}
}

func resourceLinodeDomainRecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return false, fmt.Errorf("Error parsing Linode DomainRecord ID %s as int: %s", d.Id(), err)
	}
	if domainID, ok := d.GetOkExists("domain_id"); ok {
		_, err = client.GetDomainRecord(context.Background(), domainID.(int), int(id))
	} else {
		return false, fmt.Errorf("Error parsing Linode Domain ID")
	}

	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			return false, nil
		}

		return false, fmt.Errorf("Error getting Linode DomainRecord ID %s: %s", d.Id(), err)
	}
	return true, nil
}

func resourceLinodeDomainRecordImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if strings.Contains(d.Id(), ",") {
		s := strings.Split(d.Id(), ",")
		// Validate that this is an ID by making sure it can be converted into an int
		_, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, fmt.Errorf("invalid domain_record ID: %v", err)
		}

		domainID, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, fmt.Errorf("invalid domain ID: %v", err)
		}

		d.SetId(s[1])
		d.Set("domain_id", domainID)
	}

	err := resourceLinodeDomainRecordRead(d, meta)
	if err != nil {
		return nil, fmt.Errorf("unable to import %v as domain_record: %v", d.Id(), err)
	}

	results := make([]*schema.ResourceData, 0)
	results = append(results, d)

	return results, nil
}

func resourceLinodeDomainRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode DomainRecord ID %s as int: %s", d.Id(), err)
	}
	domainID := d.Get("domain_id").(int)
	record, err := client.GetDomainRecord(context.Background(), int(domainID), int(id))

	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			log.Printf("[WARN] removing Linode Domain Record ID %q from state because it no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding the specified Linode DomainRecord: %s", err)
	}

	d.Set("name", record.Name)
	d.Set("port", record.Port)
	d.Set("priority", record.Priority)
	d.Set("protocol", record.Protocol)
	d.Set("service", record.Service)
	d.Set("tag", record.Tag)
	d.Set("target", record.Target)
	d.Set("ttl_sec", record.TTLSec)
	d.Set("record_type", record.Type)
	d.Set("weight", record.Weight)

	return nil
}

func resourceDataStringOrNil(d *schema.ResourceData, name string) *string {
	if val, ok := d.GetOkExists(name); ok {
		i := val.(string)
		if len(i) == 0 {
			return nil
		}
		return &i
	}
	return nil
}

func resourceDataIntOrNil(d *schema.ResourceData, name string) *int {
	if val, ok := d.GetOkExists(name); ok {
		i := val.(int)
		return &i
	}
	return nil
}

func resourceLinodeDomainRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode DomainRecord")
	}
	domainID := d.Get("domain_id").(int)

	createOpts := linodego.DomainRecordCreateOptions{
		Type:     linodego.DomainRecordType(d.Get("record_type").(string)),
		Name:     d.Get("name").(string),
		Priority: resourceDataIntOrNil(d, "priority"),
		Target:   d.Get("target").(string),
		Weight:   resourceDataIntOrNil(d, "weight"),
		Port:     resourceDataIntOrNil(d, "port"),
		Service:  resourceDataStringOrNil(d, "service"),
		Protocol: resourceDataStringOrNil(d, "protocol"),
		TTLSec:   d.Get("ttl_sec").(int),
		Tag:      resourceDataStringOrNil(d, "tag"),
	}

	domainRecord, err := client.CreateDomainRecord(context.Background(), domainID, createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode DomainRecord: %s", err)
	}

	d.SetId(fmt.Sprintf("%d", domainRecord.ID))

	return resourceLinodeDomainRecordRead(d, meta)
}

func resourceLinodeDomainRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	domainID := d.Get("domain_id").(int)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode DomainRecord id %s as int: %s", d.Id(), err)
	}
	updateOpts := linodego.DomainRecordUpdateOptions{
		Type:     linodego.DomainRecordType(d.Get("record_type").(string)),
		Name:     d.Get("name").(string),
		Priority: resourceDataIntOrNil(d, "priority"),
		Target:   d.Get("target").(string),
		Weight:   resourceDataIntOrNil(d, "weight"),
		Port:     resourceDataIntOrNil(d, "port"),
		Service:  resourceDataStringOrNil(d, "service"),
		Protocol: resourceDataStringOrNil(d, "protocol"),
		TTLSec:   d.Get("ttl_sec").(int),
		Tag:      resourceDataStringOrNil(d, "tag"),
	}

	_, err = client.UpdateDomainRecord(context.Background(), domainID, int(id), updateOpts)
	if err != nil {
		return fmt.Errorf("Error updating Domain Record: %s", err)
	}

	return resourceLinodeDomainRecordRead(d, meta)
}

func resourceLinodeDomainRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	domainID := d.Get("domain_id").(int)
	id, err := strconv.ParseInt(d.Id(), 10, 64)

	if err != nil {
		return fmt.Errorf("Error parsing Linode DomainRecord id %s as int", d.Id())
	}
	err = client.DeleteDomainRecord(context.Background(), domainID, int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode DomainRecord %d: %s", id, err)
	}
	d.SetId("")

	return nil
}
