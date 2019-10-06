package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StreamAnalyticsOutputEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputEventhubSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsOutputEventhubStatus `json:"status,omitempty"`
}

type StreamAnalyticsOutputEventhubSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`
	// +optional
	Format string `json:"format,omitempty" tf:"format,omitempty"`
	Type   string `json:"type" tf:"type"`
}

type StreamAnalyticsOutputEventhubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	EventhubName      string `json:"eventhubName" tf:"eventhub_name"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsOutputEventhubSpecSerialization `json:"serialization" tf:"serialization"`
	ServicebusNamespace    string                                           `json:"servicebusNamespace" tf:"servicebus_namespace"`
	SharedAccessPolicyKey  string                                           `json:"-" sensitive:"true" tf:"shared_access_policy_key"`
	SharedAccessPolicyName string                                           `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`
	StreamAnalyticsJobName string                                           `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
}

type StreamAnalyticsOutputEventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StreamAnalyticsOutputEventhubSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsOutputEventhubList is a list of StreamAnalyticsOutputEventhubs
type StreamAnalyticsOutputEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsOutputEventhub CRD objects
	Items []StreamAnalyticsOutputEventhub `json:"items,omitempty"`
}
