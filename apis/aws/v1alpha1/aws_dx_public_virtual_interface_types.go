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

type AwsDxPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxPublicVirtualInterfaceSpec struct {
	Vlan                int               `json:"vlan"`
	BgpAuthKey          string            `json:"bgp_auth_key"`
	AddressFamily       string            `json:"address_family"`
	AmazonAddress       string            `json:"amazon_address"`
	RouteFilterPrefixes []string          `json:"route_filter_prefixes"`
	Tags                map[string]string `json:"tags"`
	Arn                 string            `json:"arn"`
	Name                string            `json:"name"`
	BgpAsn              int               `json:"bgp_asn"`
	CustomerAddress     string            `json:"customer_address"`
	AwsDevice           string            `json:"aws_device"`
	ConnectionId        string            `json:"connection_id"`
}

type AwsDxPublicVirtualInterfaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxPublicVirtualInterfaceList is a list of AwsDxPublicVirtualInterfaces
type AwsDxPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxPublicVirtualInterface CRD objects
	Items []AwsDxPublicVirtualInterface `json:"items,omitempty"`
}
