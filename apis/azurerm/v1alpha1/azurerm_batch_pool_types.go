package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermBatchPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermBatchPoolSpec   `json:"spec,omitempty"`
	Status            AzurermBatchPoolStatus `json:"status,omitempty"`
}

type AzurermBatchPoolSpecAutoScale struct {
	EvaluationInterval string `json:"evaluation_interval"`
	Formula            string `json:"formula"`
}

type AzurermBatchPoolSpecContainerConfiguration struct {
	Type string `json:"type"`
}

type AzurermBatchPoolSpecStorageImageReference struct {
	Id        string `json:"id"`
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type AzurermBatchPoolSpecStartTaskResourceFile struct {
	FileMode                 string `json:"file_mode"`
	FilePath                 string `json:"file_path"`
	HttpUrl                  string `json:"http_url"`
	StorageContainerUrl      string `json:"storage_container_url"`
	AutoStorageContainerName string `json:"auto_storage_container_name"`
	BlobPrefix               string `json:"blob_prefix"`
}

type AzurermBatchPoolSpecStartTaskUserIdentityAutoUser struct {
	ElevationLevel string `json:"elevation_level"`
	Scope          string `json:"scope"`
}

type AzurermBatchPoolSpecStartTaskUserIdentity struct {
	UserName string                                      `json:"user_name"`
	AutoUser []AzurermBatchPoolSpecStartTaskUserIdentity `json:"auto_user"`
}

type AzurermBatchPoolSpecStartTask struct {
	ResourceFile      []AzurermBatchPoolSpecStartTask `json:"resource_file"`
	CommandLine       string                          `json:"command_line"`
	MaxTaskRetryCount int                             `json:"max_task_retry_count"`
	WaitForSuccess    bool                            `json:"wait_for_success"`
	Environment       map[string]string               `json:"environment"`
	UserIdentity      []AzurermBatchPoolSpecStartTask `json:"user_identity"`
}

type AzurermBatchPoolSpecFixedScale struct {
	TargetLowPriorityNodes int    `json:"target_low_priority_nodes"`
	ResizeTimeout          string `json:"resize_timeout"`
	TargetDedicatedNodes   int    `json:"target_dedicated_nodes"`
}

type AzurermBatchPoolSpecCertificate struct {
	Id            string   `json:"id"`
	StoreLocation string   `json:"store_location"`
	StoreName     string   `json:"store_name"`
	Visibility    []string `json:"visibility"`
}

type AzurermBatchPoolSpec struct {
	StopPendingResizeOperation bool                   `json:"stop_pending_resize_operation"`
	Name                       string                 `json:"name"`
	ResourceGroupName          string                 `json:"resource_group_name"`
	AutoScale                  []AzurermBatchPoolSpec `json:"auto_scale"`
	NodeAgentSkuId             string                 `json:"node_agent_sku_id"`
	AccountName                string                 `json:"account_name"`
	ContainerConfiguration     []AzurermBatchPoolSpec `json:"container_configuration"`
	StorageImageReference      []AzurermBatchPoolSpec `json:"storage_image_reference"`
	DisplayName                string                 `json:"display_name"`
	MaxTasksPerNode            int                    `json:"max_tasks_per_node"`
	StartTask                  []AzurermBatchPoolSpec `json:"start_task"`
	VmSize                     string                 `json:"vm_size"`
	FixedScale                 []AzurermBatchPoolSpec `json:"fixed_scale"`
	Certificate                []AzurermBatchPoolSpec `json:"certificate"`
}

type AzurermBatchPoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermBatchPoolList is a list of AzurermBatchPools
type AzurermBatchPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermBatchPool CRD objects
	Items []AzurermBatchPool `json:"items,omitempty"`
}
