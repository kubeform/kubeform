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

type StreamAnalyticsStreamInputEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsStreamInputEventhubSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsStreamInputEventhubStatus `json:"status,omitempty"`
}

type StreamAnalyticsStreamInputEventhubSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"field_delimiter,omitempty"`
	Type           string `json:"type"`
}

type StreamAnalyticsStreamInputEventhubSpec struct {
	EventhubConsumerGroupName string `json:"eventhub_consumer_group_name"`
	EventhubName              string `json:"eventhub_name"`
	Name                      string `json:"name"`
	ResourceGroupName         string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsStreamInputEventhubSpec `json:"serialization"`
	ServicebusNamespace    string                                   `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                   `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                   `json:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                   `json:"stream_analytics_job_name"`
}

type StreamAnalyticsStreamInputEventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsStreamInputEventhubList is a list of StreamAnalyticsStreamInputEventhubs
type StreamAnalyticsStreamInputEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsStreamInputEventhub CRD objects
	Items []StreamAnalyticsStreamInputEventhub `json:"items,omitempty"`
}
