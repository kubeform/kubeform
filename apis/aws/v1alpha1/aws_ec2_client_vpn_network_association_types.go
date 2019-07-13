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

type AwsEc2ClientVpnNetworkAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2ClientVpnNetworkAssociationSpec   `json:"spec,omitempty"`
	Status            AwsEc2ClientVpnNetworkAssociationStatus `json:"status,omitempty"`
}

type AwsEc2ClientVpnNetworkAssociationSpec struct {
	Status              string   `json:"status"`
	VpcId               string   `json:"vpc_id"`
	ClientVpnEndpointId string   `json:"client_vpn_endpoint_id"`
	SubnetId            string   `json:"subnet_id"`
	SecurityGroups      []string `json:"security_groups"`
}



type AwsEc2ClientVpnNetworkAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2ClientVpnNetworkAssociationList is a list of AwsEc2ClientVpnNetworkAssociations
type AwsEc2ClientVpnNetworkAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2ClientVpnNetworkAssociation CRD objects
	Items []AwsEc2ClientVpnNetworkAssociation `json:"items,omitempty"`
}