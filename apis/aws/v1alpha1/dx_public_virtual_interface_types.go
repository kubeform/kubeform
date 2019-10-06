package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DxPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxPublicVirtualInterfaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressFamily string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AwsDevice string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`
	BgpAsn    int    `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey   string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`
	ConnectionID string `json:"connectionID" tf:"connection_id"`
	// +optional
	CustomerAddress string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +kubebuilder:validation:MinItems=1
	RouteFilterPrefixes []string `json:"routeFilterPrefixes" tf:"route_filter_prefixes"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Vlan int               `json:"vlan" tf:"vlan"`
}

type DxPublicVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxPublicVirtualInterfaceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxPublicVirtualInterfaceList is a list of DxPublicVirtualInterfaces
type DxPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxPublicVirtualInterface CRD objects
	Items []DxPublicVirtualInterface `json:"items,omitempty"`
}
