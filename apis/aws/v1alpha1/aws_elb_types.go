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

type AwsElbSpecListener struct {
	InstanceProtocol string `json:"instance_protocol"`
	LbPort           int    `json:"lb_port"`
	LbProtocol       string `json:"lb_protocol"`
	SslCertificateId string `json:"ssl_certificate_id"`
	InstancePort     int    `json:"instance_port"`
}

type AwsElbSpecAccessLogs struct {
	Enabled      bool   `json:"enabled"`
	Interval     int    `json:"interval"`
	Bucket       string `json:"bucket"`
	BucketPrefix string `json:"bucket_prefix"`
}

type AwsElbSpecHealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Target             string `json:"target"`
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
}

type AwsElbSpec struct {
	Name                      string            `json:"name"`
	CrossZoneLoadBalancing    bool              `json:"cross_zone_load_balancing"`
	Listener                  []AwsElbSpec      `json:"listener"`
	Tags                      map[string]string `json:"tags"`
	SourceSecurityGroupId     string            `json:"source_security_group_id"`
	Subnets                   []string          `json:"subnets"`
	AccessLogs                []AwsElbSpec      `json:"access_logs"`
	HealthCheck               []AwsElbSpec      `json:"health_check"`
	NamePrefix                string            `json:"name_prefix"`
	Arn                       string            `json:"arn"`
	Internal                  bool              `json:"internal"`
	Instances                 []string          `json:"instances"`
	SecurityGroups            []string          `json:"security_groups"`
	IdleTimeout               int               `json:"idle_timeout"`
	ConnectionDrainingTimeout int               `json:"connection_draining_timeout"`
	AvailabilityZones         []string          `json:"availability_zones"`
	SourceSecurityGroup       string            `json:"source_security_group"`
	ConnectionDraining        bool              `json:"connection_draining"`
	DnsName                   string            `json:"dns_name"`
	ZoneId                    string            `json:"zone_id"`
}



type AwsElbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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