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

type CognitoIdentityPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolSpec   `json:"spec,omitempty"`
	Status            CognitoIdentityPoolStatus `json:"status,omitempty"`
}

type CognitoIdentityPoolSpecCognitoIdentityProviders struct {
	// +optional
	ClientId string `json:"client_id,omitempty"`
	// +optional
	ProviderName string `json:"provider_name,omitempty"`
	// +optional
	ServerSideTokenCheck bool `json:"server_side_token_check,omitempty"`
}

type CognitoIdentityPoolSpec struct {
	// +optional
	AllowUnauthenticatedIdentities bool `json:"allow_unauthenticated_identities,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CognitoIdentityProviders *[]CognitoIdentityPoolSpec `json:"cognito_identity_providers,omitempty"`
	// +optional
	DeveloperProviderName string `json:"developer_provider_name,omitempty"`
	IdentityPoolName      string `json:"identity_pool_name"`
	// +optional
	OpenidConnectProviderArns []string `json:"openid_connect_provider_arns,omitempty"`
	// +optional
	SamlProviderArns []string `json:"saml_provider_arns,omitempty"`
	// +optional
	SupportedLoginProviders map[string]string `json:"supported_login_providers,omitempty"`
}

type CognitoIdentityPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoIdentityPoolList is a list of CognitoIdentityPools
type CognitoIdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoIdentityPool CRD objects
	Items []CognitoIdentityPool `json:"items,omitempty"`
}
