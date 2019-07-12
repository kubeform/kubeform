package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointServiceSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointServiceStatus `json:"status,omitempty"`
}

type AwsVpcEndpointServiceSpec struct {
	AllowedPrincipals       []string `json:"allowed_principals"`
	AvailabilityZones       []string `json:"availability_zones"`
	PrivateDnsName          string   `json:"private_dns_name"`
	BaseEndpointDnsNames    []string `json:"base_endpoint_dns_names"`
	AcceptanceRequired      bool     `json:"acceptance_required"`
	NetworkLoadBalancerArns []string `json:"network_load_balancer_arns"`
	State                   string   `json:"state"`
	ServiceName             string   `json:"service_name"`
	ServiceType             string   `json:"service_type"`
}

type AwsVpcEndpointServiceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointServiceList is a list of AwsVpcEndpointServices
type AwsVpcEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpointService CRD objects
	Items []AwsVpcEndpointService `json:"items,omitempty"`
}
