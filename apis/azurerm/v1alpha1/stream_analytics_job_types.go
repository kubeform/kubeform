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

type StreamAnalyticsJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsJobSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsJobStatus `json:"status,omitempty"`
}

type StreamAnalyticsJobSpec struct {
	CompatibilityLevel                 string `json:"compatibilityLevel" tf:"compatibility_level"`
	DataLocale                         string `json:"dataLocale" tf:"data_locale"`
	EventsLateArrivalMaxDelayInSeconds int    `json:"eventsLateArrivalMaxDelayInSeconds" tf:"events_late_arrival_max_delay_in_seconds"`
	EventsOutOfOrderMaxDelayInSeconds  int    `json:"eventsOutOfOrderMaxDelayInSeconds" tf:"events_out_of_order_max_delay_in_seconds"`
	EventsOutOfOrderPolicy             string `json:"eventsOutOfOrderPolicy" tf:"events_out_of_order_policy"`
	Location                           string `json:"location" tf:"location"`
	Name                               string `json:"name" tf:"name"`
	OutputErrorPolicy                  string `json:"outputErrorPolicy" tf:"output_error_policy"`
	ResourceGroupName                  string `json:"resourceGroupName" tf:"resource_group_name"`
	StreamingUnits                     int    `json:"streamingUnits" tf:"streaming_units"`
	// +optional
	Tags                map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	TransformationQuery string                    `json:"transformationQuery" tf:"transformation_query"`
	ProviderRef         core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type StreamAnalyticsJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
