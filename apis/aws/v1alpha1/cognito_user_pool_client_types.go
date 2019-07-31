package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=3
	// +kubebuilder:validation:UniqueItems=true
	AllowedOauthFlows []string `json:"allowedOauthFlows,omitempty" tf:"allowed_oauth_flows,omitempty"`
	// +optional
	AllowedOauthFlowsUserPoolClient bool `json:"allowedOauthFlowsUserPoolClient,omitempty" tf:"allowed_oauth_flows_user_pool_client,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	AllowedOauthScopes []string `json:"allowedOauthScopes,omitempty" tf:"allowed_oauth_scopes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	CallbackUrls []string `json:"callbackUrls,omitempty" tf:"callback_urls,omitempty"`
	// +optional
	DefaultRedirectURI string `json:"defaultRedirectURI,omitempty" tf:"default_redirect_uri,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ExplicitAuthFlows []string `json:"explicitAuthFlows,omitempty" tf:"explicit_auth_flows,omitempty"`
	// +optional
	GenerateSecret bool `json:"generateSecret,omitempty" tf:"generate_secret,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	LogoutUrls []string `json:"logoutUrls,omitempty" tf:"logout_urls,omitempty"`
	Name       string   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ReadAttributes []string `json:"readAttributes,omitempty" tf:"read_attributes,omitempty"`
	// +optional
	RefreshTokenValidity int `json:"refreshTokenValidity,omitempty" tf:"refresh_token_validity,omitempty"`
	// +optional
	SupportedIdentityProviders []string `json:"supportedIdentityProviders,omitempty" tf:"supported_identity_providers,omitempty"`
	UserPoolID                 string   `json:"userPoolID" tf:"user_pool_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WriteAttributes []string `json:"writeAttributes,omitempty" tf:"write_attributes,omitempty"`
}

type CognitoUserPoolClientStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
