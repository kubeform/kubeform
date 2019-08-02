package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpcEndpointService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointServiceSpec   `json:"spec,omitempty"`
	Status            VpcEndpointServiceStatus `json:"status,omitempty"`
}

type VpcEndpointServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AcceptanceRequired bool `json:"acceptanceRequired" tf:"acceptance_required"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AllowedPrincipals []string `json:"allowedPrincipals,omitempty" tf:"allowed_principals,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	NetworkLoadBalancerArns []string `json:"networkLoadBalancerArns" tf:"network_load_balancer_arns"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VpcEndpointServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *VpcEndpointServiceSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointServiceList is a list of VpcEndpointServices
type VpcEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointService CRD objects
	Items []VpcEndpointService `json:"items,omitempty"`
}
