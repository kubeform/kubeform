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

type CognitoIdentityProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityProviderSpec   `json:"spec,omitempty"`
	Status            CognitoIdentityProviderStatus `json:"status,omitempty"`
}

type CognitoIdentityProviderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AttributeMapping map[string]string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`
	// +optional
	IdpIdentifiers  []string          `json:"idpIdentifiers,omitempty" tf:"idp_identifiers,omitempty"`
	ProviderDetails map[string]string `json:"providerDetails" tf:"provider_details"`
	ProviderName    string            `json:"providerName" tf:"provider_name"`
	ProviderType    string            `json:"providerType" tf:"provider_type"`
	UserPoolID      string            `json:"userPoolID" tf:"user_pool_id"`
}

type CognitoIdentityProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CognitoIdentityProviderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoIdentityProviderList is a list of CognitoIdentityProviders
type CognitoIdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoIdentityProvider CRD objects
	Items []CognitoIdentityProvider `json:"items,omitempty"`
}
