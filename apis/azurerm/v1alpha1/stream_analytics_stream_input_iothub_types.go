package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type StreamAnalyticsStreamInputIothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsStreamInputIothubSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsStreamInputIothubStatus `json:"status,omitempty"`
}

type StreamAnalyticsStreamInputIothubSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type StreamAnalyticsStreamInputIothubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Endpoint                  string `json:"endpoint" tf:"endpoint"`
	EventhubConsumerGroupName string `json:"eventhubConsumerGroupName" tf:"eventhub_consumer_group_name"`
	IothubNamespace           string `json:"iothubNamespace" tf:"iothub_namespace"`
	Name                      string `json:"name" tf:"name"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsStreamInputIothubSpecSerialization `json:"serialization" tf:"serialization"`
	SharedAccessPolicyKey  string                                              `json:"-" sensitive:"true" tf:"shared_access_policy_key"`
	SharedAccessPolicyName string                                              `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                              `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
}

type StreamAnalyticsStreamInputIothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StreamAnalyticsStreamInputIothubSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
