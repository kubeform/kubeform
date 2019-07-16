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

type AzureadServicePrincipalPassword struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureadServicePrincipalPasswordSpec   `json:"spec,omitempty"`
	Status            AzureadServicePrincipalPasswordStatus `json:"status,omitempty"`
}

type AzureadServicePrincipalPasswordSpec struct {
	EndDate            string `json:"end_date"`
	ServicePrincipalId string `json:"service_principal_id"`
	Value              string `json:"value"`
}

type AzureadServicePrincipalPasswordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzureadServicePrincipalPasswordList is a list of AzureadServicePrincipalPasswords
type AzureadServicePrincipalPasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzureadServicePrincipalPassword CRD objects
	Items []AzureadServicePrincipalPassword `json:"items,omitempty"`
}
