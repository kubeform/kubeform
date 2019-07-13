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

type AzurermAutomationDscNodeconfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationDscNodeconfigurationSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationDscNodeconfigurationStatus `json:"status,omitempty"`
}

type AzurermAutomationDscNodeconfigurationSpec struct {
	Name                  string `json:"name"`
	AutomationAccountName string `json:"automation_account_name"`
	ResourceGroupName     string `json:"resource_group_name"`
	ContentEmbedded       string `json:"content_embedded"`
	ConfigurationName     string `json:"configuration_name"`
}



type AzurermAutomationDscNodeconfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutomationDscNodeconfigurationList is a list of AzurermAutomationDscNodeconfigurations
type AzurermAutomationDscNodeconfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationDscNodeconfiguration CRD objects
	Items []AzurermAutomationDscNodeconfiguration `json:"items,omitempty"`
}