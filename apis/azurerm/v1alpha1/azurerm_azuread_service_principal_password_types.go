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

type AzurermAzureadServicePrincipalPassword struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAzureadServicePrincipalPasswordSpec   `json:"spec,omitempty"`
	Status            AzurermAzureadServicePrincipalPasswordStatus `json:"status,omitempty"`
}

type AzurermAzureadServicePrincipalPasswordSpec struct {
	KeyId              string `json:"key_id"`
	Value              string `json:"value"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	ServicePrincipalId string `json:"service_principal_id"`
}



type AzurermAzureadServicePrincipalPasswordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAzureadServicePrincipalPasswordList is a list of AzurermAzureadServicePrincipalPasswords
type AzurermAzureadServicePrincipalPasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAzureadServicePrincipalPassword CRD objects
	Items []AzurermAzureadServicePrincipalPassword `json:"items,omitempty"`
}