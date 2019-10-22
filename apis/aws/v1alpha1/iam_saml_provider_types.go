package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type IamSamlProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamSamlProviderSpec   `json:"spec,omitempty"`
	Status            IamSamlProviderStatus `json:"status,omitempty"`
}

type IamSamlProviderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                  string `json:"arn,omitempty" tf:"arn,omitempty"`
	Name                 string `json:"name" tf:"name"`
	SamlMetadataDocument string `json:"samlMetadataDocument" tf:"saml_metadata_document"`
	// +optional
	ValidUntil string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type IamSamlProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IamSamlProviderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamSamlProviderList is a list of IamSamlProviders
type IamSamlProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamSamlProvider CRD objects
	Items []IamSamlProvider `json:"items,omitempty"`
}
