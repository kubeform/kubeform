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

type VpcEndpointService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointServiceSpec   `json:"spec,omitempty"`
	Status            VpcEndpointServiceStatus `json:"status,omitempty"`
}

type VpcEndpointServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AcceptanceRequired bool `json:"acceptanceRequired" tf:"acceptance_required"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AllowedPrincipals []string `json:"allowedPrincipals,omitempty" tf:"allowed_principals,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BaseEndpointDNSNames []string `json:"baseEndpointDNSNames,omitempty" tf:"base_endpoint_dns_names,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	NetworkLoadBalancerArns []string `json:"networkLoadBalancerArns" tf:"network_load_balancer_arns"`
	// +optional
	PrivateDNSName string `json:"privateDNSName,omitempty" tf:"private_dns_name,omitempty"`
	// +optional
	ServiceName string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
	// +optional
	ServiceType string `json:"serviceType,omitempty" tf:"service_type,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
}

type VpcEndpointServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcEndpointServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
