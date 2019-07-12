package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayVpcLink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayVpcLinkSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayVpcLinkStatus `json:"status,omitempty"`
}

type AwsApiGatewayVpcLinkSpec struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TargetArns  []string `json:"target_arns"`
}

type AwsApiGatewayVpcLinkStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayVpcLinkList is a list of AwsApiGatewayVpcLinks
type AwsApiGatewayVpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayVpcLink CRD objects
	Items []AwsApiGatewayVpcLink `json:"items,omitempty"`
}
