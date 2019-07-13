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

type GoogleStorageBucketAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketAclSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketAclStatus `json:"status,omitempty"`
}

type GoogleStorageBucketAclSpec struct {
	RoleEntity    []string `json:"role_entity"`
	Bucket        string   `json:"bucket"`
	DefaultAcl    string   `json:"default_acl"`
	PredefinedAcl string   `json:"predefined_acl"`
}



type GoogleStorageBucketAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageBucketAclList is a list of GoogleStorageBucketAcls
type GoogleStorageBucketAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucketAcl CRD objects
	Items []GoogleStorageBucketAcl `json:"items,omitempty"`
}