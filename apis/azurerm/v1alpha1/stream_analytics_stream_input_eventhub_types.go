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

type StreamAnalyticsStreamInputEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsStreamInputEventhubSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsStreamInputEventhubStatus `json:"status,omitempty"`
}

type StreamAnalyticsStreamInputEventhubSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type StreamAnalyticsStreamInputEventhubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	EventhubConsumerGroupName string `json:"eventhubConsumerGroupName" tf:"eventhub_consumer_group_name"`
	EventhubName              string `json:"eventhubName" tf:"eventhub_name"`
	Name                      string `json:"name" tf:"name"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsStreamInputEventhubSpecSerialization `json:"serialization" tf:"serialization"`
	ServicebusNamespace    string                                                `json:"servicebusNamespace" tf:"servicebus_namespace"`
	SharedAccessPolicyName string                                                `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                                `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
}

type StreamAnalyticsStreamInputEventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
