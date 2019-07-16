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

type AutomationCredential struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationCredentialSpec   `json:"spec,omitempty"`
	Status            AutomationCredentialStatus `json:"status,omitempty"`
}

type AutomationCredentialSpec struct {
	AccountName string `json:"account_name"`
	// +optional
	Description       string `json:"description,omitempty"`
	Name              string `json:"name"`
	Password          string `json:"password"`
	ResourceGroupName string `json:"resource_group_name"`
	Username          string `json:"username"`
}

type AutomationCredentialStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationCredentialList is a list of AutomationCredentials
type AutomationCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationCredential CRD objects
	Items []AutomationCredential `json:"items,omitempty"`
}
