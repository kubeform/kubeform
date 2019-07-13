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

type AzurermMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            AzurermMetricAlertruleStatus `json:"status,omitempty"`
}

type AzurermMetricAlertruleSpecEmailAction struct {
	SendToServiceOwners bool     `json:"send_to_service_owners"`
	CustomEmails        []string `json:"custom_emails"`
}

type AzurermMetricAlertruleSpecWebhookAction struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermMetricAlertruleSpec struct {
	Description       string                       `json:"description"`
	ResourceId        string                       `json:"resource_id"`
	MetricName        string                       `json:"metric_name"`
	Location          string                       `json:"location"`
	Period            string                       `json:"period"`
	Aggregation       string                       `json:"aggregation"`
	EmailAction       []AzurermMetricAlertruleSpec `json:"email_action"`
	WebhookAction     []AzurermMetricAlertruleSpec `json:"webhook_action"`
	Enabled           bool                         `json:"enabled"`
	Tags              map[string]string            `json:"tags"`
	Name              string                       `json:"name"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Operator          string                       `json:"operator"`
	Threshold         float64                      `json:"threshold"`
}



type AzurermMetricAlertruleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMetricAlertruleList is a list of AzurermMetricAlertrules
type AzurermMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMetricAlertrule CRD objects
	Items []AzurermMetricAlertrule `json:"items,omitempty"`
}