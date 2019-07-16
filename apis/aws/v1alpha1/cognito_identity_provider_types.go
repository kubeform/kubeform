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

type CognitoIdentityProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityProviderSpec   `json:"spec,omitempty"`
	Status            CognitoIdentityProviderStatus `json:"status,omitempty"`
}

type CognitoIdentityProviderSpec struct {
	// +optional
	AttributeMapping map[string]string `json:"attribute_mapping,omitempty"`
	// +optional
	IdpIdentifiers  []string          `json:"idp_identifiers,omitempty"`
	ProviderDetails map[string]string `json:"provider_details"`
	ProviderName    string            `json:"provider_name"`
	ProviderType    string            `json:"provider_type"`
	UserPoolId      string            `json:"user_pool_id"`
}

type CognitoIdentityProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
