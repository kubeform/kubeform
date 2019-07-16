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

type MonitorActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActivityLogAlertSpec   `json:"spec,omitempty"`
	Status            MonitorActivityLogAlertStatus `json:"status,omitempty"`
}

type MonitorActivityLogAlertSpecAction struct {
	ActionGroupId string `json:"action_group_id"`
	// +optional
	WebhookProperties map[string]string `json:"webhook_properties,omitempty"`
}

type MonitorActivityLogAlertSpecCriteria struct {
	// +optional
	Caller   string `json:"caller,omitempty"`
	Category string `json:"category"`
	// +optional
	Level string `json:"level,omitempty"`
	// +optional
	OperationName string `json:"operation_name,omitempty"`
	// +optional
	ResourceGroup string `json:"resource_group,omitempty"`
	// +optional
	ResourceId string `json:"resource_id,omitempty"`
	// +optional
	ResourceProvider string `json:"resource_provider,omitempty"`
	// +optional
	ResourceType string `json:"resource_type,omitempty"`
	// +optional
	Status string `json:"status,omitempty"`
	// +optional
	SubStatus string `json:"sub_status,omitempty"`
}

type MonitorActivityLogAlertSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Action *[]MonitorActivityLogAlertSpec `json:"action,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Criteria []MonitorActivityLogAlertSpec `json:"criteria"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes"`
}

type MonitorActivityLogAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorActivityLogAlertList is a list of MonitorActivityLogAlerts
type MonitorActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorActivityLogAlert CRD objects
	Items []MonitorActivityLogAlert `json:"items,omitempty"`
}
