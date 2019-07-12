package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpnGatewayAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsVpnGatewayAttachmentStatus `json:"status,omitempty"`
}

type AwsVpnGatewayAttachmentSpec struct {
	VpcId        string `json:"vpc_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}



type AwsVpnGatewayAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayAttachmentList is a list of AwsVpnGatewayAttachments
type AwsVpnGatewayAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpnGatewayAttachment CRD objects
	Items []AwsVpnGatewayAttachment `json:"items,omitempty"`
}