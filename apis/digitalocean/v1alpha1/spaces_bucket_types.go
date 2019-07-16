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

type SpacesBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpacesBucketSpec   `json:"spec,omitempty"`
	Status            SpacesBucketStatus `json:"status,omitempty"`
}

type SpacesBucketSpec struct {
	// +optional
	Acl string `json:"acl,omitempty"`
	// +optional
	ForceDestroy bool   `json:"force_destroy,omitempty"`
	Name         string `json:"name"`
	// +optional
	Region string `json:"region,omitempty"`
}

type SpacesBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpacesBucketList is a list of SpacesBuckets
type SpacesBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpacesBucket CRD objects
	Items []SpacesBucket `json:"items,omitempty"`
}
