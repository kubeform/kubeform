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

type StorageObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageObjectAccessControlSpec   `json:"spec,omitempty"`
	Status            StorageObjectAccessControlStatus `json:"status,omitempty"`
}

type StorageObjectAccessControlSpec struct {
	Bucket string `json:"bucket"`
	Entity string `json:"entity"`
	Object string `json:"object"`
	Role   string `json:"role"`
}

type StorageObjectAccessControlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageObjectAccessControlList is a list of StorageObjectAccessControls
type StorageObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageObjectAccessControl CRD objects
	Items []StorageObjectAccessControl `json:"items,omitempty"`
}
