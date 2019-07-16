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

type Lb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSpec   `json:"spec,omitempty"`
	Status            LbStatus `json:"status,omitempty"`
}

type LbSpecAccessLogs struct {
	Bucket string `json:"bucket"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

type LbSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs *[]LbSpec `json:"access_logs,omitempty"`
	// +optional
	EnableCrossZoneLoadBalancing bool `json:"enable_cross_zone_load_balancing,omitempty"`
	// +optional
	EnableDeletionProtection bool `json:"enable_deletion_protection,omitempty"`
	// +optional
	EnableHttp2 bool `json:"enable_http2,omitempty"`
	// +optional
	IdleTimeout int `json:"idle_timeout,omitempty"`
	// +optional
	LoadBalancerType string `json:"load_balancer_type,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type LbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbList is a list of Lbs
type LbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Lb CRD objects
	Items []Lb `json:"items,omitempty"`
}
