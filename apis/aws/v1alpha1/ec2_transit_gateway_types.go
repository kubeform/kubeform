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

type Ec2TransitGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewaySpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayStatus `json:"status,omitempty"`
}

type Ec2TransitGatewaySpec struct {
	// +optional
	AmazonSideAsn int `json:"amazon_side_asn,omitempty"`
	// +optional
	AutoAcceptSharedAttachments string `json:"auto_accept_shared_attachments,omitempty"`
	// +optional
	DefaultRouteTableAssociation string `json:"default_route_table_association,omitempty"`
	// +optional
	DefaultRouteTablePropagation string `json:"default_route_table_propagation,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DnsSupport string `json:"dns_support,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	VpnEcmpSupport string `json:"vpn_ecmp_support,omitempty"`
}

type Ec2TransitGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2TransitGatewayList is a list of Ec2TransitGateways
type Ec2TransitGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2TransitGateway CRD objects
	Items []Ec2TransitGateway `json:"items,omitempty"`
}
