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

type AwsEc2TransitGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2TransitGatewaySpec   `json:"spec,omitempty"`
	Status            AwsEc2TransitGatewayStatus `json:"status,omitempty"`
}

type AwsEc2TransitGatewaySpec struct {
	AmazonSideAsn                  int               `json:"amazon_side_asn"`
	Arn                            string            `json:"arn"`
	DnsSupport                     string            `json:"dns_support"`
	OwnerId                        string            `json:"owner_id"`
	VpnEcmpSupport                 string            `json:"vpn_ecmp_support"`
	Tags                           map[string]string `json:"tags"`
	AssociationDefaultRouteTableId string            `json:"association_default_route_table_id"`
	AutoAcceptSharedAttachments    string            `json:"auto_accept_shared_attachments"`
	DefaultRouteTableAssociation   string            `json:"default_route_table_association"`
	DefaultRouteTablePropagation   string            `json:"default_route_table_propagation"`
	Description                    string            `json:"description"`
	PropagationDefaultRouteTableId string            `json:"propagation_default_route_table_id"`
}



type AwsEc2TransitGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2TransitGatewayList is a list of AwsEc2TransitGateways
type AwsEc2TransitGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2TransitGateway CRD objects
	Items []AwsEc2TransitGateway `json:"items,omitempty"`
}