package linode

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/linode/linodego"
	"github.com/linode/linodego/pkg/condition"
)

const (
	linodeLKECreateTimeout = 15 * time.Minute
	linodeLKEUpdateTimeout = 20 * time.Minute
	linodeLKEDeleteTimeout = 10 * time.Minute
)

func resourceLinodeLKECluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeLKEClusterCreate,
		Read:   resourceLinodeLKEClusterRead,
		Update: resourceLinodeLKEClusterUpdate,
		Delete: resourceLinodeLKEClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(linodeLKECreateTimeout),
			Update: schema.DefaultTimeout(linodeLKEUpdateTimeout),
			Delete: schema.DefaultTimeout(linodeLKEDeleteTimeout),
		},
		Schema: map[string]*schema.Schema{
			"label": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique label for the cluster.",
			},
			"k8s_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "An array of tags applied to this object. Tags are for organizational purposes only.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This cluster's location.",
			},
			"api_endpoints": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "The API endpoints for the cluster.",
			},
			"kubeconfig": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The Base64-encoded Kubeconfig for the cluster.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the cluster.",
			},
			"pool": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the Node Pool.",
						},
						"count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntAtLeast(1),
							Description:  "The number of nodes in the Node Pool.",
							Required:     true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "A Linode Type for all of the nodes in the Node Pool.",
							Required:    true,
						},
						"nodes": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "The ID of the node.",
										Computed:    true,
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Description: "The ID of the underlying Linode instance.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: `The status of the node.`,
										Computed:    true,
									},
								},
							},
							Computed:    true,
							Description: "The nodes in the node pool.",
						},
					},
				},
				MinItems:    1,
				Required:    true,
				Description: "A node pool in the cluster.",
			},
		},
	}
}

func resourceLinodeLKEClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Linode LKE Cluster ID: %s", err)
	}

	cluster, err := client.GetLKECluster(context.Background(), id)
	if err != nil {
		return fmt.Errorf("failed to get LKE cluster %d: %s", id, err)
	}

	pools, err := client.ListLKEClusterPools(context.Background(), id, nil)
	if err != nil {
		return fmt.Errorf("failed to get pools for LKE cluster %d: %s", id, err)
	}

	kubeconfig, err := client.GetLKEClusterKubeconfig(context.Background(), id)
	if err != nil {
		return fmt.Errorf("failed to get kubeconfig for LKE cluster %d: %s", id, err)
	}

	endpoints, err := client.ListLKEClusterAPIEndpoints(context.Background(), id, nil)
	if err != nil {
		return fmt.Errorf("failed to get API endpoints for LKE cluster %d: %s", id, err)
	}

	d.Set("label", cluster.Label)
	d.Set("k8s_version", cluster.K8sVersion)
	d.Set("region", cluster.Region)
	d.Set("tags", cluster.Tags)
	d.Set("status", cluster.Status)
	d.Set("kubeconfig", kubeconfig.KubeConfig)
	d.Set("pool", flattenLinodeLKEClusterPools(pools))
	d.Set("api_endpoints", flattenLinodeLKEClusterAPIEndpoints(endpoints))
	return nil
}

func resourceLinodeLKEClusterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	createOpts := linodego.LKEClusterCreateOptions{
		Label:      d.Get("label").(string),
		Region:     d.Get("region").(string),
		K8sVersion: d.Get("k8s_version").(string),
	}

	for _, nodePool := range d.Get("pool").([]interface{}) {
		poolSpec := nodePool.(map[string]interface{})
		createOpts.NodePools = append(createOpts.NodePools, linodego.LKEClusterPoolCreateOptions{
			Type:  poolSpec["type"].(string),
			Count: poolSpec["count"].(int),
		})
	}

	if tagsRaw, tagsOk := d.GetOk("tags"); tagsOk {
		for _, tag := range tagsRaw.(*schema.Set).List() {
			createOpts.Tags = append(createOpts.Tags, tag.(string))
		}
	}

	cluster, err := client.CreateLKECluster(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("failed to create LKE cluster: %s", err)
	}
	d.SetId(strconv.Itoa(cluster.ID))

	client.WaitForLKEClusterConditions(context.Background(), cluster.ID, linodego.LKEClusterPollOptions{
		TimeoutSeconds: 10 * 60,
	}, condition.ClusterHasReadyNode)
	return resourceLinodeLKEClusterRead(d, meta)
}

func resourceLinodeLKEClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("failed parsing Linode LKE Cluster ID: %s", err)
	}

	updateOpts := linodego.LKEClusterUpdateOptions{}
	if d.HasChange("label") {
		updateOpts.Label = d.Get("label").(string)
	}
	if d.HasChange("tags") {
		tags := []string{}
		for _, tag := range d.Get("tags").(*schema.Set).List() {
			tags = append(tags, tag.(string))
		}

		updateOpts.Tags = &tags
	}
	if d.HasChanges("label", "tags") {
		if _, err := client.UpdateLKECluster(context.Background(), id, updateOpts); err != nil {
			return fmt.Errorf("failed to update LKE Cluster %d: %s", id, err)
		}
	}

	poolSpecs := getLinodeLKEClusterPoolSpecs(d.Get("pool"))
	pools, err := client.ListLKEClusterPools(context.Background(), id, nil)
	if err != nil {
		return fmt.Errorf("failed to get node pools for LKE Cluster %d: %s", id, err)
	}

	// map pool specs to provisioned clusters
	provisionedPools := map[linodeLKEClusterPoolSpec][]int{}
	for _, pool := range pools {
		spec := linodeLKEClusterPoolSpec{pool.Type, pool.Count}
		provisionedPools[spec] = append(provisionedPools[spec], pool.ID)
	}

	// keep track of all specs visited for accounting
	visitedSpecs := make(map[linodeLKEClusterPoolSpec]struct{})

	toDelete := []int{}
	for spec, count := range poolSpecs {
		diff := 0
		ids, ok := provisionedPools[spec]
		if !ok {
			diff = count
		} else {
			diff = count - len(ids)
		}

		if diff > 0 {
			createOpts := linodego.LKEClusterPoolCreateOptions{
				Count: spec.Count,
				Type:  spec.Type,
			}
			// stage cluster pools for creation
			for i := 0; i < diff; i++ {
				if _, err := client.CreateLKEClusterPool(context.Background(), id, createOpts); err != nil {
					return fmt.Errorf("failed to create node pool for cluster %d: %s", id, err)
				}
			}
		} else if diff < 0 {
			// stage cluster pools for deletion
			deleteCount := int(math.Abs(float64(diff)))
			toDelete = append(toDelete, ids[:deleteCount]...)
		}

		visitedSpecs[spec] = struct{}{}
	}

	// ensure there are no provisioned cluster pools for which there are no
	// declared specifications
	for spec, ids := range provisionedPools {
		if _, ok := visitedSpecs[spec]; !ok {
			// stage these cluster pools for deletion
			toDelete = append(toDelete, ids...)
		}
	}

	for _, poolID := range toDelete {
		if err := client.DeleteLKEClusterPool(context.Background(), id, poolID); err != nil {
			return fmt.Errorf("failed to delete node pool for cluster %d: %s", id, err)
		}
	}

	return resourceLinodeLKEClusterRead(d, meta)
}

func resourceLinodeLKEClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("failed parsing Linode LKE Cluster ID: %s", err)
	}

	err = client.DeleteLKECluster(context.Background(), id)
	if err != nil {
		return fmt.Errorf("failed to delete Linode LKE cluster %d: %s", id, err)
	}
	client.WaitForLKEClusterStatus(context.Background(), id, "not_ready", int(d.Timeout(schema.TimeoutCreate).Seconds()))
	return nil
}

type linodeLKEClusterPoolSpec struct {
	Type  string
	Count int
}

func getLinodeLKEClusterPoolSpecs(pool interface{}) map[linodeLKEClusterPoolSpec]int {
	specs := pool.([]interface{})
	poolSpecs := map[linodeLKEClusterPoolSpec]int{}
	for _, spec := range specs {
		specMap := spec.(map[string]interface{})
		poolSpecs[linodeLKEClusterPoolSpec{
			Type:  specMap["type"].(string),
			Count: specMap["count"].(int),
		}]++
	}
	return poolSpecs
}

func flattenLinodeLKEClusterAPIEndpoints(apiEndpoints []linodego.LKEClusterAPIEndpoint) []string {
	flattened := make([]string, len(apiEndpoints))
	for i, endpoint := range apiEndpoints {
		flattened[i] = endpoint.Endpoint
	}
	return flattened
}

func flattenLinodeLKEClusterPools(pools []linodego.LKEClusterPool) []map[string]interface{} {
	flattened := make([]map[string]interface{}, len(pools))
	for i, pool := range pools {

		nodes := make([]map[string]interface{}, len(pool.Linodes))
		for i, node := range pool.Linodes {
			nodes[i] = map[string]interface{}{
				"id":          node.ID,
				"instance_id": node.InstanceID,
				"status":      node.Status,
			}
		}

		flattened[i] = map[string]interface{}{
			"id":    pool.ID,
			"count": pool.Count,
			"type":  pool.Type,
			"nodes": nodes,
		}
	}
	return flattened
}
