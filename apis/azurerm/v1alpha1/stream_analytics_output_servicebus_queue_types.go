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

type StreamAnalyticsOutputServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputServicebusQueueSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsOutputServicebusQueueStatus `json:"status,omitempty"`
}

type StreamAnalyticsOutputServicebusQueueSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`
	// +optional
	Format string `json:"format,omitempty" tf:"format,omitempty"`
	Type   string `json:"type" tf:"type"`
}

type StreamAnalyticsOutputServicebusQueueSpec struct {
	Name              string `json:"name" tf:"name"`
	QueueName         string `json:"queueName" tf:"queue_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization       []StreamAnalyticsOutputServicebusQueueSpecSerialization `json:"serialization" tf:"serialization"`
	ServicebusNamespace string                                                  `json:"servicebusNamespace" tf:"servicebus_namespace"`
	// Sensitive Data. Provide secret name which contains one value only
	SharedAccessPolicyKey  *core.LocalObjectReference `json:"sharedAccessPolicyKey" tf:"shared_access_policy_key"`
	SharedAccessPolicyName string                     `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`
	StreamAnalyticsJobName string                     `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
	ProviderRef            core.LocalObjectReference  `json:"providerRef" tf:"-"`
}

type StreamAnalyticsOutputServicebusQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
