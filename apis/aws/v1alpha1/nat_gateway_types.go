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

type NatGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatGatewaySpec   `json:"spec,omitempty"`
	Status            NatGatewayStatus `json:"status,omitempty"`
}

type NatGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AllocationID string `json:"allocationID" tf:"allocation_id"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	// +optional
	PrivateIP string `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	// +optional
	PublicIP string `json:"publicIP,omitempty" tf:"public_ip,omitempty"`
	SubnetID string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NatGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NatGatewaySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NatGatewayList is a list of NatGateways
type NatGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NatGateway CRD objects
	Items []NatGateway `json:"items,omitempty"`
}
