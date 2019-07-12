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

type AwsEipAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEipAssociationSpec   `json:"spec,omitempty"`
	Status            AwsEipAssociationStatus `json:"status,omitempty"`
}

type AwsEipAssociationSpec struct {
	AllowReassociation bool   `json:"allow_reassociation"`
	InstanceId         string `json:"instance_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
	PrivateIpAddress   string `json:"private_ip_address"`
	PublicIp           string `json:"public_ip"`
	AllocationId       string `json:"allocation_id"`
}

type AwsEipAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEipAssociationList is a list of AwsEipAssociations
type AwsEipAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEipAssociation CRD objects
	Items []AwsEipAssociation `json:"items,omitempty"`
}
