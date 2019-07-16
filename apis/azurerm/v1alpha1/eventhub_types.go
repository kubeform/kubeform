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

type Eventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubSpec   `json:"spec,omitempty"`
	Status            EventhubStatus `json:"status,omitempty"`
}

type EventhubSpecCaptureDescriptionDestination struct {
	ArchiveNameFormat string `json:"archive_name_format"`
	BlobContainerName string `json:"blob_container_name"`
	Name              string `json:"name"`
	StorageAccountId  string `json:"storage_account_id"`
}

type EventhubSpecCaptureDescription struct {
	// +kubebuilder:validation:MaxItems=1
	Destination []EventhubSpecCaptureDescription `json:"destination"`
	Enabled     bool                             `json:"enabled"`
	Encoding    string                           `json:"encoding"`
	// +optional
	IntervalInSeconds int `json:"interval_in_seconds,omitempty"`
	// +optional
	SizeLimitInBytes int `json:"size_limit_in_bytes,omitempty"`
	// +optional
	SkipEmptyArchives bool `json:"skip_empty_archives,omitempty"`
}

type EventhubSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CaptureDescription *[]EventhubSpec `json:"capture_description,omitempty"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty"`
	MessageRetention  int    `json:"message_retention"`
	Name              string `json:"name"`
	NamespaceName     string `json:"namespace_name"`
	PartitionCount    int    `json:"partition_count"`
	ResourceGroupName string `json:"resource_group_name"`
}

type EventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
