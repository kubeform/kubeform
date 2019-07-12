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

type AzurermAutomationModule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationModuleSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationModuleStatus `json:"status,omitempty"`
}

type AzurermAutomationModuleSpecModuleLinkHash struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type AzurermAutomationModuleSpecModuleLink struct {
	Uri  string                                  `json:"uri"`
	Hash []AzurermAutomationModuleSpecModuleLink `json:"hash"`
}

type AzurermAutomationModuleSpec struct {
	Name                  string                        `json:"name"`
	AutomationAccountName string                        `json:"automation_account_name"`
	ResourceGroupName     string                        `json:"resource_group_name"`
	ModuleLink            []AzurermAutomationModuleSpec `json:"module_link"`
}

type AzurermAutomationModuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutomationModuleList is a list of AzurermAutomationModules
type AzurermAutomationModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationModule CRD objects
	Items []AzurermAutomationModule `json:"items,omitempty"`
}
