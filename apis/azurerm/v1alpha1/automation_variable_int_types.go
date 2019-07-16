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

type AutomationVariableInt struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationVariableIntSpec   `json:"spec,omitempty"`
	Status            AutomationVariableIntStatus `json:"status,omitempty"`
}

type AutomationVariableIntSpec struct {
	AutomationAccountName string `json:"automation_account_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Encrypted         bool   `json:"encrypted,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Value int `json:"value,omitempty"`
}

type AutomationVariableIntStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationVariableIntList is a list of AutomationVariableInts
type AutomationVariableIntList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationVariableInt CRD objects
	Items []AutomationVariableInt `json:"items,omitempty"`
}
