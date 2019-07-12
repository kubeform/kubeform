package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbSpec   `json:"spec,omitempty"`
	Status            AwsAlbStatus `json:"status,omitempty"`
}

type AwsAlbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

type AwsAlbSpecAccessLogs struct {
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
}

type AwsAlbSpec struct {
	SecurityGroups               []string          `json:"security_groups"`
	SubnetMapping                []AwsAlbSpec      `json:"subnet_mapping"`
	AccessLogs                   []AwsAlbSpec      `json:"access_logs"`
	IdleTimeout                  int               `json:"idle_timeout"`
	EnableCrossZoneLoadBalancing bool              `json:"enable_cross_zone_load_balancing"`
	IpAddressType                string            `json:"ip_address_type"`
	DnsName                      string            `json:"dns_name"`
	NamePrefix                   string            `json:"name_prefix"`
	EnableDeletionProtection     bool              `json:"enable_deletion_protection"`
	Arn                          string            `json:"arn"`
	Name                         string            `json:"name"`
	Internal                     bool              `json:"internal"`
	Subnets                      []string          `json:"subnets"`
	EnableHttp2                  bool              `json:"enable_http2"`
	Tags                         map[string]string `json:"tags"`
	ArnSuffix                    string            `json:"arn_suffix"`
	VpcId                        string            `json:"vpc_id"`
	ZoneId                       string            `json:"zone_id"`
	LoadBalancerType             string            `json:"load_balancer_type"`
}



type AwsAlbStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbList is a list of AwsAlbs
type AwsAlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlb CRD objects
	Items []AwsAlb `json:"items,omitempty"`
}