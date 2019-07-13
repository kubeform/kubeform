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

type AwsLbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

type AwsLbSpecAccessLogs struct {
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
}

type AwsLbSpec struct {
	EnableCrossZoneLoadBalancing bool              `json:"enable_cross_zone_load_balancing"`
	SubnetMapping                []AwsLbSpec       `json:"subnet_mapping"`
	AccessLogs                   []AwsLbSpec       `json:"access_logs"`
	EnableDeletionProtection     bool              `json:"enable_deletion_protection"`
	IdleTimeout                  int               `json:"idle_timeout"`
	IpAddressType                string            `json:"ip_address_type"`
	Arn                          string            `json:"arn"`
	Internal                     bool              `json:"internal"`
	Subnets                      []string          `json:"subnets"`
	Tags                         map[string]string `json:"tags"`
	ArnSuffix                    string            `json:"arn_suffix"`
	Name                         string            `json:"name"`
	SecurityGroups               []string          `json:"security_groups"`
	VpcId                        string            `json:"vpc_id"`
	ZoneId                       string            `json:"zone_id"`
	DnsName                      string            `json:"dns_name"`
	NamePrefix                   string            `json:"name_prefix"`
	LoadBalancerType             string            `json:"load_balancer_type"`
	EnableHttp2                  bool              `json:"enable_http2"`
}



type AwsLbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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