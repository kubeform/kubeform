package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DataLakeStoreFile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeStoreFileSpec   `json:"spec,omitempty"`
	Status            DataLakeStoreFileStatus `json:"status,omitempty"`
}

type DataLakeStoreFileSpec struct {
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
	AccountName    string                    `json:"accountName" tf:"account_name"`
	LocalFilePath  string                    `json:"localFilePath" tf:"local_file_path"`
	RemoteFilePath string                    `json:"remoteFilePath" tf:"remote_file_path"`
}

type DataLakeStoreFileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataLakeStoreFileList is a list of DataLakeStoreFiles
type DataLakeStoreFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeStoreFile CRD objects
	Items []DataLakeStoreFile `json:"items,omitempty"`
}
