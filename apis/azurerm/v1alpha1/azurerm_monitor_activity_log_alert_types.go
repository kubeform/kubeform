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
	ResourceId       string `json:"resource_id"`
	SubStatus        string `json:"sub_status"`
	ResourceProvider string `json:"resource_provider"`
	ResourceGroup    string `json:"resource_group"`
	Caller           string `json:"caller"`
	Level            string `json:"level"`
	ResourceType     string `json:"resource_type"`
	Status           string `json:"status"`
	Category         string `json:"category"`
	OperationName    string `json:"operation_name"`
}

type AzurermMonitorActivityLogAlertSpecAction struct {
	WebhookProperties map[string]string `json:"webhook_properties"`
	ActionGroupId     string            `json:"action_group_id"`
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
