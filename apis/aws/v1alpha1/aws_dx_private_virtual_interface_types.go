package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxPrivateVirtualInterfaceSpec struct {
	AddressFamily     string            `json:"address_family"`
	AwsDevice         string            `json:"aws_device"`
	CustomerAddress   string            `json:"customer_address"`
	AmazonAddress     string            `json:"amazon_address"`
	JumboFrameCapable bool              `json:"jumbo_frame_capable"`
	Name              string            `json:"name"`
	Vlan              int               `json:"vlan"`
	BgpAsn            int               `json:"bgp_asn"`
	Mtu               int               `json:"mtu"`
	Arn               string            `json:"arn"`
	VpnGatewayId      string            `json:"vpn_gateway_id"`
	BgpAuthKey        string            `json:"bgp_auth_key"`
	Tags              map[string]string `json:"tags"`
	ConnectionId      string            `json:"connection_id"`
	DxGatewayId       string            `json:"dx_gateway_id"`
}



type AwsDxPrivateVirtualInterfaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxPrivateVirtualInterfaceList is a list of AwsDxPrivateVirtualInterfaces
type AwsDxPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxPrivateVirtualInterface CRD objects
	Items []AwsDxPrivateVirtualInterface `json:"items,omitempty"`
}