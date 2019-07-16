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

type StreamAnalyticsOutputServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputServicebusQueueSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsOutputServicebusQueueStatus `json:"status,omitempty"`
}

type StreamAnalyticsOutputServicebusQueueSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"field_delimiter,omitempty"`
	// +optional
	Format string `json:"format,omitempty"`
	Type   string `json:"type"`
}

type StreamAnalyticsOutputServicebusQueueSpec struct {
	Name              string `json:"name"`
	QueueName         string `json:"queue_name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsOutputServicebusQueueSpec `json:"serialization"`
	ServicebusNamespace    string                                     `json:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                     `json:"shared_access_policy_key"`
	SharedAccessPolicyName string                                     `json:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                     `json:"stream_analytics_job_name"`
}

type StreamAnalyticsOutputServicebusQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsOutputServicebusQueueList is a list of StreamAnalyticsOutputServicebusQueues
type StreamAnalyticsOutputServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsOutputServicebusQueue CRD objects
	Items []StreamAnalyticsOutputServicebusQueue `json:"items,omitempty"`
}
