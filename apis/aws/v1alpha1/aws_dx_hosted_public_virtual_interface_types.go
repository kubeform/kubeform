package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxHostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxHostedPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxHostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxHostedPublicVirtualInterfaceSpec struct {
	BgpAsn              int      `json:"bgp_asn"`
	BgpAuthKey          string   `json:"bgp_auth_key"`
	CustomerAddress     string   `json:"customer_address"`
	AmazonAddress       string   `json:"amazon_address"`
	OwnerAccountId      string   `json:"owner_account_id"`
	RouteFilterPrefixes []string `json:"route_filter_prefixes"`
	AwsDevice           string   `json:"aws_device"`
	Arn                 string   `json:"arn"`
	ConnectionId        string   `json:"connection_id"`
	Name                string   `json:"name"`
	Vlan                int      `json:"vlan"`
	AddressFamily       string   `json:"address_family"`
}



type AwsDxHostedPublicVirtualInterfaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxHostedPublicVirtualInterfaceList is a list of AwsDxHostedPublicVirtualInterfaces
type AwsDxHostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxHostedPublicVirtualInterface CRD objects
	Items []AwsDxHostedPublicVirtualInterface `json:"items,omitempty"`
}