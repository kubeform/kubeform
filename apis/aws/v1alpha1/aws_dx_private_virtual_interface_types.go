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
	DxGatewayId       string            `json:"dx_gateway_id"`
	Mtu               int               `json:"mtu"`
	VpnGatewayId      string            `json:"vpn_gateway_id"`
	JumboFrameCapable bool              `json:"jumbo_frame_capable"`
	BgpAuthKey        string            `json:"bgp_auth_key"`
	AddressFamily     string            `json:"address_family"`
	Tags              map[string]string `json:"tags"`
	AwsDevice         string            `json:"aws_device"`
	ConnectionId      string            `json:"connection_id"`
	Name              string            `json:"name"`
	Vlan              int               `json:"vlan"`
	BgpAsn            int               `json:"bgp_asn"`
	CustomerAddress   string            `json:"customer_address"`
	AmazonAddress     string            `json:"amazon_address"`
	Arn               string            `json:"arn"`
}



type AwsDxPrivateVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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