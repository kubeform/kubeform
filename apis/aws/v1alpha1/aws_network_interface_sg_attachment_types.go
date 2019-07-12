package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceSgAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNetworkInterfaceSgAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsNetworkInterfaceSgAttachmentStatus `json:"status,omitempty"`
}

type AwsNetworkInterfaceSgAttachmentSpec struct {
	SecurityGroupId    string `json:"security_group_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
}



type AwsNetworkInterfaceSgAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkInterfaceSgAttachmentList is a list of AwsNetworkInterfaceSgAttachments
type AwsNetworkInterfaceSgAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNetworkInterfaceSgAttachment CRD objects
	Items []AwsNetworkInterfaceSgAttachment `json:"items,omitempty"`
}