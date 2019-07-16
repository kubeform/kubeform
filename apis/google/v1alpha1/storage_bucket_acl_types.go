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

type StorageBucketAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketAclSpec   `json:"spec,omitempty"`
	Status            StorageBucketAclStatus `json:"status,omitempty"`
}

type StorageBucketAclSpec struct {
	Bucket string `json:"bucket"`
	// +optional
	DefaultAcl string `json:"default_acl,omitempty"`
	// +optional
	PredefinedAcl string `json:"predefined_acl,omitempty"`
}

type StorageBucketAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketAclList is a list of StorageBucketAcls
type StorageBucketAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketAcl CRD objects
	Items []StorageBucketAcl `json:"items,omitempty"`
}
