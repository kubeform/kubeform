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

type AwsVpnGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpnGatewaySpec   `json:"spec,omitempty"`
	Status            AwsVpnGatewayStatus `json:"status,omitempty"`
}

type AwsVpnGatewaySpec struct {
	Tags             map[string]string `json:"tags"`
	AvailabilityZone string            `json:"availability_zone"`
	AmazonSideAsn    string            `json:"amazon_side_asn"`
	VpcId            string            `json:"vpc_id"`
}



type AwsVpnGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpnGatewayList is a list of AwsVpnGateways
type AwsVpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpnGateway CRD objects
	Items []AwsVpnGateway `json:"items,omitempty"`
}