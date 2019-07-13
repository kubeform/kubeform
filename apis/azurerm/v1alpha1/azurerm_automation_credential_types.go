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

type AzurermAutomationCredential struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationCredentialSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationCredentialStatus `json:"status,omitempty"`
}

type AzurermAutomationCredentialSpec struct {
	Password          string `json:"password"`
	Description       string `json:"description"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	AccountName       string `json:"account_name"`
	Username          string `json:"username"`
}



type AzurermAutomationCredentialStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutomationCredentialList is a list of AzurermAutomationCredentials
type AzurermAutomationCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationCredential CRD objects
	Items []AzurermAutomationCredential `json:"items,omitempty"`
}