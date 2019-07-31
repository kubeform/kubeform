package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Eventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubSpec   `json:"spec,omitempty"`
	Status            EventhubStatus `json:"status,omitempty"`
}

type EventhubSpecCaptureDescriptionDestination struct {
	ArchiveNameFormat string `json:"archiveNameFormat" tf:"archive_name_format"`
	BlobContainerName string `json:"blobContainerName" tf:"blob_container_name"`
	Name              string `json:"name" tf:"name"`
	StorageAccountID  string `json:"storageAccountID" tf:"storage_account_id"`
}

type EventhubSpecCaptureDescription struct {
	// +kubebuilder:validation:MaxItems=1
	Destination []EventhubSpecCaptureDescriptionDestination `json:"destination" tf:"destination"`
	Enabled     bool                                        `json:"enabled" tf:"enabled"`
	Encoding    string                                      `json:"encoding" tf:"encoding"`
	// +optional
	IntervalInSeconds int `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`
	// +optional
	SizeLimitInBytes int `json:"sizeLimitInBytes,omitempty" tf:"size_limit_in_bytes,omitempty"`
	// +optional
	SkipEmptyArchives bool `json:"skipEmptyArchives,omitempty" tf:"skip_empty_archives,omitempty"`
}

type EventhubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	CaptureDescription []EventhubSpecCaptureDescription `json:"captureDescription,omitempty" tf:"capture_description,omitempty"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty" tf:"location,omitempty"`
	MessageRetention  int    `json:"messageRetention" tf:"message_retention"`
	Name              string `json:"name" tf:"name"`
	NamespaceName     string `json:"namespaceName" tf:"namespace_name"`
	PartitionCount    int    `json:"partitionCount" tf:"partition_count"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type EventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubList is a list of Eventhubs
type EventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Eventhub CRD objects
	Items []Eventhub `json:"items,omitempty"`
}
