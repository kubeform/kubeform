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

type StreamAnalyticsJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsJobSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsJobStatus `json:"status,omitempty"`
}

type StreamAnalyticsJobSpec struct {
	CompatibilityLevel                 string `json:"compatibility_level"`
	DataLocale                         string `json:"data_locale"`
	EventsLateArrivalMaxDelayInSeconds int    `json:"events_late_arrival_max_delay_in_seconds"`
	EventsOutOfOrderMaxDelayInSeconds  int    `json:"events_out_of_order_max_delay_in_seconds"`
	EventsOutOfOrderPolicy             string `json:"events_out_of_order_policy"`
	Location                           string `json:"location"`
	Name                               string `json:"name"`
	OutputErrorPolicy                  string `json:"output_error_policy"`
	ResourceGroupName                  string `json:"resource_group_name"`
	StreamingUnits                     int    `json:"streaming_units"`
	TransformationQuery                string `json:"transformation_query"`
}

type StreamAnalyticsJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsJobList is a list of StreamAnalyticsJobs
type StreamAnalyticsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsJob CRD objects
	Items []StreamAnalyticsJob `json:"items,omitempty"`
}
