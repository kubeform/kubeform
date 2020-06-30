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

func resourceLinodeNodeBalancerConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeNodeBalancerConfigCreate,
		Read:   resourceLinodeNodeBalancerConfigRead,
		Update: resourceLinodeNodeBalancerConfigUpdate,
		Delete: resourceLinodeNodeBalancerConfigDelete,
		Importer: &schema.ResourceImporter{
			State: resourceLinodeNodeBalancerConfigImport,
		},
		Schema: map[string]*schema.Schema{
			"nodebalancer_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the NodeBalancer to access.",
				Required:    true,
				ForceNew:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				Description:  "The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key.",
				ValidateFunc: validation.StringInSlice([]string{"http", "https", "tcp"}, false),
				Optional:     true,
				Default:      linodego.ProtocolHTTP,
			},
			"port": {
				Type:         schema.TypeInt,
				Description:  "The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.",
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Default:      80,
			},
			"check_interval": {
				Type:        schema.TypeInt,
				Description: "How often, in seconds, to check that backends are up and serving requests.",
				Optional:    true,
				Computed:    true,
			},
			"check_timeout": {
				Type:         schema.TypeInt,
				Description:  "How long, in seconds, to wait for a check attempt before considering it failed. (1-30)",
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"check_attempts": {
				Type:         schema.TypeInt,
				Description:  "How many times to attempt a check before considering a backend to be down. (1-30)",
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"algorithm": {
				Type:         schema.TypeString,
				Description:  "What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source",
				ValidateFunc: validation.StringInSlice([]string{"roundrobin", "leastconn", "source"}, false),
				Optional:     true,
				Computed:     true,
			},
			"stickiness": {
				Type:         schema.TypeString,
				Description:  "Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'",
				ValidateFunc: validation.StringInSlice([]string{"none", "table", "http_cookie"}, false),
				Optional:     true,
				Computed:     true,
			},
			"check": {
				Type:         schema.TypeString,
				Description:  "The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.",
				ValidateFunc: validation.StringInSlice([]string{"none", "connection", "http", "http_body"}, false),
				Optional:     true,
				Computed:     true,
			},
			"check_path": {
				Type:        schema.TypeString,
				Description: "The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.",
				Optional:    true,
				Computed:    true,
			},
			"check_body": {
				Type:        schema.TypeString,
				Description: "This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down",
				Optional:    true,
				Computed:    true,
			},
			"check_passive": {
				Type:        schema.TypeBool,
				Description: "If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.",
				Optional:    true,
				Computed:    true,
			},
			"cipher_suite": {
				Type:         schema.TypeString,
				Description:  "What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.",
				ValidateFunc: validation.StringInSlice([]string{"recommended", "legacy"}, false),
				Optional:     true,
				Computed:     true,
			},
			"ssl_commonname": {
				Type:        schema.TypeString,
				Description: "The common name for the SSL certification this port is serving if this port is not configured to use SSL.",
				Computed:    true,
			},
			"ssl_fingerprint": {
				Type:        schema.TypeString,
				Description: "The fingerprint for the SSL certification this port is serving if this port is not configured to use SSL.",
				Computed:    true,
			},
			"ssl_cert": {
				Type:        schema.TypeString,
				Description: "The certificate this port is serving. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.",
				Optional:    true,
			},
			"ssl_key": {
				Type:        schema.TypeString,
				Description: "The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.",
				Optional:    true,
				Sensitive:   true,
			},
			"node_status": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_up": {
							Type:        schema.TypeInt,
							Description: "The number of backends considered to be 'UP' and healthy, and that are serving requests.",
							Computed:    true,
						},
						"status_down": {
							Type:        schema.TypeInt,
							Description: "The number of backends considered to be 'DOWN' and unhealthy. These are not in rotation, and not serving requests.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceLinodeNodeBalancerConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if strings.Contains(d.Id(), ",") {
		s := strings.Split(d.Id(), ",")
		// Validate that this is an ID by making sure it can be converted into an int
		_, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, fmt.Errorf("invalid nodebalancer_config ID: %v", err)
		}

		nodebalancerID, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, fmt.Errorf("invalid nodebalancer ID: %v", err)
		}

		d.SetId(s[1])
		d.Set("nodebalancer_id", nodebalancerID)
	}

	err := resourceLinodeNodeBalancerConfigRead(d, meta)
	if err != nil {
		return nil, fmt.Errorf("unable to import %v as nodebalancer_config: %v", d.Id(), err)
	}

	results := make([]*schema.ResourceData, 0)
	results = append(results, d)

	return results, nil
}

func resourceLinodeNodeBalancerConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode NodeBalancerConfig ID %s as int: %s", d.Id(), err)
	}
	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("nodebalancer_id"))
	}

	config, err := client.GetNodeBalancerConfig(context.Background(), int(nodebalancerID), int(id))
	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			log.Printf("[WARN] removing NodeBalancer Config ID %q from state because it no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding the specified Linode NodeBalancerConfig: %s", err)
	}

	d.Set("algorithm", config.Algorithm)
	d.Set("stickiness", config.Stickiness)
	d.Set("check", config.Check)
	d.Set("check_attempts", config.CheckAttempts)
	d.Set("check_body", config.CheckBody)
	d.Set("check_interval", config.CheckInterval)
	d.Set("check_timeout", config.CheckTimeout)
	d.Set("check_passive", config.CheckPassive)
	d.Set("check_path", config.CheckPath)
	d.Set("cipher_suite", config.CipherSuite)
	d.Set("port", config.Port)
	d.Set("protocol", config.Protocol)
	d.Set("ssl_key", config.SSLKey)
	d.Set("ssl_fingerprint", config.SSLFingerprint)
	d.Set("ssl_commonname", config.SSLCommonName)

	nodeStatus := map[string]interface{}{
		"up":   fmt.Sprintf("%d", config.NodesStatus.Up),
		"down": fmt.Sprintf("%d", config.NodesStatus.Down),
	}
	d.Set("node_status", nodeStatus)

	return nil
}

func resourceLinodeNodeBalancerConfigCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode NodeBalancerConfig")
	}

	nodebalancerID := d.Get("nodebalancer_id").(int)

	createOpts := linodego.NodeBalancerConfigCreateOptions{
		Algorithm:     linodego.ConfigAlgorithm(d.Get("algorithm").(string)),
		Check:         linodego.ConfigCheck(d.Get("check").(string)),
		Stickiness:    linodego.ConfigStickiness(d.Get("stickiness").(string)),
		CheckAttempts: d.Get("check_attempts").(int),
		CheckBody:     d.Get("check_body").(string),
		CheckInterval: d.Get("check_interval").(int),
		CheckPath:     d.Get("check_path").(string),
		CheckTimeout:  d.Get("check_timeout").(int),
		Port:          d.Get("port").(int),
		Protocol:      linodego.ConfigProtocol(d.Get("protocol").(string)),
		SSLCert:       d.Get("ssl_cert").(string),
		SSLKey:        d.Get("ssl_key").(string),
	}

	if checkPassiveRaw, ok := d.GetOkExists("check_passive"); ok {
		checkPassive := checkPassiveRaw.(bool)
		createOpts.CheckPassive = &checkPassive
	}

	config, err := client.CreateNodeBalancerConfig(context.Background(), nodebalancerID, createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode NodeBalancerConfig: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", config.ID))
	d.Set("nodebalancer_id", nodebalancerID)

	return resourceLinodeNodeBalancerConfigRead(d, meta)
}

func resourceLinodeNodeBalancerConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode NodeBalancerConfig ID %s as int: %s", d.Id(), err)
	}
	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %s as int", d.Get("nodebalancer_id"))
	}

	updateOpts := linodego.NodeBalancerConfigUpdateOptions{
		Algorithm:     linodego.ConfigAlgorithm(d.Get("algorithm").(string)),
		Check:         linodego.ConfigCheck(d.Get("check").(string)),
		Stickiness:    linodego.ConfigStickiness(d.Get("stickiness").(string)),
		CheckAttempts: d.Get("check_attempts").(int),
		CheckBody:     d.Get("check_body").(string),
		CheckInterval: d.Get("check_interval").(int),
		CheckPath:     d.Get("check_path").(string),
		CheckTimeout:  d.Get("check_timeout").(int),
		Port:          d.Get("port").(int),
		Protocol:      linodego.ConfigProtocol(d.Get("protocol").(string)),
		SSLCert:       d.Get("ssl_cert").(string),
		SSLKey:        d.Get("ssl_key").(string),
	}

	if ok := d.HasChange("check_passive"); ok {
		checkPassive := d.Get("check_passive").(bool)
		updateOpts.CheckPassive = &checkPassive
	}

	if _, err = client.UpdateNodeBalancerConfig(context.Background(), int(nodebalancerID), int(id), updateOpts); err != nil {
		return fmt.Errorf("Error updating Nodebalancer %d Config %d: %s", int(nodebalancerID), int(id), err)
	}

	return resourceLinodeNodeBalancerConfigRead(d, meta)
}

func resourceLinodeNodeBalancerConfigDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode NodeBalancerConfig ID %s as int: %s", d.Id(), err)
	}
	nodebalancerID, ok := d.Get("nodebalancer_id").(int)
	if !ok {
		return fmt.Errorf("Error parsing Linode NodeBalancer ID %v as int", d.Get("nodebalancer_id"))
	}
	err = client.DeleteNodeBalancerConfig(context.Background(), nodebalancerID, int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode NodeBalancerConfig %d: %s", id, err)
	}
	return nil
}
