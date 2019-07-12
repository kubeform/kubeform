package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStorageBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageBlobSpec   `json:"spec,omitempty"`
	Status            AzurermStorageBlobStatus `json:"status,omitempty"`
}

type AzurermStorageBlobSpec struct {
	Source               string            `json:"source"`
	Attempts             int               `json:"attempts"`
	Type                 string            `json:"type"`
	ContentType          string            `json:"content_type"`
	StorageAccountName   string            `json:"storage_account_name"`
	StorageContainerName string            `json:"storage_container_name"`
	Size                 int               `json:"size"`
	SourceUri            string            `json:"source_uri"`
	Url                  string            `json:"url"`
	Parallelism          int               `json:"parallelism"`
	Name                 string            `json:"name"`
	ResourceGroupName    string            `json:"resource_group_name"`
	Metadata             map[string]string `json:"metadata"`
}

type AzurermStorageBlobStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStorageBlobList is a list of AzurermStorageBlobs
type AzurermStorageBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageBlob CRD objects
	Items []AzurermStorageBlob `json:"items,omitempty"`
}
