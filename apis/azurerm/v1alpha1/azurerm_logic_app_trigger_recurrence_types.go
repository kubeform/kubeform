package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLogicAppTriggerRecurrence struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppTriggerRecurrenceSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppTriggerRecurrenceStatus `json:"status,omitempty"`
}

type AzurermLogicAppTriggerRecurrenceSpec struct {
	Name       string `json:"name"`
	LogicAppId string `json:"logic_app_id"`
	Frequency  string `json:"frequency"`
	Interval   int    `json:"interval"`
}



type AzurermLogicAppTriggerRecurrenceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLogicAppTriggerRecurrenceList is a list of AzurermLogicAppTriggerRecurrences
type AzurermLogicAppTriggerRecurrenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppTriggerRecurrence CRD objects
	Items []AzurermLogicAppTriggerRecurrence `json:"items,omitempty"`
}