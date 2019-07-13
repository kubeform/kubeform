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

type AzurermEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventhubSpec   `json:"spec,omitempty"`
	Status            AzurermEventhubStatus `json:"status,omitempty"`
}

type AzurermEventhubSpecCaptureDescriptionDestination struct {
	Name              string `json:"name"`
	ArchiveNameFormat string `json:"archive_name_format"`
	BlobContainerName string `json:"blob_container_name"`
	StorageAccountId  string `json:"storage_account_id"`
}

type AzurermEventhubSpecCaptureDescription struct {
	Enabled           bool                                    `json:"enabled"`
	SkipEmptyArchives bool                                    `json:"skip_empty_archives"`
	Encoding          string                                  `json:"encoding"`
	IntervalInSeconds int                                     `json:"interval_in_seconds"`
	SizeLimitInBytes  int                                     `json:"size_limit_in_bytes"`
	Destination       []AzurermEventhubSpecCaptureDescription `json:"destination"`
}

type AzurermEventhubSpec struct {
	ResourceGroupName  string                `json:"resource_group_name"`
	Location           string                `json:"location"`
	PartitionCount     int                   `json:"partition_count"`
	MessageRetention   int                   `json:"message_retention"`
	CaptureDescription []AzurermEventhubSpec `json:"capture_description"`
	PartitionIds       []string              `json:"partition_ids"`
	Name               string                `json:"name"`
	NamespaceName      string                `json:"namespace_name"`
}



type AzurermEventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermEventhubList is a list of AzurermEventhubs
type AzurermEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventhub CRD objects
	Items []AzurermEventhub `json:"items,omitempty"`
}