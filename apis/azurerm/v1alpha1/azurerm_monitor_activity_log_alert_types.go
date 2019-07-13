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

type AzurermMonitorActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorActivityLogAlertSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorActivityLogAlertStatus `json:"status,omitempty"`
}

type AzurermMonitorActivityLogAlertSpecCriteria struct {
	OperationName    string `json:"operation_name"`
	Level            string `json:"level"`
	ResourceProvider string `json:"resource_provider"`
	Status           string `json:"status"`
	SubStatus        string `json:"sub_status"`
	Category         string `json:"category"`
	Caller           string `json:"caller"`
	ResourceType     string `json:"resource_type"`
	ResourceGroup    string `json:"resource_group"`
	ResourceId       string `json:"resource_id"`
}

type AzurermMonitorActivityLogAlertSpecAction struct {
	ActionGroupId     string            `json:"action_group_id"`
	WebhookProperties map[string]string `json:"webhook_properties"`
}

type AzurermMonitorActivityLogAlertSpec struct {
	Scopes            []string                             `json:"scopes"`
	Criteria          []AzurermMonitorActivityLogAlertSpec `json:"criteria"`
	Action            []AzurermMonitorActivityLogAlertSpec `json:"action"`
	Description       string                               `json:"description"`
	Enabled           bool                                 `json:"enabled"`
	Tags              map[string]string                    `json:"tags"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
}



type AzurermMonitorActivityLogAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMonitorActivityLogAlertList is a list of AzurermMonitorActivityLogAlerts
type AzurermMonitorActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorActivityLogAlert CRD objects
	Items []AzurermMonitorActivityLogAlert `json:"items,omitempty"`
}