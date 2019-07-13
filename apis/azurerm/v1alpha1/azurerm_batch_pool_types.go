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

type AzurermBatchPoolSpecStorageImageReference struct {
	Id        string `json:"id"`
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type AzurermBatchPoolSpecAutoScale struct {
	EvaluationInterval string `json:"evaluation_interval"`
	Formula            string `json:"formula"`
}

type AzurermBatchPoolSpecCertificate struct {
	Id            string   `json:"id"`
	StoreLocation string   `json:"store_location"`
	StoreName     string   `json:"store_name"`
	Visibility    []string `json:"visibility"`
}

type AzurermBatchPoolSpecStartTaskUserIdentityAutoUser struct {
	ElevationLevel string `json:"elevation_level"`
	Scope          string `json:"scope"`
}

type AzurermBatchPoolSpecStartTaskUserIdentity struct {
	UserName string                                      `json:"user_name"`
	AutoUser []AzurermBatchPoolSpecStartTaskUserIdentity `json:"auto_user"`
}

type AzurermBatchPoolSpecStartTaskResourceFile struct {
	HttpUrl                  string `json:"http_url"`
	StorageContainerUrl      string `json:"storage_container_url"`
	AutoStorageContainerName string `json:"auto_storage_container_name"`
	BlobPrefix               string `json:"blob_prefix"`
	FileMode                 string `json:"file_mode"`
	FilePath                 string `json:"file_path"`
}

type AzurermBatchPoolSpecStartTask struct {
	Environment       map[string]string               `json:"environment"`
	UserIdentity      []AzurermBatchPoolSpecStartTask `json:"user_identity"`
	ResourceFile      []AzurermBatchPoolSpecStartTask `json:"resource_file"`
	CommandLine       string                          `json:"command_line"`
	MaxTaskRetryCount int                             `json:"max_task_retry_count"`
	WaitForSuccess    bool                            `json:"wait_for_success"`
}

type AzurermBatchPoolSpecFixedScale struct {
	TargetDedicatedNodes   int    `json:"target_dedicated_nodes"`
	TargetLowPriorityNodes int    `json:"target_low_priority_nodes"`
	ResizeTimeout          string `json:"resize_timeout"`
}

type AzurermBatchPoolSpecContainerConfiguration struct {
	Type string `json:"type"`
}

type AzurermBatchPoolSpec struct {
	Name                       string                 `json:"name"`
	StorageImageReference      []AzurermBatchPoolSpec `json:"storage_image_reference"`
	NodeAgentSkuId             string                 `json:"node_agent_sku_id"`
	AccountName                string                 `json:"account_name"`
	AutoScale                  []AzurermBatchPoolSpec `json:"auto_scale"`
	StopPendingResizeOperation bool                   `json:"stop_pending_resize_operation"`
	Certificate                []AzurermBatchPoolSpec `json:"certificate"`
	StartTask                  []AzurermBatchPoolSpec `json:"start_task"`
	ResourceGroupName          string                 `json:"resource_group_name"`
	FixedScale                 []AzurermBatchPoolSpec `json:"fixed_scale"`
	MaxTasksPerNode            int                    `json:"max_tasks_per_node"`
	ContainerConfiguration     []AzurermBatchPoolSpec `json:"container_configuration"`
	DisplayName                string                 `json:"display_name"`
	VmSize                     string                 `json:"vm_size"`
}



type AzurermBatchPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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