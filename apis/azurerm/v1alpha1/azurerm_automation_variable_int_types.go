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

type AzurermAutomationVariableInt struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationVariableIntSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationVariableIntStatus `json:"status,omitempty"`
}

type AzurermAutomationVariableIntSpec struct {
	Value                 int    `json:"value"`
	ResourceGroupName     string `json:"resource_group_name"`
	Name                  string `json:"name"`
	AutomationAccountName string `json:"automation_account_name"`
	Description           string `json:"description"`
	Encrypted             bool   `json:"encrypted"`
}

type AzurermAutomationVariableIntStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutomationVariableIntList is a list of AzurermAutomationVariableInts
type AzurermAutomationVariableIntList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationVariableInt CRD objects
	Items []AzurermAutomationVariableInt `json:"items,omitempty"`
}
