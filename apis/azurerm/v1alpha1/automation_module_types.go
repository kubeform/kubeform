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

type AutomationModule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationModuleSpec   `json:"spec,omitempty"`
	Status            AutomationModuleStatus `json:"status,omitempty"`
}

type AutomationModuleSpecModuleLinkHash struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type AutomationModuleSpecModuleLink struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Hash *[]AutomationModuleSpecModuleLink `json:"hash,omitempty"`
	Uri  string                            `json:"uri"`
}

type AutomationModuleSpec struct {
	AutomationAccountName string `json:"automation_account_name"`
	// +kubebuilder:validation:MaxItems=1
	ModuleLink        []AutomationModuleSpec `json:"module_link"`
	Name              string                 `json:"name"`
	ResourceGroupName string                 `json:"resource_group_name"`
}

type AutomationModuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationModuleList is a list of AutomationModules
type AutomationModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationModule CRD objects
	Items []AutomationModule `json:"items,omitempty"`
}
