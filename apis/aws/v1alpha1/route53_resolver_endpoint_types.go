package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Route53ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverEndpointSpec   `json:"spec,omitempty"`
	Status            Route53ResolverEndpointStatus `json:"status,omitempty"`
}

type Route53ResolverEndpointSpecIpAddress struct {
	// +optional
	Ip string `json:"ip,omitempty" tf:"ip,omitempty"`
	// +optional
	IpID     string `json:"ipID,omitempty" tf:"ip_id,omitempty"`
	SubnetID string `json:"subnetID" tf:"subnet_id"`
}

type Route53ResolverEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn       string `json:"arn,omitempty" tf:"arn,omitempty"`
	Direction string `json:"direction" tf:"direction"`
	// +optional
	HostVpcID string `json:"hostVpcID,omitempty" tf:"host_vpc_id,omitempty"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=2
	// +kubebuilder:validation:UniqueItems=true
	IpAddress []Route53ResolverEndpointSpecIpAddress `json:"ipAddress" tf:"ip_address"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type Route53ResolverEndpointStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *Route53ResolverEndpointSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ResolverEndpointList is a list of Route53ResolverEndpoints
type Route53ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53ResolverEndpoint CRD objects
	Items []Route53ResolverEndpoint `json:"items,omitempty"`
}