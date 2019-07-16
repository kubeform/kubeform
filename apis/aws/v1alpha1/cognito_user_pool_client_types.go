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

type CognitoUserPoolClient struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolClientSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolClientStatus `json:"status,omitempty"`
}

type CognitoUserPoolClientSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=3
	// +kubebuilder:validation:UniqueItems=true
	AllowedOauthFlows []string `json:"allowed_oauth_flows,omitempty"`
	// +optional
	AllowedOauthFlowsUserPoolClient bool `json:"allowed_oauth_flows_user_pool_client,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	AllowedOauthScopes []string `json:"allowed_oauth_scopes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	CallbackUrls []string `json:"callback_urls,omitempty"`
	// +optional
	DefaultRedirectUri string `json:"default_redirect_uri,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ExplicitAuthFlows []string `json:"explicit_auth_flows,omitempty"`
	// +optional
	GenerateSecret bool `json:"generate_secret,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	LogoutUrls []string `json:"logout_urls,omitempty"`
	Name       string   `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ReadAttributes []string `json:"read_attributes,omitempty"`
	// +optional
	RefreshTokenValidity int `json:"refresh_token_validity,omitempty"`
	// +optional
	SupportedIdentityProviders []string `json:"supported_identity_providers,omitempty"`
	UserPoolId                 string   `json:"user_pool_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WriteAttributes []string `json:"write_attributes,omitempty"`
}

type CognitoUserPoolClientStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoUserPoolClientList is a list of CognitoUserPoolClients
type CognitoUserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoUserPoolClient CRD objects
	Items []CognitoUserPoolClient `json:"items,omitempty"`
}
