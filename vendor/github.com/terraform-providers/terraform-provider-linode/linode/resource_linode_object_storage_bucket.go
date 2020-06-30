package linode

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func resourceLinodeObjectStorageBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeObjectStorageBucketCreate,
		Read:   resourceLinodeObjectStorageBucketRead,
		Delete: resourceLinodeObjectStorageBucketDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeString,
				Description: "The cluster of the Linode Object Storage Bucket.",
				Required:    true,
				ForceNew:    true,
			},
			"label": {
				Type:        schema.TypeString,
				Description: "The label of the Linode Object Storage Bucket.",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceLinodeObjectStorageBucketRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	cluster, label, err := decodeLinodeObjectStorageBucketID(d.Id())

	if err != nil {
		return fmt.Errorf("Error parsing Linode ObjectStorageBucket id %s", d.Id())
	}

	bucket, err := client.GetObjectStorageBucket(context.Background(), cluster, label)

	if err != nil {
		return fmt.Errorf("Error finding the specified Linode ObjectStorageBucket: %s", err)
	}

	d.SetId(fmt.Sprintf("%s:%s", bucket.Cluster, bucket.Label))
	d.Set("cluster", bucket.Cluster)
	d.Set("label", bucket.Label)

	return nil
}

func resourceLinodeObjectStorageBucketCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode ObjectStorageBucket")
	}

	createOpts := linodego.ObjectStorageBucketCreateOptions{
		Cluster: d.Get("cluster").(string),
		Label:   d.Get("label").(string),
	}
	bucket, err := client.CreateObjectStorageBucket(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode ObjectStorageBucket: %s", err)
	}
	d.SetId(fmt.Sprintf("%s:%s", bucket.Cluster, bucket.Label))
	d.Set("cluster", bucket.Cluster)
	d.Set("label", bucket.Label)

	return resourceLinodeObjectStorageBucketRead(d, meta)
}

func resourceLinodeObjectStorageBucketDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	cluster, label, err := decodeLinodeObjectStorageBucketID(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Linode ObjectStorageBucket id %s", d.Id())
	}
	err = client.DeleteObjectStorageBucket(context.Background(), cluster, label)
	if err != nil {
		return fmt.Errorf("Error deleting Linode ObjectStorageBucket %s: %s", d.Id(), err)
	}
	return nil
}

func decodeLinodeObjectStorageBucketID(id string) (cluster, label string, err error) {
	parts := strings.Split(id, ":")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		err = fmt.Errorf("Linode Object Storage Bucket ID must be of the form <Cluster>:<Label>, was provided: %s", id)
		return
	}
	cluster = parts[0]
	label = parts[1]
	return
}
