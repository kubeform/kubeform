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

type ResourceManagerLien struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceManagerLienSpec   `json:"spec,omitempty"`
	Status            ResourceManagerLienStatus `json:"status,omitempty"`
}

type ResourceManagerLienSpec struct {
	Origin       string   `json:"origin"`
	Parent       string   `json:"parent"`
	Reason       string   `json:"reason"`
	Restrictions []string `json:"restrictions"`
}

type ResourceManagerLienStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ResourceManagerLienList is a list of ResourceManagerLiens
type ResourceManagerLienList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ResourceManagerLien CRD objects
	Items []ResourceManagerLien `json:"items,omitempty"`
}
