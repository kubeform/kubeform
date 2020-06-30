// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func comparePubsubSubscriptionExpirationPolicy(_, old, new string, _ *schema.ResourceData) bool {
	trimmedNew := strings.TrimLeft(new, "0")
	trimmedOld := strings.TrimLeft(old, "0")
	if strings.Contains(trimmedNew, ".") {
		trimmedNew = strings.TrimRight(strings.TrimSuffix(trimmedNew, "s"), "0") + "s"
	}
	if strings.Contains(trimmedOld, ".") {
		trimmedOld = strings.TrimRight(strings.TrimSuffix(trimmedOld, "s"), "0") + "s"
	}
	return trimmedNew == trimmedOld
}

func resourcePubsubSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubSubscriptionCreate,
		Read:   resourcePubsubSubscriptionRead,
		Update: resourcePubsubSubscriptionUpdate,
		Delete: resourcePubsubSubscriptionDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePubsubSubscriptionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: comparePubsubSubscriptionBasename,
			},
			"topic": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"ack_deadline_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"expiration_policy": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: comparePubsubSubscriptionExpirationPolicy,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"message_retention_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "604800s",
			},
			"push_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"push_endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},
						"attributes": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"retain_acked_messages": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePubsubSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandPubsubSubscriptionName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	topicProp, err := expandPubsubSubscriptionTopic(d.Get("topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topic"); !isEmptyValue(reflect.ValueOf(topicProp)) && (ok || !reflect.DeepEqual(v, topicProp)) {
		obj["topic"] = topicProp
	}
	labelsProp, err := expandPubsubSubscriptionLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pushConfigProp, err := expandPubsubSubscriptionPushConfig(d.Get("push_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("push_config"); !isEmptyValue(reflect.ValueOf(pushConfigProp)) && (ok || !reflect.DeepEqual(v, pushConfigProp)) {
		obj["pushConfig"] = pushConfigProp
	}
	ackDeadlineSecondsProp, err := expandPubsubSubscriptionAckDeadlineSeconds(d.Get("ack_deadline_seconds"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ack_deadline_seconds"); !isEmptyValue(reflect.ValueOf(ackDeadlineSecondsProp)) && (ok || !reflect.DeepEqual(v, ackDeadlineSecondsProp)) {
		obj["ackDeadlineSeconds"] = ackDeadlineSecondsProp
	}
	messageRetentionDurationProp, err := expandPubsubSubscriptionMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !isEmptyValue(reflect.ValueOf(messageRetentionDurationProp)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	retainAckedMessagesProp, err := expandPubsubSubscriptionRetainAckedMessages(d.Get("retain_acked_messages"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retain_acked_messages"); !isEmptyValue(reflect.ValueOf(retainAckedMessagesProp)) && (ok || !reflect.DeepEqual(v, retainAckedMessagesProp)) {
		obj["retainAckedMessages"] = retainAckedMessagesProp
	}
	expirationPolicyProp, err := expandPubsubSubscriptionExpirationPolicy(d.Get("expiration_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_policy"); ok || !reflect.DeepEqual(v, expirationPolicyProp) {
		obj["expirationPolicy"] = expirationPolicyProp
	}

	obj, err = resourcePubsubSubscriptionEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Subscription: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Subscription: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Subscription %q: %#v", d.Id(), res)

	return resourcePubsubSubscriptionRead(d, meta)
}

func resourcePubsubSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("PubsubSubscription %q", d.Id()))
	}

	res, err = resourcePubsubSubscriptionDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing PubsubSubscription because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}

	if err := d.Set("name", flattenPubsubSubscriptionName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("topic", flattenPubsubSubscriptionTopic(res["topic"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("labels", flattenPubsubSubscriptionLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("push_config", flattenPubsubSubscriptionPushConfig(res["pushConfig"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("ack_deadline_seconds", flattenPubsubSubscriptionAckDeadlineSeconds(res["ackDeadlineSeconds"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("message_retention_duration", flattenPubsubSubscriptionMessageRetentionDuration(res["messageRetentionDuration"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("retain_acked_messages", flattenPubsubSubscriptionRetainAckedMessages(res["retainAckedMessages"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("expiration_policy", flattenPubsubSubscriptionExpirationPolicy(res["expirationPolicy"], d)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}

	return nil
}

func resourcePubsubSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandPubsubSubscriptionLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pushConfigProp, err := expandPubsubSubscriptionPushConfig(d.Get("push_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("push_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pushConfigProp)) {
		obj["pushConfig"] = pushConfigProp
	}
	ackDeadlineSecondsProp, err := expandPubsubSubscriptionAckDeadlineSeconds(d.Get("ack_deadline_seconds"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ack_deadline_seconds"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ackDeadlineSecondsProp)) {
		obj["ackDeadlineSeconds"] = ackDeadlineSecondsProp
	}
	messageRetentionDurationProp, err := expandPubsubSubscriptionMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	retainAckedMessagesProp, err := expandPubsubSubscriptionRetainAckedMessages(d.Get("retain_acked_messages"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retain_acked_messages"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retainAckedMessagesProp)) {
		obj["retainAckedMessages"] = retainAckedMessagesProp
	}
	expirationPolicyProp, err := expandPubsubSubscriptionExpirationPolicy(d.Get("expiration_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_policy"); ok || !reflect.DeepEqual(v, expirationPolicyProp) {
		obj["expirationPolicy"] = expirationPolicyProp
	}

	obj, err = resourcePubsubSubscriptionUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Subscription %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("push_config") {
		updateMask = append(updateMask, "pushConfig")
	}

	if d.HasChange("ack_deadline_seconds") {
		updateMask = append(updateMask, "ackDeadlineSeconds")
	}

	if d.HasChange("message_retention_duration") {
		updateMask = append(updateMask, "messageRetentionDuration")
	}

	if d.HasChange("retain_acked_messages") {
		updateMask = append(updateMask, "retainAckedMessages")
	}

	if d.HasChange("expiration_policy") {
		updateMask = append(updateMask, "expirationPolicy")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Subscription %q: %s", d.Id(), err)
	}

	return resourcePubsubSubscriptionRead(d, meta)
}

func resourcePubsubSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Subscription %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Subscription")
	}

	log.Printf("[DEBUG] Finished deleting Subscription %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubSubscriptionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/subscriptions/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubSubscriptionName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenPubsubSubscriptionTopic(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenPubsubSubscriptionLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubSubscriptionPushConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["push_endpoint"] =
		flattenPubsubSubscriptionPushConfigPushEndpoint(original["pushEndpoint"], d)
	transformed["attributes"] =
		flattenPubsubSubscriptionPushConfigAttributes(original["attributes"], d)
	return []interface{}{transformed}
}
func flattenPubsubSubscriptionPushConfigPushEndpoint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubSubscriptionPushConfigAttributes(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubSubscriptionAckDeadlineSeconds(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenPubsubSubscriptionMessageRetentionDuration(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubSubscriptionRetainAckedMessages(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubSubscriptionExpirationPolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["ttl"] =
		flattenPubsubSubscriptionExpirationPolicyTtl(original["ttl"], d)
	return []interface{}{transformed}
}
func flattenPubsubSubscriptionExpirationPolicyTtl(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandPubsubSubscriptionName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return "", err
	}

	subscription := d.Get("name").(string)

	re := regexp.MustCompile("projects\\/(.*)\\/subscriptions\\/(.*)")
	match := re.FindStringSubmatch(subscription)
	if len(match) == 3 {
		// We need to preserve the behavior where the user passes the subscription name already in the long form,
		// however we need it to be stored as the short form since it's used for the replaceVars in the URL.
		// The unintuitive behavior is that if the user provides the long form, we use the project from there, not the one
		// specified on the resource or provider.
		// TODO(drebes): consider deprecating the long form behavior for 3.0
		d.Set("project", match[1])
		d.Set("name", match[2])
		return subscription, nil
	}
	return fmt.Sprintf("projects/%s/subscriptions/%s", project, subscription), nil
}

func expandPubsubSubscriptionTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return "", err
	}

	topic := d.Get("topic").(string)

	re := regexp.MustCompile("projects\\/(.*)\\/topics\\/(.*)")
	match := re.FindStringSubmatch(topic)
	if len(match) == 3 {
		return topic, nil
	} else {
		// If no full topic given, we expand it to a full topic on the same project
		fullTopic := fmt.Sprintf("projects/%s/topics/%s", project, topic)
		d.Set("topic", fullTopic)
		return fullTopic, nil
	}
}

func expandPubsubSubscriptionLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionPushConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPushEndpoint, err := expandPubsubSubscriptionPushConfigPushEndpoint(original["push_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPushEndpoint); val.IsValid() && !isEmptyValue(val) {
		transformed["pushEndpoint"] = transformedPushEndpoint
	}

	transformedAttributes, err := expandPubsubSubscriptionPushConfigAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !isEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigPushEndpoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigAttributes(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionAckDeadlineSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionMessageRetentionDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionRetainAckedMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionExpirationPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTtl, err := expandPubsubSubscriptionExpirationPolicyTtl(original["ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTtl); val.IsValid() && !isEmptyValue(val) {
		transformed["ttl"] = transformedTtl
	}

	return transformed, nil
}

func expandPubsubSubscriptionExpirationPolicyTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourcePubsubSubscriptionEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func resourcePubsubSubscriptionUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["subscription"] = obj
	return newObj, nil
}

func resourcePubsubSubscriptionDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {

	// path is a derived field from the API-side `name`
	path := fmt.Sprintf("projects/%s/subscriptions/%s", d.Get("project"), d.Get("name"))
	d.Set("path", path)

	return res, nil
}