package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpnConnectionRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnConnectionRouteSpec   `json:"spec,omitempty"`
	Status            VpnConnectionRouteStatus `json:"status,omitempty"`
}

type VpnConnectionRouteSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DestinationCIDRBlock string `json:"destinationCIDRBlock" tf:"destination_cidr_block"`
	VpnConnectionID      string `json:"vpnConnectionID" tf:"vpn_connection_id"`
}

type VpnConnectionRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnConnectionRouteList is a list of VpnConnectionRoutes
type VpnConnectionRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnConnectionRoute CRD objects
	Items []VpnConnectionRoute `json:"items,omitempty"`
}
