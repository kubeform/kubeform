package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInternetGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsInternetGatewaySpec   `json:"spec,omitempty"`
	Status            AwsInternetGatewayStatus `json:"status,omitempty"`
}

type AwsInternetGatewaySpec struct {
	VpcId   string            `json:"vpc_id"`
	Tags    map[string]string `json:"tags"`
	OwnerId string            `json:"owner_id"`
}



type AwsInternetGatewayStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInternetGatewayList is a list of AwsInternetGateways
type AwsInternetGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInternetGateway CRD objects
	Items []AwsInternetGateway `json:"items,omitempty"`
}