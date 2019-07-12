package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxHostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxHostedPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxHostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxHostedPrivateVirtualInterfaceSpec struct {
	CustomerAddress   string `json:"customer_address"`
	AmazonAddress     string `json:"amazon_address"`
	Mtu               int    `json:"mtu"`
	AwsDevice         string `json:"aws_device"`
	ConnectionId      string `json:"connection_id"`
	BgpAsn            int    `json:"bgp_asn"`
	Vlan              int    `json:"vlan"`
	BgpAuthKey        string `json:"bgp_auth_key"`
	AddressFamily     string `json:"address_family"`
	OwnerAccountId    string `json:"owner_account_id"`
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
	Arn               string `json:"arn"`
	Name              string `json:"name"`
}



type AwsDxHostedPrivateVirtualInterfaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxHostedPrivateVirtualInterfaceList is a list of AwsDxHostedPrivateVirtualInterfaces
type AwsDxHostedPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxHostedPrivateVirtualInterface CRD objects
	Items []AwsDxHostedPrivateVirtualInterface `json:"items,omitempty"`
}