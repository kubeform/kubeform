package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpnGatewaySpec   `json:"spec,omitempty"`
	Status            AwsVpnGatewayStatus `json:"status,omitempty"`
}

type AwsVpnGatewaySpec struct {
	AvailabilityZone string            `json:"availability_zone"`
	AmazonSideAsn    string            `json:"amazon_side_asn"`
	VpcId            string            `json:"vpc_id"`
	Tags             map[string]string `json:"tags"`
}



type AwsVpnGatewayStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayList is a list of AwsVpnGateways
type AwsVpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpnGateway CRD objects
	Items []AwsVpnGateway `json:"items,omitempty"`
}