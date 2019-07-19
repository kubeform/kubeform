package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Elb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElbSpec   `json:"spec,omitempty"`
	Status            ElbStatus `json:"status,omitempty"`
}

type ElbSpecAccessLogs struct {
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	BucketPrefix string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Interval int `json:"interval,omitempty" tf:"interval,omitempty"`
}

type ElbSpecHealthCheck struct {
	HealthyThreshold   int    `json:"healthyThreshold" tf:"healthy_threshold"`
	Interval           int    `json:"interval" tf:"interval"`
	Target             string `json:"target" tf:"target"`
	Timeout            int    `json:"timeout" tf:"timeout"`
	UnhealthyThreshold int    `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type ElbSpecListener struct {
	InstancePort     int    `json:"instancePort" tf:"instance_port"`
	InstanceProtocol string `json:"instanceProtocol" tf:"instance_protocol"`
	LbPort           int    `json:"lbPort" tf:"lb_port"`
	LbProtocol       string `json:"lbProtocol" tf:"lb_protocol"`
	// +optional
	SslCertificateID string `json:"sslCertificateID,omitempty" tf:"ssl_certificate_id,omitempty"`
}

type ElbSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs []ElbSpecAccessLogs `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	ConnectionDraining bool `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`
	// +optional
	ConnectionDrainingTimeout int `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout,omitempty"`
	// +optional
	CrossZoneLoadBalancing bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheck []ElbSpecHealthCheck `json:"healthCheck,omitempty" tf:"health_check,omitempty"`
	// +optional
	IdleTimeout int `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Instances []string `json:"instances,omitempty" tf:"instances,omitempty"`
	// +optional
	Internal bool `json:"internal,omitempty" tf:"internal,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Listener []ElbSpecListener `json:"listener" tf:"listener"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SourceSecurityGroup string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets,omitempty" tf:"subnets,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ElbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElbList is a list of Elbs
type ElbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Elb CRD objects
	Items []Elb `json:"items,omitempty"`
}
