package linode

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeDomainRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeDomainRecordRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Record.",
				Optional:    true,
			},
			"domain_id": {
				Type:        schema.TypeInt,
				Description: "The associated domain's ID.",
				Required:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "The type of Record this is in the DNS system.",
				Computed:    true,
			},
			"ttl_sec": {
				Type:        schema.TypeInt,
				Description: "The amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.",
				Computed:    true,
			},
			"target": {
				Type:        schema.TypeString,
				Description: "The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "The priority of the target host. Lower values are preferred.",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "The relative weight of this Record. Higher values are preferred.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "The port this Record points to.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "The protocol this Record's service communicates with. Only valid for SRV records.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeString,
				Description: "The service this Record identified. Only valid for SRV records.",
				Computed:    true,
			},
			"tag": {
				Type:        schema.TypeString,
				Description: "The tag portion of a CAA record.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeDomainRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	domainID := d.Get("domain_id").(int)
	recordName := d.Get("name").(string)
	recordID := d.Get("id").(int)

	if recordName == "" && recordID == 0 {
		return fmt.Errorf("Record name or ID is required")
	}

	var record *linodego.DomainRecord

	if recordID != 0 {
		rec, err := client.GetDomainRecord(context.Background(), domainID, recordID)
		if err != nil {
			return fmt.Errorf("Error fetching domain record: %v", err)
		}
		record = rec
	} else if recordName != "" {
		filter, _ := json.Marshal(map[string]interface{}{"name": recordName})
		records, err := client.ListDomainRecords(context.Background(), domainID, linodego.NewListOptions(0, string(filter)))
		if err != nil {
			return fmt.Errorf("Error listing domain records: %v", err)
		}
		if len(records) > 0 {
			record = &records[0]
		}
	}

	if record != nil {
		d.SetId(strconv.Itoa(recordID))
		d.Set("id", record.ID)
		d.Set("name", record.Name)
		d.Set("type", record.Type)
		d.Set("ttl_sec", record.TTLSec)
		d.Set("target", record.Target)
		d.Set("priority", record.Priority)
		d.Set("protocol", record.Protocol)
		d.Set("weight", record.Weight)
		d.Set("port", record.Port)
		d.Set("service", record.Service)
		d.Set("tag", record.Tag)
		return nil
	}

	d.SetId("")

	return fmt.Errorf(`Domain record "%s" for domain %d was not found`, recordName, domainID)
}
