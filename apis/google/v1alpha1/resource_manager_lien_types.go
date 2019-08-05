package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty"`
	// +optional
	Name         string   `json:"name,omitempty" tf:"name,omitempty"`
	Origin       string   `json:"origin" tf:"origin"`
	Parent       string   `json:"parent" tf:"parent"`
	Reason       string   `json:"reason" tf:"reason"`
	Restrictions []string `json:"restrictions" tf:"restrictions"`
}

type ResourceManagerLienStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ResourceManagerLienSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
