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

type AzurermAutomationVariableDatetime struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationVariableDatetimeSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationVariableDatetimeStatus `json:"status,omitempty"`
}

type AzurermAutomationVariableDatetimeSpec struct {
	Encrypted             bool   `json:"encrypted"`
	Value                 string `json:"value"`
	ResourceGroupName     string `json:"resource_group_name"`
	Name                  string `json:"name"`
	AutomationAccountName string `json:"automation_account_name"`
	Description           string `json:"description"`
}

type AzurermAutomationVariableDatetimeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutomationVariableDatetimeList is a list of AzurermAutomationVariableDatetimes
type AzurermAutomationVariableDatetimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationVariableDatetime CRD objects
	Items []AzurermAutomationVariableDatetime `json:"items,omitempty"`
}
