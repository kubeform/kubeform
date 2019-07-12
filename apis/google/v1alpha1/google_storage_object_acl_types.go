package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleStorageObjectAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageObjectAclSpec   `json:"spec,omitempty"`
	Status            GoogleStorageObjectAclStatus `json:"status,omitempty"`
}

type GoogleStorageObjectAclSpec struct {
	RoleEntity    []string `json:"role_entity"`
	Bucket        string   `json:"bucket"`
	Object        string   `json:"object"`
	PredefinedAcl string   `json:"predefined_acl"`
}

type GoogleStorageObjectAclStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleStorageObjectAclList is a list of GoogleStorageObjectAcls
type GoogleStorageObjectAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageObjectAcl CRD objects
	Items []GoogleStorageObjectAcl `json:"items,omitempty"`
}
