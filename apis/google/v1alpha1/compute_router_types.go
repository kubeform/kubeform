package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ComputeRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterSpec   `json:"spec,omitempty"`
	Status            ComputeRouterStatus `json:"status,omitempty"`
}

type ComputeRouterSpecBgpAdvertisedIPRanges struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Range string `json:"range,omitempty" tf:"range,omitempty"`
}

type ComputeRouterSpecBgp struct {
	// +optional
	AdvertiseMode string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`
	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`
	// +optional
	AdvertisedIPRanges []ComputeRouterSpecBgpAdvertisedIPRanges `json:"advertisedIPRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`
	Asn                int                                      `json:"asn" tf:"asn"`
}

type ComputeRouterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Bgp []ComputeRouterSpecBgp `json:"bgp,omitempty" tf:"bgp,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Network     string `json:"network" tf:"network"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeRouterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeRouterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterList is a list of ComputeRouters
type ComputeRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouter CRD objects
	Items []ComputeRouter `json:"items,omitempty"`
}
