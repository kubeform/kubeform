package linode

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/linode/linodego"
)

func resourceLinodeNodeBalancerNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeNodeBalancerNodeCreate,
		Read:   resourceLinodeNodeBalancerNodeRead,
		Update: resourceLinodeNodeBalancerNodeUpdate,
		Delete: resourceLinodeNodeBalancerNodeDelete,
		Importer: &schema.ResourceImporter{
			State: resourceLinodeNodeBalancerNodeImport,
		},
		Schema: map[string]*schema.Schema{
			"nodebalancer_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the NodeBalancer to access.",
				Required:    true,
				ForceNew:    true,
			},
			"config_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the NodeBalancerConfig to access.",
				Required:    true,
				ForceNew:    true,
			},
			"label": {
				Type:        schema.TypeString,
				Description: "The label for this node. This is for display purposes only.",
				Required:    true,
			},
			"weight": {
				Type:         schema.TypeInt,
				Description:  "Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)",
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"mode": {
				Type:         schema.TypeString,
				Description:  "The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. If set to `backup` this backend will only accept traffic if all other nodes are down.",
				ValidateFunc: validation.StringInSlice([]string{"accept", "reject", "drain", "backup"}, false),
				Optional:     true,
				Computed:     true,
			},
			"address": {
				Type:        schema.TypeString,
				Description: "The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.",
				Required:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "The current status of this node, based on the configured checks of its NodeBalancer Config. (unknown, UP, DOWN)",
				Computed:    true,
			},
		},
	}
}

func resourceLinodeNodeBalancerNodeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode NodeBalancerNode ID %s as int: %s", d.Id(), err)
	}
	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("nodebalancer_id"))
	}
	configID, ok := d.Get("config_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("config_id"))
	}

	node, err := client.GetNodeBalancerNode(context.Background(), nodebalancerID, configID, int(id))

	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			log.Printf("[WARN] removing NodeBalancer Node ID %q from state because it no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error finding the specified Linode NodeBalancerNode: %s", err)
	}

	d.Set("label", node.Label)
	d.Set("weight", node.Weight)
	d.Set("mode", node.Mode)
	d.Set("address", node.Address)
	d.Set("status", node.Status)
	return nil
}

func resourceLinodeNodeBalancerNodeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if strings.Contains(d.Id(), ",") {
		s := strings.Split(d.Id(), ",")
		// Validate that this is an ID by making sure it can be converted into an int
		_, err := strconv.Atoi(s[2])
		if err != nil {
			return nil, fmt.Errorf("invalid nodebalancer_node ID: %v", err)
		}

		configID, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, fmt.Errorf("invalid config ID: %v", err)
		}

		nodebalancerID, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, fmt.Errorf("invalid nodebalancer ID: %v", err)
		}

		d.SetId(s[2])
		d.Set("nodebalancer_id", nodebalancerID)
		d.Set("config_id", configID)
	}

	err := resourceLinodeNodeBalancerNodeRead(d, meta)
	if err != nil {
		return nil, fmt.Errorf("unable to import %v as nodebalancer_node: %v", d.Id(), err)
	}

	results := make([]*schema.ResourceData, 0)
	results = append(results, d)

	return results, nil
}

func resourceLinodeNodeBalancerNodeCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode NodeBalancerNode")
	}

	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("nodebalancer_id"))
	}
	configID, ok := d.Get("config_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("config_id"))
	}

	createOpts := linodego.NodeBalancerNodeCreateOptions{
		Address: d.Get("address").(string),
		Label:   d.Get("label").(string),
		Mode:    linodego.NodeMode(d.Get("mode").(string)),
		Weight:  d.Get("weight").(int),
	}
	node, err := client.CreateNodeBalancerNode(context.Background(), int(nodebalancerID), int(configID), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode NodeBalancerNode: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", node.ID))
	d.Set("config_id", configID)
	d.Set("nodebalancer_id", nodebalancerID)

	return resourceLinodeNodeBalancerNodeRead(d, meta)
}

func resourceLinodeNodeBalancerNodeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode NodeBalancerConfig ID %v as int: %s", d.Id(), err)
	}
	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("nodebalancer_id"))
	}
	configID, ok := d.Get("config_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("config_id"))
	}

	updateOpts := linodego.NodeBalancerNodeUpdateOptions{
		Address: d.Get("address").(string),
		Label:   d.Get("label").(string),
		Mode:    linodego.NodeMode(d.Get("mode").(string)),
		Weight:  d.Get("weight").(int),
	}

	if _, err = client.UpdateNodeBalancerNode(context.Background(), nodebalancerID, configID, int(id), updateOpts); err != nil {
		return fmt.Errorf("Error updating Linode Nodebalancer %d Config %d Node %d: %s", nodebalancerID, configID, int(id), err)
	}

	return resourceLinodeNodeBalancerNodeRead(d, meta)
}

func resourceLinodeNodeBalancerNodeDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode NodeBalancerConfig ID %s as int: %s", d.Id(), err)
	}
	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("nodebalancer_id"))
	}
	configID, ok := d.Get("config_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("config_id"))
	}
	err = client.DeleteNodeBalancerNode(context.Background(), nodebalancerID, configID, int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode NodeBalancerNode %d: %s", id, err)
	}
	return nil
}
