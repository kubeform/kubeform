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

type AzurermDataLakeStoreFile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataLakeStoreFileSpec   `json:"spec,omitempty"`
	Status            AzurermDataLakeStoreFileStatus `json:"status,omitempty"`
}

type AzurermDataLakeStoreFileSpec struct {
	AccountName    string `json:"account_name"`
	RemoteFilePath string `json:"remote_file_path"`
	LocalFilePath  string `json:"local_file_path"`
}

type AzurermDataLakeStoreFileStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataLakeStoreFileList is a list of AzurermDataLakeStoreFiles
type AzurermDataLakeStoreFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataLakeStoreFile CRD objects
	Items []AzurermDataLakeStoreFile `json:"items,omitempty"`
}
