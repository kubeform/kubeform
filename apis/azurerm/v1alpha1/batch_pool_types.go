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

type BatchPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchPoolSpec   `json:"spec,omitempty"`
	Status            BatchPoolStatus `json:"status,omitempty"`
}

type BatchPoolSpecAutoScale struct {
	// +optional
	EvaluationInterval string `json:"evaluation_interval,omitempty"`
	Formula            string `json:"formula"`
}

type BatchPoolSpecCertificate struct {
	Id            string `json:"id"`
	StoreLocation string `json:"store_location"`
	// +optional
	StoreName string `json:"store_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Visibility []string `json:"visibility,omitempty"`
}

type BatchPoolSpecContainerConfiguration struct {
	// +optional
	Type string `json:"type,omitempty"`
}

type BatchPoolSpecFixedScale struct {
	// +optional
	ResizeTimeout string `json:"resize_timeout,omitempty"`
	// +optional
	TargetDedicatedNodes int `json:"target_dedicated_nodes,omitempty"`
	// +optional
	TargetLowPriorityNodes int `json:"target_low_priority_nodes,omitempty"`
}

type BatchPoolSpecStartTaskResourceFile struct {
	// +optional
	AutoStorageContainerName string `json:"auto_storage_container_name,omitempty"`
	// +optional
	BlobPrefix string `json:"blob_prefix,omitempty"`
	// +optional
	FileMode string `json:"file_mode,omitempty"`
	// +optional
	FilePath string `json:"file_path,omitempty"`
	// +optional
	HttpUrl string `json:"http_url,omitempty"`
	// +optional
	StorageContainerUrl string `json:"storage_container_url,omitempty"`
}

type BatchPoolSpecStartTaskUserIdentityAutoUser struct {
	// +optional
	ElevationLevel string `json:"elevation_level,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty"`
}

type BatchPoolSpecStartTaskUserIdentity struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoUser *[]BatchPoolSpecStartTaskUserIdentity `json:"auto_user,omitempty"`
	// +optional
	UserName string `json:"user_name,omitempty"`
}

type BatchPoolSpecStartTask struct {
	CommandLine string `json:"command_line"`
	// +optional
	Environment map[string]string `json:"environment,omitempty"`
	// +optional
	MaxTaskRetryCount int `json:"max_task_retry_count,omitempty"`
	// +optional
	ResourceFile *[]BatchPoolSpecStartTask `json:"resource_file,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	UserIdentity []BatchPoolSpecStartTask `json:"user_identity"`
	// +optional
	WaitForSuccess bool `json:"wait_for_success,omitempty"`
}

type BatchPoolSpecStorageImageReference struct {
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	Offer string `json:"offer,omitempty"`
	// +optional
	Publisher string `json:"publisher,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty"`
	// +optional
	Version string `json:"version,omitempty"`
}

type BatchPoolSpec struct {
	AccountName string `json:"account_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoScale *[]BatchPoolSpec `json:"auto_scale,omitempty"`
	// +optional
	Certificate *[]BatchPoolSpec `json:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ContainerConfiguration *[]BatchPoolSpec `json:"container_configuration,omitempty"`
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedScale *[]BatchPoolSpec `json:"fixed_scale,omitempty"`
	// +optional
	MaxTasksPerNode   int    `json:"max_tasks_per_node,omitempty"`
	Name              string `json:"name"`
	NodeAgentSkuId    string `json:"node_agent_sku_id"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StartTask *[]BatchPoolSpec `json:"start_task,omitempty"`
	// +optional
	StopPendingResizeOperation bool `json:"stop_pending_resize_operation,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	StorageImageReference []BatchPoolSpec `json:"storage_image_reference"`
	VmSize                string          `json:"vm_size"`
}

type BatchPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchPoolList is a list of BatchPools
type BatchPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchPool CRD objects
	Items []BatchPool `json:"items,omitempty"`
}
