package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	EndDate string `json:"endDate" tf:"end_date"`
	// +optional
	KeyID              string `json:"keyID,omitempty" tf:"key_id,omitempty"`
	ServicePrincipalID string `json:"servicePrincipalID" tf:"service_principal_id"`
	// +optional
	StartDate string `json:"startDate,omitempty" tf:"start_date,omitempty"`
	Value     string `json:"-" sensitive:"true" tf:"value"`
}

type AzureadServicePrincipalPasswordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AzureadServicePrincipalPasswordSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
