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

type AwsDxHostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxHostedPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxHostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxHostedPublicVirtualInterfaceSpec struct {
	Arn                 string   `json:"arn"`
	ConnectionId        string   `json:"connection_id"`
	Vlan                int      `json:"vlan"`
	BgpAuthKey          string   `json:"bgp_auth_key"`
	OwnerAccountId      string   `json:"owner_account_id"`
	RouteFilterPrefixes []string `json:"route_filter_prefixes"`
	Name                string   `json:"name"`
	BgpAsn              int      `json:"bgp_asn"`
	AddressFamily       string   `json:"address_family"`
	CustomerAddress     string   `json:"customer_address"`
	AmazonAddress       string   `json:"amazon_address"`
	AwsDevice           string   `json:"aws_device"`
}



type AwsDxHostedPublicVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxHostedPublicVirtualInterfaceList is a list of AwsDxHostedPublicVirtualInterfaces
type AwsDxHostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxHostedPublicVirtualInterface CRD objects
	Items []AwsDxHostedPublicVirtualInterface `json:"items,omitempty"`
}