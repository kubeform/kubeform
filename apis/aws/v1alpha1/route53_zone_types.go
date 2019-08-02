package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Route53Zone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ZoneSpec   `json:"spec,omitempty"`
	Status            Route53ZoneStatus `json:"status,omitempty"`
}

type Route53ZoneSpecVpc struct {
	VpcID string `json:"vpcID" tf:"vpc_id"`
	// +optional
	VpcRegion string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`
}

type Route53ZoneSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
	// +optional
	DelegationSetID string `json:"delegationSetID,omitempty" tf:"delegation_set_id,omitempty"`
	// +optional
	ForceDestroy bool   `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Vpc []Route53ZoneSpecVpc `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type Route53ZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *Route53ZoneSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ZoneList is a list of Route53Zones
type Route53ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53Zone CRD objects
	Items []Route53Zone `json:"items,omitempty"`
}
