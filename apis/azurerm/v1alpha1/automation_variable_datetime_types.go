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

type AutomationVariableDatetime struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationVariableDatetimeSpec   `json:"spec,omitempty"`
	Status            AutomationVariableDatetimeStatus `json:"status,omitempty"`
}

type AutomationVariableDatetimeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AutomationAccountName string `json:"automationAccountName" tf:"automation_account_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Encrypted         bool   `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type AutomationVariableDatetimeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutomationVariableDatetimeSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationVariableDatetimeList is a list of AutomationVariableDatetimes
type AutomationVariableDatetimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationVariableDatetime CRD objects
	Items []AutomationVariableDatetime `json:"items,omitempty"`
}
