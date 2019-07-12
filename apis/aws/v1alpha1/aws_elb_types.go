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

type AwsElb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElbSpec   `json:"spec,omitempty"`
	Status            AwsElbStatus `json:"status,omitempty"`
}

type AwsElbSpecAccessLogs struct {
	Interval     int    `json:"interval"`
	Bucket       string `json:"bucket"`
	BucketPrefix string `json:"bucket_prefix"`
	Enabled      bool   `json:"enabled"`
}

type AwsElbSpecHealthCheck struct {
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Target             string `json:"target"`
	Interval           int    `json:"interval"`
}

type AwsElbSpecListener struct {
	InstancePort     int    `json:"instance_port"`
	InstanceProtocol string `json:"instance_protocol"`
	LbPort           int    `json:"lb_port"`
	LbProtocol       string `json:"lb_protocol"`
	SslCertificateId string `json:"ssl_certificate_id"`
}

type AwsElbSpec struct {
	NamePrefix                string            `json:"name_prefix"`
	Internal                  bool              `json:"internal"`
	Instances                 []string          `json:"instances"`
	SourceSecurityGroup       string            `json:"source_security_group"`
	AccessLogs                []AwsElbSpec      `json:"access_logs"`
	HealthCheck               []AwsElbSpec      `json:"health_check"`
	Tags                      map[string]string `json:"tags"`
	Name                      string            `json:"name"`
	CrossZoneLoadBalancing    bool              `json:"cross_zone_load_balancing"`
	SecurityGroups            []string          `json:"security_groups"`
	ZoneId                    string            `json:"zone_id"`
	AvailabilityZones         []string          `json:"availability_zones"`
	DnsName                   string            `json:"dns_name"`
	Arn                       string            `json:"arn"`
	SourceSecurityGroupId     string            `json:"source_security_group_id"`
	Subnets                   []string          `json:"subnets"`
	IdleTimeout               int               `json:"idle_timeout"`
	ConnectionDraining        bool              `json:"connection_draining"`
	ConnectionDrainingTimeout int               `json:"connection_draining_timeout"`
	Listener                  []AwsElbSpec      `json:"listener"`
}

type AwsElbStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElbList is a list of AwsElbs
type AwsElbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElb CRD objects
	Items []AwsElb `json:"items,omitempty"`
}
