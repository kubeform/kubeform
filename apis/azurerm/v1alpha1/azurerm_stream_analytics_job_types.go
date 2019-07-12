package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStreamAnalyticsJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsJobSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsJobStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsJobSpec struct {
	DataLocale                         string            `json:"data_locale"`
	EventsOutOfOrderMaxDelayInSeconds  int               `json:"events_out_of_order_max_delay_in_seconds"`
	EventsOutOfOrderPolicy             string            `json:"events_out_of_order_policy"`
	OutputErrorPolicy                  string            `json:"output_error_policy"`
	TransformationQuery                string            `json:"transformation_query"`
	Tags                               map[string]string `json:"tags"`
	Name                               string            `json:"name"`
	ResourceGroupName                  string            `json:"resource_group_name"`
	Location                           string            `json:"location"`
	CompatibilityLevel                 string            `json:"compatibility_level"`
	EventsLateArrivalMaxDelayInSeconds int               `json:"events_late_arrival_max_delay_in_seconds"`
	StreamingUnits                     int               `json:"streaming_units"`
	JobId                              string            `json:"job_id"`
}

type AzurermStreamAnalyticsJobStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStreamAnalyticsJobList is a list of AzurermStreamAnalyticsJobs
type AzurermStreamAnalyticsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsJob CRD objects
	Items []AzurermStreamAnalyticsJob `json:"items,omitempty"`
}
