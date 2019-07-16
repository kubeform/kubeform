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

type AutomationVariableString struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationVariableStringSpec   `json:"spec,omitempty"`
	Status            AutomationVariableStringStatus `json:"status,omitempty"`
}

type AutomationVariableStringSpec struct {
	AutomationAccountName string `json:"automation_account_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Encrypted         bool   `json:"encrypted,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Value string `json:"value,omitempty"`
}

type AutomationVariableStringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationVariableStringList is a list of AutomationVariableStrings
type AutomationVariableStringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationVariableString CRD objects
	Items []AutomationVariableString `json:"items,omitempty"`
}
