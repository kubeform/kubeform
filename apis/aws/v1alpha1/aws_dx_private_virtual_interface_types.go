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

type AwsDxPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxPrivateVirtualInterfaceSpec struct {
	ConnectionId      string            `json:"connection_id"`
	VpnGatewayId      string            `json:"vpn_gateway_id"`
	Vlan              int               `json:"vlan"`
	AmazonAddress     string            `json:"amazon_address"`
	BgpAsn            int               `json:"bgp_asn"`
	AddressFamily     string            `json:"address_family"`
	Mtu               int               `json:"mtu"`
	JumboFrameCapable bool              `json:"jumbo_frame_capable"`
	Tags              map[string]string `json:"tags"`
	AwsDevice         string            `json:"aws_device"`
	Arn               string            `json:"arn"`
	Name              string            `json:"name"`
	DxGatewayId       string            `json:"dx_gateway_id"`
	BgpAuthKey        string            `json:"bgp_auth_key"`
	CustomerAddress   string            `json:"customer_address"`
}

type AwsDxPrivateVirtualInterfaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxPrivateVirtualInterfaceList is a list of AwsDxPrivateVirtualInterfaces
type AwsDxPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxPrivateVirtualInterface CRD objects
	Items []AwsDxPrivateVirtualInterface `json:"items,omitempty"`
}
