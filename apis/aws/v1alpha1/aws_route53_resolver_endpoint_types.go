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

type AwsRoute53ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53ResolverEndpointSpec   `json:"spec,omitempty"`
	Status            AwsRoute53ResolverEndpointStatus `json:"status,omitempty"`
}

type AwsRoute53ResolverEndpointSpecIpAddress struct {
	Ip       string `json:"ip"`
	IpId     string `json:"ip_id"`
	SubnetId string `json:"subnet_id"`
}

type AwsRoute53ResolverEndpointSpec struct {
	Direction        string                           `json:"direction"`
	IpAddress        []AwsRoute53ResolverEndpointSpec `json:"ip_address"`
	SecurityGroupIds []string                         `json:"security_group_ids"`
	Name             string                           `json:"name"`
	Tags             map[string]string                `json:"tags"`
	Arn              string                           `json:"arn"`
	HostVpcId        string                           `json:"host_vpc_id"`
}

type AwsRoute53ResolverEndpointStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRoute53ResolverEndpointList is a list of AwsRoute53ResolverEndpoints
type AwsRoute53ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53ResolverEndpoint CRD objects
	Items []AwsRoute53ResolverEndpoint `json:"items,omitempty"`
}
