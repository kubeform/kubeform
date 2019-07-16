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

type Elb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElbSpec   `json:"spec,omitempty"`
	Status            ElbStatus `json:"status,omitempty"`
}

type ElbSpecAccessLogs struct {
	Bucket string `json:"bucket"`
	// +optional
	BucketPrefix string `json:"bucket_prefix,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Interval int `json:"interval,omitempty"`
}

type ElbSpecListener struct {
	InstancePort     int    `json:"instance_port"`
	InstanceProtocol string `json:"instance_protocol"`
	LbPort           int    `json:"lb_port"`
	LbProtocol       string `json:"lb_protocol"`
	// +optional
	SslCertificateId string `json:"ssl_certificate_id,omitempty"`
}

type ElbSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs *[]ElbSpec `json:"access_logs,omitempty"`
	// +optional
	ConnectionDraining bool `json:"connection_draining,omitempty"`
	// +optional
	ConnectionDrainingTimeout int `json:"connection_draining_timeout,omitempty"`
	// +optional
	CrossZoneLoadBalancing bool `json:"cross_zone_load_balancing,omitempty"`
	// +optional
	IdleTimeout int `json:"idle_timeout,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Listener []ElbSpec `json:"listener"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type ElbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
