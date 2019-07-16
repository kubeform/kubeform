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

type VpnConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnConnectionSpec   `json:"spec,omitempty"`
	Status            VpnConnectionStatus `json:"status,omitempty"`
}

type VpnConnectionSpec struct {
	CustomerGatewayId string `json:"customer_gateway_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TransitGatewayId string `json:"transit_gateway_id,omitempty"`
	Type             string `json:"type"`
	// +optional
	VpnGatewayId string `json:"vpn_gateway_id,omitempty"`
}

type VpnConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnConnectionList is a list of VpnConnections
type VpnConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnConnection CRD objects
	Items []VpnConnection `json:"items,omitempty"`
}
