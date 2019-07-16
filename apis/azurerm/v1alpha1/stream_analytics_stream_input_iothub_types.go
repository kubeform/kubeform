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

type StreamAnalyticsStreamInputIothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsStreamInputIothubSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsStreamInputIothubStatus `json:"status,omitempty"`
}

type StreamAnalyticsStreamInputIothubSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"field_delimiter,omitempty"`
	Type           string `json:"type"`
}

type StreamAnalyticsStreamInputIothubSpec struct {
	Endpoint                  string `json:"endpoint"`
	EventhubConsumerGroupName string `json:"eventhub_consumer_group_name"`
	IothubNamespace           string `json:"iothub_namespace"`
	Name                      string `json:"name"`
	ResourceGroupName         string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsStreamInputIothubSpec `json:"serialization"`
	SharedAccessPolicyKey  string                                 `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                 `json:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                 `json:"stream_analytics_job_name"`
}

type StreamAnalyticsStreamInputIothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsStreamInputIothubList is a list of StreamAnalyticsStreamInputIothubs
type StreamAnalyticsStreamInputIothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsStreamInputIothub CRD objects
	Items []StreamAnalyticsStreamInputIothub `json:"items,omitempty"`
}
