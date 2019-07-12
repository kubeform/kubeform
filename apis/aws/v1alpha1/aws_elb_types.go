package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElbSpec   `json:"spec,omitempty"`
	Status            AwsElbStatus `json:"status,omitempty"`
}

type AwsElbSpecHealthCheck struct {
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Target             string `json:"target"`
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
}

type AwsElbSpecAccessLogs struct {
	Interval     int    `json:"interval"`
	Bucket       string `json:"bucket"`
	BucketPrefix string `json:"bucket_prefix"`
	Enabled      bool   `json:"enabled"`
}

type AwsElbSpecListener struct {
	InstancePort     int    `json:"instance_port"`
	InstanceProtocol string `json:"instance_protocol"`
	LbPort           int    `json:"lb_port"`
	LbProtocol       string `json:"lb_protocol"`
	SslCertificateId string `json:"ssl_certificate_id"`
}

type AwsElbSpec struct {
	ConnectionDrainingTimeout int               `json:"connection_draining_timeout"`
	DnsName                   string            `json:"dns_name"`
	Name                      string            `json:"name"`
	Arn                       string            `json:"arn"`
	Subnets                   []string          `json:"subnets"`
	ConnectionDraining        bool              `json:"connection_draining"`
	Tags                      map[string]string `json:"tags"`
	Internal                  bool              `json:"internal"`
	CrossZoneLoadBalancing    bool              `json:"cross_zone_load_balancing"`
	IdleTimeout               int               `json:"idle_timeout"`
	HealthCheck               []AwsElbSpec      `json:"health_check"`
	NamePrefix                string            `json:"name_prefix"`
	SecurityGroups            []string          `json:"security_groups"`
	SourceSecurityGroup       string            `json:"source_security_group"`
	SourceSecurityGroupId     string            `json:"source_security_group_id"`
	ZoneId                    string            `json:"zone_id"`
	AvailabilityZones         []string          `json:"availability_zones"`
	Instances                 []string          `json:"instances"`
	AccessLogs                []AwsElbSpec      `json:"access_logs"`
	Listener                  []AwsElbSpec      `json:"listener"`
}

type AwsElbStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbList is a list of AwsElbs
type AwsElbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElb CRD objects
	Items []AwsElb `json:"items,omitempty"`
}
