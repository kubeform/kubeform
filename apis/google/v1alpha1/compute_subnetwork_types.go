package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSubnetworkSpec   `json:"spec,omitempty"`
	Status            ComputeSubnetworkStatus `json:"status,omitempty"`
}

type ComputeSubnetworkSpecSecondaryIPRange struct {
	IpCIDRRange string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	RangeName   string `json:"rangeName" tf:"range_name"`
}

type ComputeSubnetworkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EnableFlowLogs bool `json:"enableFlowLogs,omitempty" tf:"enable_flow_logs,omitempty"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`
	// +optional
	GatewayAddress string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`
	IpCIDRRange    string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	Name           string `json:"name" tf:"name"`
	Network        string `json:"network" tf:"network"`
	// +optional
	PrivateIPGoogleAccess bool `json:"privateIPGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	SecondaryIPRange []ComputeSubnetworkSpecSecondaryIPRange `json:"secondaryIPRange,omitempty" tf:"secondary_ip_range,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeSubnetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeSubnetworkSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkList is a list of ComputeSubnetworks
type ComputeSubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSubnetwork CRD objects
	Items []ComputeSubnetwork `json:"items,omitempty"`
}
