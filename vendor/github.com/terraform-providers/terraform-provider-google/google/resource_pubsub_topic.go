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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePubsubTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubTopicCreate,
		Read:   resourcePubsubTopicRead,
		Update: resourcePubsubTopicUpdate,
		Delete: resourcePubsubTopicDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePubsubTopicImport,
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
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"kms_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"message_storage_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_persistence_regions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
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

func resourcePubsubTopicCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandPubsubTopicName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	kmsKeyNameProp, err := expandPubsubTopicKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key_name"); !isEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	labelsProp, err := expandPubsubTopicLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	messageStoragePolicyProp, err := expandPubsubTopicMessageStoragePolicy(d.Get("message_storage_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("message_storage_policy"); !isEmptyValue(reflect.ValueOf(messageStoragePolicyProp)) && (ok || !reflect.DeepEqual(v, messageStoragePolicyProp)) {
		obj["messageStoragePolicy"] = messageStoragePolicyProp
	}

	obj, err = resourcePubsubTopicEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/topics/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Topic: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutCreate), pubsubTopicProjectNotReady)
	if err != nil {
		return fmt.Errorf("Error creating Topic: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/topics/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Topic %q: %#v", d.Id(), res)

	return resourcePubsubTopicRead(d, meta)
}

func resourcePubsubTopicRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/topics/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("PubsubTopic %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Topic: %s", err)
	}

	if err := d.Set("name", flattenPubsubTopicName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Topic: %s", err)
	}
	if err := d.Set("kms_key_name", flattenPubsubTopicKmsKeyName(res["kmsKeyName"], d)); err != nil {
		return fmt.Errorf("Error reading Topic: %s", err)
	}
	if err := d.Set("labels", flattenPubsubTopicLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Topic: %s", err)
	}
	if err := d.Set("message_storage_policy", flattenPubsubTopicMessageStoragePolicy(res["messageStoragePolicy"], d)); err != nil {
		return fmt.Errorf("Error reading Topic: %s", err)
	}

	return nil
}

func resourcePubsubTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandPubsubTopicLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	messageStoragePolicyProp, err := expandPubsubTopicMessageStoragePolicy(d.Get("message_storage_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("message_storage_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, messageStoragePolicyProp)) {
		obj["messageStoragePolicy"] = messageStoragePolicyProp
	}

	obj, err = resourcePubsubTopicUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/topics/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Topic %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("message_storage_policy") {
		updateMask = append(updateMask, "messageStoragePolicy")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Topic %q: %s", d.Id(), err)
	}

	return resourcePubsubTopicRead(d, meta)
}

func resourcePubsubTopicDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/topics/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Topic %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Topic")
	}

	log.Printf("[DEBUG] Finished deleting Topic %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubTopicImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/topics/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/topics/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubTopicName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenPubsubTopicKmsKeyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubTopicLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenPubsubTopicMessageStoragePolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["allowed_persistence_regions"] =
		flattenPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(original["allowedPersistenceRegions"], d)
	return []interface{}{transformed}
}
func flattenPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandPubsubTopicName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func expandPubsubTopicKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubTopicMessageStoragePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedPersistenceRegions, err := expandPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(original["allowed_persistence_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedPersistenceRegions); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedPersistenceRegions"] = transformedAllowedPersistenceRegions
	}

	return transformed, nil
}

func expandPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourcePubsubTopicEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func resourcePubsubTopicUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["topic"] = obj
	return newObj, nil
}