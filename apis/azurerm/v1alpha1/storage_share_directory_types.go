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

type StorageShareDirectory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageShareDirectorySpec   `json:"spec,omitempty"`
	Status            StorageShareDirectoryStatus `json:"status,omitempty"`
}

type StorageShareDirectorySpec struct {
	// +optional
	Metadata           map[string]string `json:"metadata,omitempty"`
	Name               string            `json:"name"`
	ShareName          string            `json:"share_name"`
	StorageAccountName string            `json:"storage_account_name"`
}

type StorageShareDirectoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageShareDirectoryList is a list of StorageShareDirectorys
type StorageShareDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageShareDirectory CRD objects
	Items []StorageShareDirectory `json:"items,omitempty"`
}
