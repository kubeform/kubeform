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

type AzurermStreamAnalyticsJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsJobSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsJobStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsJobSpec struct {
	Name                               string            `json:"name"`
	ResourceGroupName                  string            `json:"resource_group_name"`
	EventsLateArrivalMaxDelayInSeconds int               `json:"events_late_arrival_max_delay_in_seconds"`
	EventsOutOfOrderPolicy             string            `json:"events_out_of_order_policy"`
	StreamingUnits                     int               `json:"streaming_units"`
	TransformationQuery                string            `json:"transformation_query"`
	JobId                              string            `json:"job_id"`
	Tags                               map[string]string `json:"tags"`
	Location                           string            `json:"location"`
	CompatibilityLevel                 string            `json:"compatibility_level"`
	DataLocale                         string            `json:"data_locale"`
	EventsOutOfOrderMaxDelayInSeconds  int               `json:"events_out_of_order_max_delay_in_seconds"`
	OutputErrorPolicy                  string            `json:"output_error_policy"`
}



type AzurermStreamAnalyticsJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStreamAnalyticsJobList is a list of AzurermStreamAnalyticsJobs
type AzurermStreamAnalyticsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsJob CRD objects
	Items []AzurermStreamAnalyticsJob `json:"items,omitempty"`
}