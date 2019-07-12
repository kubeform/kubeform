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

type AzurermAzureadServicePrincipal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAzureadServicePrincipalSpec   `json:"spec,omitempty"`
	Status            AzurermAzureadServicePrincipalStatus `json:"status,omitempty"`
}

type AzurermAzureadServicePrincipalSpec struct {
	ApplicationId string `json:"application_id"`
	DisplayName   string `json:"display_name"`
}

type AzurermAzureadServicePrincipalStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAzureadServicePrincipalList is a list of AzurermAzureadServicePrincipals
type AzurermAzureadServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAzureadServicePrincipal CRD objects
	Items []AzurermAzureadServicePrincipal `json:"items,omitempty"`
}
