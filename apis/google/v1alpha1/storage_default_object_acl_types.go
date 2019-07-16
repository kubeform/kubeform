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

type StorageDefaultObjectAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageDefaultObjectAclSpec   `json:"spec,omitempty"`
	Status            StorageDefaultObjectAclStatus `json:"status,omitempty"`
}

type StorageDefaultObjectAclSpec struct {
	Bucket string `json:"bucket"`
}

type StorageDefaultObjectAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageDefaultObjectAclList is a list of StorageDefaultObjectAcls
type StorageDefaultObjectAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageDefaultObjectAcl CRD objects
	Items []StorageDefaultObjectAcl `json:"items,omitempty"`
}
