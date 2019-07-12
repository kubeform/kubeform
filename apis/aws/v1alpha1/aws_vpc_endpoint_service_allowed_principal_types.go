package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceAllowedPrincipal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointServiceAllowedPrincipalSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointServiceAllowedPrincipalStatus `json:"status,omitempty"`
}

type AwsVpcEndpointServiceAllowedPrincipalSpec struct {
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
	PrincipalArn         string `json:"principal_arn"`
}

type AwsVpcEndpointServiceAllowedPrincipalStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointServiceAllowedPrincipalList is a list of AwsVpcEndpointServiceAllowedPrincipals
type AwsVpcEndpointServiceAllowedPrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpointServiceAllowedPrincipal CRD objects
	Items []AwsVpcEndpointServiceAllowedPrincipal `json:"items,omitempty"`
}
