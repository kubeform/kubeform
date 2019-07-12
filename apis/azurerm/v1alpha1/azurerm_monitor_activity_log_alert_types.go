package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMonitorActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorActivityLogAlertSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorActivityLogAlertStatus `json:"status,omitempty"`
}

type AzurermMonitorActivityLogAlertSpecCriteria struct {
	Status           string `json:"status"`
	SubStatus        string `json:"sub_status"`
	Category         string `json:"category"`
	OperationName    string `json:"operation_name"`
	ResourceProvider string `json:"resource_provider"`
	ResourceType     string `json:"resource_type"`
	Caller           string `json:"caller"`
	Level            string `json:"level"`
	ResourceGroup    string `json:"resource_group"`
	ResourceId       string `json:"resource_id"`
}

type AzurermMonitorActivityLogAlertSpecAction struct {
	ActionGroupId     string            `json:"action_group_id"`
	WebhookProperties map[string]string `json:"webhook_properties"`
}

type AzurermMonitorActivityLogAlertSpec struct {
	Description       string                               `json:"description"`
	Enabled           bool                                 `json:"enabled"`
	Tags              map[string]string                    `json:"tags"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Scopes            []string                             `json:"scopes"`
	Criteria          []AzurermMonitorActivityLogAlertSpec `json:"criteria"`
	Action            []AzurermMonitorActivityLogAlertSpec `json:"action"`
}

type AzurermMonitorActivityLogAlertStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMonitorActivityLogAlertList is a list of AzurermMonitorActivityLogAlerts
type AzurermMonitorActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorActivityLogAlert CRD objects
	Items []AzurermMonitorActivityLogAlert `json:"items,omitempty"`
}
