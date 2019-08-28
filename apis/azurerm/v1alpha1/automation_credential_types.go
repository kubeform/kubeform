package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	AccountName string `json:"accountName" tf:"account_name"`
	// +optional
	Description       string `json:"description,omitempty" tf:"description,omitempty"`
	Name              string `json:"name" tf:"name"`
	Password          string `json:"-" sensitive:"true" tf:"password"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	Username          string `json:"username" tf:"username"`
}



type AutomationCredentialStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AutomationCredentialSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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