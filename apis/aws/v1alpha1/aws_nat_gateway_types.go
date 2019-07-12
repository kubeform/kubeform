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

type AwsNatGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNatGatewaySpec   `json:"spec,omitempty"`
	Status            AwsNatGatewayStatus `json:"status,omitempty"`
}

type AwsNatGatewaySpec struct {
	AllocationId       string            `json:"allocation_id"`
	SubnetId           string            `json:"subnet_id"`
	NetworkInterfaceId string            `json:"network_interface_id"`
	PrivateIp          string            `json:"private_ip"`
	PublicIp           string            `json:"public_ip"`
	Tags               map[string]string `json:"tags"`
}

type AwsNatGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNatGatewayList is a list of AwsNatGateways
type AwsNatGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNatGateway CRD objects
	Items []AwsNatGateway `json:"items,omitempty"`
}
