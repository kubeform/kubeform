package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbSpec   `json:"spec,omitempty"`
	Status            AwsLbStatus `json:"status,omitempty"`
}

type AwsLbSpecAccessLogs struct {
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
}

type AwsLbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

type AwsLbSpec struct {
	SecurityGroups               []string          `json:"security_groups"`
	AccessLogs                   []AwsLbSpec       `json:"access_logs"`
	IpAddressType                string            `json:"ip_address_type"`
	DnsName                      string            `json:"dns_name"`
	Tags                         map[string]string `json:"tags"`
	ArnSuffix                    string            `json:"arn_suffix"`
	NamePrefix                   string            `json:"name_prefix"`
	Internal                     bool              `json:"internal"`
	Arn                          string            `json:"arn"`
	Name                         string            `json:"name"`
	LoadBalancerType             string            `json:"load_balancer_type"`
	EnableHttp2                  bool              `json:"enable_http2"`
	EnableDeletionProtection     bool              `json:"enable_deletion_protection"`
	IdleTimeout                  int               `json:"idle_timeout"`
	EnableCrossZoneLoadBalancing bool              `json:"enable_cross_zone_load_balancing"`
	VpcId                        string            `json:"vpc_id"`
	ZoneId                       string            `json:"zone_id"`
	Subnets                      []string          `json:"subnets"`
	SubnetMapping                []AwsLbSpec       `json:"subnet_mapping"`
}

type AwsLbStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbList is a list of AwsLbs
type AwsLbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLb CRD objects
	Items []AwsLb `json:"items,omitempty"`
}
