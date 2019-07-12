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

type AwsVpcEndpointService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointServiceSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointServiceStatus `json:"status,omitempty"`
}

type AwsVpcEndpointServiceSpec struct {
	AllowedPrincipals       []string          `json:"allowed_principals"`
	ServiceType             string            `json:"service_type"`
	State                   string            `json:"state"`
	AcceptanceRequired      bool              `json:"acceptance_required"`
	AvailabilityZones       []string          `json:"availability_zones"`
	BaseEndpointDnsNames    []string          `json:"base_endpoint_dns_names"`
	ManagesVpcEndpoints     bool              `json:"manages_vpc_endpoints"`
	NetworkLoadBalancerArns []string          `json:"network_load_balancer_arns"`
	PrivateDnsName          string            `json:"private_dns_name"`
	ServiceName             string            `json:"service_name"`
	Tags                    map[string]string `json:"tags"`
}

type AwsVpcEndpointServiceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcEndpointServiceList is a list of AwsVpcEndpointServices
type AwsVpcEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpointService CRD objects
	Items []AwsVpcEndpointService `json:"items,omitempty"`
}
