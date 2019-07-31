package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AutomationDscNodeconfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationDscNodeconfigurationSpec   `json:"spec,omitempty"`
	Status            AutomationDscNodeconfigurationStatus `json:"status,omitempty"`
}

type AutomationDscNodeconfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AutomationAccountName string `json:"automationAccountName" tf:"automation_account_name"`
	ContentEmbedded       string `json:"contentEmbedded" tf:"content_embedded"`
	Name                  string `json:"name" tf:"name"`
	ResourceGroupName     string `json:"resourceGroupName" tf:"resource_group_name"`
}

type AutomationDscNodeconfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationDscNodeconfigurationList is a list of AutomationDscNodeconfigurations
type AutomationDscNodeconfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationDscNodeconfiguration CRD objects
	Items []AutomationDscNodeconfiguration `json:"items,omitempty"`
}
