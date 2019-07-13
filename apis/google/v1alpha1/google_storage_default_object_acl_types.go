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

type GoogleStorageDefaultObjectAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageDefaultObjectAclSpec   `json:"spec,omitempty"`
	Status            GoogleStorageDefaultObjectAclStatus `json:"status,omitempty"`
}

type GoogleStorageDefaultObjectAclSpec struct {
	Bucket     string   `json:"bucket"`
	RoleEntity []string `json:"role_entity"`
}



type GoogleStorageDefaultObjectAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageDefaultObjectAclList is a list of GoogleStorageDefaultObjectAcls
type GoogleStorageDefaultObjectAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageDefaultObjectAcl CRD objects
	Items []GoogleStorageDefaultObjectAcl `json:"items,omitempty"`
}