package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	Algorithm string `json:"algorithm" tf:"algorithm"`
	Value     string `json:"value" tf:"value"`
}

type AutomationModuleSpecModuleLink struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Hash []AutomationModuleSpecModuleLinkHash `json:"hash,omitempty" tf:"hash,omitempty"`
	Uri  string                               `json:"uri" tf:"uri"`
}

type AutomationModuleSpec struct {
	ProviderRef           core.LocalObjectReference `json:"providerRef" tf:"-"`
	AutomationAccountName string                    `json:"automationAccountName" tf:"automation_account_name"`
	// +kubebuilder:validation:MaxItems=1
	ModuleLink        []AutomationModuleSpecModuleLink `json:"moduleLink" tf:"module_link"`
	Name              string                           `json:"name" tf:"name"`
	ResourceGroupName string                           `json:"resourceGroupName" tf:"resource_group_name"`
}

type AutomationModuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
