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

type AwsDxHostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxHostedPrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            AwsDxHostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type AwsDxHostedPrivateVirtualInterfaceSpec struct {
	Arn               string `json:"arn"`
	ConnectionId      string `json:"connection_id"`
	Vlan              int    `json:"vlan"`
	BgpAsn            int    `json:"bgp_asn"`
	AmazonAddress     string `json:"amazon_address"`
	OwnerAccountId    string `json:"owner_account_id"`
	AwsDevice         string `json:"aws_device"`
	Name              string `json:"name"`
	BgpAuthKey        string `json:"bgp_auth_key"`
	AddressFamily     string `json:"address_family"`
	CustomerAddress   string `json:"customer_address"`
	Mtu               int    `json:"mtu"`
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
}



type AwsDxHostedPrivateVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxHostedPrivateVirtualInterfaceList is a list of AwsDxHostedPrivateVirtualInterfaces
type AwsDxHostedPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxHostedPrivateVirtualInterface CRD objects
	Items []AwsDxHostedPrivateVirtualInterface `json:"items,omitempty"`
}