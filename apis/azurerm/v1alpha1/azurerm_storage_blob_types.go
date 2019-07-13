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

type AzurermStorageBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageBlobSpec   `json:"spec,omitempty"`
	Status            AzurermStorageBlobStatus `json:"status,omitempty"`
}

type AzurermStorageBlobSpec struct {
	Name                 string            `json:"name"`
	SourceUri            string            `json:"source_uri"`
	Parallelism          int               `json:"parallelism"`
	Attempts             int               `json:"attempts"`
	Size                 int               `json:"size"`
	ContentType          string            `json:"content_type"`
	Source               string            `json:"source"`
	Url                  string            `json:"url"`
	ResourceGroupName    string            `json:"resource_group_name"`
	StorageAccountName   string            `json:"storage_account_name"`
	StorageContainerName string            `json:"storage_container_name"`
	Type                 string            `json:"type"`
	Metadata             map[string]string `json:"metadata"`
}



type AzurermStorageBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStorageBlobList is a list of AzurermStorageBlobs
type AzurermStorageBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageBlob CRD objects
	Items []AzurermStorageBlob `json:"items,omitempty"`
}