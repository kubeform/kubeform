package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAutomationVariableString struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationVariableStringSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationVariableStringStatus `json:"status,omitempty"`
}

type AzurermAutomationVariableStringSpec struct {
	ResourceGroupName     string `json:"resource_group_name"`
	Name                  string `json:"name"`
	AutomationAccountName string `json:"automation_account_name"`
	Description           string `json:"description"`
	Encrypted             bool   `json:"encrypted"`
	Value                 string `json:"value"`
}



type AzurermAutomationVariableStringStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAutomationVariableStringList is a list of AzurermAutomationVariableStrings
type AzurermAutomationVariableStringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationVariableString CRD objects
	Items []AzurermAutomationVariableString `json:"items,omitempty"`
}