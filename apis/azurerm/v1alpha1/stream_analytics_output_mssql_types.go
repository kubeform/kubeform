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

type StreamAnalyticsOutputMssql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputMssqlSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsOutputMssqlStatus `json:"status,omitempty"`
}

type StreamAnalyticsOutputMssqlSpec struct {
	Database               string `json:"database"`
	Name                   string `json:"name"`
	Password               string `json:"password"`
	ResourceGroupName      string `json:"resource_group_name"`
	Server                 string `json:"server"`
	StreamAnalyticsJobName string `json:"stream_analytics_job_name"`
	Table                  string `json:"table"`
	User                   string `json:"user"`
}

type StreamAnalyticsOutputMssqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsOutputMssqlList is a list of StreamAnalyticsOutputMssqls
type StreamAnalyticsOutputMssqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsOutputMssql CRD objects
	Items []StreamAnalyticsOutputMssql `json:"items,omitempty"`
}
