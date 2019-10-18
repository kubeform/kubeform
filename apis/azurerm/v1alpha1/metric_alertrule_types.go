package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MetricAlertruleSpec   `json:"spec,omitempty"`
	Status            MetricAlertruleStatus `json:"status,omitempty"`
}

type MetricAlertruleSpecEmailAction struct {
	// +optional
	CustomEmails []string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`
	// +optional
	SendToServiceOwners bool `json:"sendToServiceOwners,omitempty" tf:"send_to_service_owners,omitempty"`
}

type MetricAlertruleSpecWebhookAction struct {
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	ServiceURI string            `json:"serviceURI" tf:"service_uri"`
}

type MetricAlertruleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Aggregation string `json:"aggregation" tf:"aggregation"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EmailAction []MetricAlertruleSpecEmailAction `json:"emailAction,omitempty" tf:"email_action,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Location          string `json:"location" tf:"location"`
	MetricName        string `json:"metricName" tf:"metric_name"`
	Name              string `json:"name" tf:"name"`
	Operator          string `json:"operator" tf:"operator"`
	Period            string `json:"period" tf:"period"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ResourceID        string `json:"resourceID" tf:"resource_id"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Threshold json.Number       `json:"threshold" tf:"threshold"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WebhookAction []MetricAlertruleSpecWebhookAction `json:"webhookAction,omitempty" tf:"webhook_action,omitempty"`
}

type MetricAlertruleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MetricAlertruleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MetricAlertruleList is a list of MetricAlertrules
type MetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MetricAlertrule CRD objects
	Items []MetricAlertrule `json:"items,omitempty"`
}
