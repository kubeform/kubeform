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

type AwsVpnConnectionRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpnConnectionRouteSpec   `json:"spec,omitempty"`
	Status            AwsVpnConnectionRouteStatus `json:"status,omitempty"`
}

type AwsVpnConnectionRouteSpec struct {
	DestinationCidrBlock string `json:"destination_cidr_block"`
	VpnConnectionId      string `json:"vpn_connection_id"`
}

type AwsVpnConnectionRouteStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpnConnectionRouteList is a list of AwsVpnConnectionRoutes
type AwsVpnConnectionRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpnConnectionRoute CRD objects
	Items []AwsVpnConnectionRoute `json:"items,omitempty"`
}
