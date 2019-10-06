package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	EvaluationInterval string `json:"evaluationInterval,omitempty" tf:"evaluation_interval,omitempty"`
	Formula            string `json:"formula" tf:"formula"`
}

type BatchPoolSpecCertificate struct {
	ID            string `json:"ID" tf:"id"`
	StoreLocation string `json:"storeLocation" tf:"store_location"`
	// +optional
	StoreName string `json:"storeName,omitempty" tf:"store_name,omitempty"`
	// +optional
	Visibility []string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type BatchPoolSpecContainerConfiguration struct {
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type BatchPoolSpecFixedScale struct {
	// +optional
	ResizeTimeout string `json:"resizeTimeout,omitempty" tf:"resize_timeout,omitempty"`
	// +optional
	TargetDedicatedNodes int `json:"targetDedicatedNodes,omitempty" tf:"target_dedicated_nodes,omitempty"`
	// +optional
	TargetLowPriorityNodes int `json:"targetLowPriorityNodes,omitempty" tf:"target_low_priority_nodes,omitempty"`
}

type BatchPoolSpecStartTaskResourceFile struct {
	// +optional
	AutoStorageContainerName string `json:"autoStorageContainerName,omitempty" tf:"auto_storage_container_name,omitempty"`
	// +optional
	BlobPrefix string `json:"blobPrefix,omitempty" tf:"blob_prefix,omitempty"`
	// +optional
	FileMode string `json:"fileMode,omitempty" tf:"file_mode,omitempty"`
	// +optional
	FilePath string `json:"filePath,omitempty" tf:"file_path,omitempty"`
	// +optional
	HttpURL string `json:"httpURL,omitempty" tf:"http_url,omitempty"`
	// +optional
	StorageContainerURL string `json:"storageContainerURL,omitempty" tf:"storage_container_url,omitempty"`
}

type BatchPoolSpecStartTaskUserIdentityAutoUser struct {
	// +optional
	ElevationLevel string `json:"elevationLevel,omitempty" tf:"elevation_level,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type BatchPoolSpecStartTaskUserIdentity struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoUser []BatchPoolSpecStartTaskUserIdentityAutoUser `json:"autoUser,omitempty" tf:"auto_user,omitempty"`
	// +optional
	UserName string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type BatchPoolSpecStartTask struct {
	CommandLine string `json:"commandLine" tf:"command_line"`
	// +optional
	Environment map[string]string `json:"environment,omitempty" tf:"environment,omitempty"`
	// +optional
	MaxTaskRetryCount int `json:"maxTaskRetryCount,omitempty" tf:"max_task_retry_count,omitempty"`
	// +optional
	ResourceFile []BatchPoolSpecStartTaskResourceFile `json:"resourceFile,omitempty" tf:"resource_file,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	UserIdentity []BatchPoolSpecStartTaskUserIdentity `json:"userIdentity" tf:"user_identity"`
	// +optional
	WaitForSuccess bool `json:"waitForSuccess,omitempty" tf:"wait_for_success,omitempty"`
}

type BatchPoolSpecStorageImageReference struct {
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Offer string `json:"offer,omitempty" tf:"offer,omitempty"`
	// +optional
	Publisher string `json:"publisher,omitempty" tf:"publisher,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type BatchPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName string `json:"accountName" tf:"account_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoScale []BatchPoolSpecAutoScale `json:"autoScale,omitempty" tf:"auto_scale,omitempty"`
	// +optional
	Certificate []BatchPoolSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ContainerConfiguration []BatchPoolSpecContainerConfiguration `json:"containerConfiguration,omitempty" tf:"container_configuration,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedScale []BatchPoolSpecFixedScale `json:"fixedScale,omitempty" tf:"fixed_scale,omitempty"`
	// +optional
	MaxTasksPerNode   int    `json:"maxTasksPerNode,omitempty" tf:"max_tasks_per_node,omitempty"`
	Name              string `json:"name" tf:"name"`
	NodeAgentSkuID    string `json:"nodeAgentSkuID" tf:"node_agent_sku_id"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StartTask []BatchPoolSpecStartTask `json:"startTask,omitempty" tf:"start_task,omitempty"`
	// +optional
	StopPendingResizeOperation bool `json:"stopPendingResizeOperation,omitempty" tf:"stop_pending_resize_operation,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	StorageImageReference []BatchPoolSpecStorageImageReference `json:"storageImageReference" tf:"storage_image_reference"`
	VmSize                string                               `json:"vmSize" tf:"vm_size"`
}

type BatchPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BatchPoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
