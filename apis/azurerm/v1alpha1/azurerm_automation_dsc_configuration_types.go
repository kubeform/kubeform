package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAutomationDscConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationDscConfigurationSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationDscConfigurationStatus `json:"status,omitempty"`
}

type AzurermAutomationDscConfigurationSpec struct {
	ResourceGroupName     string `json:"resource_group_name"`
	Location              string `json:"location"`
	LogVerbose            bool   `json:"log_verbose"`
	Description           string `json:"description"`
	State                 string `json:"state"`
	Name                  string `json:"name"`
	AutomationAccountName string `json:"automation_account_name"`
	ContentEmbedded       string `json:"content_embedded"`
}

type AzurermAutomationDscConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAutomationDscConfigurationList is a list of AzurermAutomationDscConfigurations
type AzurermAutomationDscConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationDscConfiguration CRD objects
	Items []AzurermAutomationDscConfiguration `json:"items,omitempty"`
}
