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

type AwsLb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbSpec   `json:"spec,omitempty"`
	Status            AwsLbStatus `json:"status,omitempty"`
}

type AwsLbSpecAccessLogs struct {
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
	Bucket  string `json:"bucket"`
}

type AwsLbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

type AwsLbSpec struct {
	AccessLogs                   []AwsLbSpec       `json:"access_logs"`
	EnableDeletionProtection     bool              `json:"enable_deletion_protection"`
	VpcId                        string            `json:"vpc_id"`
	ZoneId                       string            `json:"zone_id"`
	Tags                         map[string]string `json:"tags"`
	NamePrefix                   string            `json:"name_prefix"`
	LoadBalancerType             string            `json:"load_balancer_type"`
	SubnetMapping                []AwsLbSpec       `json:"subnet_mapping"`
	EnableHttp2                  bool              `json:"enable_http2"`
	IpAddressType                string            `json:"ip_address_type"`
	DnsName                      string            `json:"dns_name"`
	ArnSuffix                    string            `json:"arn_suffix"`
	Internal                     bool              `json:"internal"`
	Arn                          string            `json:"arn"`
	EnableCrossZoneLoadBalancing bool              `json:"enable_cross_zone_load_balancing"`
	Subnets                      []string          `json:"subnets"`
	IdleTimeout                  int               `json:"idle_timeout"`
	Name                         string            `json:"name"`
	SecurityGroups               []string          `json:"security_groups"`
}

type AwsLbStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLbList is a list of AwsLbs
type AwsLbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLb CRD objects
	Items []AwsLb `json:"items,omitempty"`
}
