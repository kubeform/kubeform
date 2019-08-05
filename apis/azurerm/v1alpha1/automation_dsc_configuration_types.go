package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AutomationDscConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationDscConfigurationSpec   `json:"spec,omitempty"`
	Status            AutomationDscConfigurationStatus `json:"status,omitempty"`
}

type AutomationDscConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AutomationAccountName string `json:"automationAccountName" tf:"automation_account_name"`
	ContentEmbedded       string `json:"contentEmbedded" tf:"content_embedded"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Location    string `json:"location" tf:"location"`
	// +optional
	LogVerbose        bool   `json:"logVerbose,omitempty" tf:"log_verbose,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
}

type AutomationDscConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutomationDscConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationDscConfigurationList is a list of AutomationDscConfigurations
type AutomationDscConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationDscConfiguration CRD objects
	Items []AutomationDscConfiguration `json:"items,omitempty"`
}
