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

type StorageBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBlobSpec   `json:"spec,omitempty"`
	Status            StorageBlobStatus `json:"status,omitempty"`
}

type StorageBlobSpec struct {
	// +optional
	Attempts int `json:"attempts,omitempty"`
	// +optional
	ContentType string `json:"content_type,omitempty"`
	Name        string `json:"name"`
	// +optional
	Parallelism       int    `json:"parallelism,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Size int `json:"size,omitempty"`
	// +optional
	Source string `json:"source,omitempty"`
	// +optional
	SourceUri            string `json:"source_uri,omitempty"`
	StorageAccountName   string `json:"storage_account_name"`
	StorageContainerName string `json:"storage_container_name"`
	// +optional
	Type string `json:"type,omitempty"`
}

type StorageBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBlobList is a list of StorageBlobs
type StorageBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBlob CRD objects
	Items []StorageBlob `json:"items,omitempty"`
}
