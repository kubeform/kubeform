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

type CognitoIdentityPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolSpec   `json:"spec,omitempty"`
	Status            CognitoIdentityPoolStatus `json:"status,omitempty"`
}

type CognitoIdentityPoolSpecCognitoIdentityProviders struct {
	// +optional
	ClientID string `json:"clientID,omitempty" tf:"client_id,omitempty"`
	// +optional
	ProviderName string `json:"providerName,omitempty" tf:"provider_name,omitempty"`
	// +optional
	ServerSideTokenCheck bool `json:"serverSideTokenCheck,omitempty" tf:"server_side_token_check,omitempty"`
}

type CognitoIdentityPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowUnauthenticatedIdentities bool `json:"allowUnauthenticatedIdentities,omitempty" tf:"allow_unauthenticated_identities,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CognitoIdentityProviders []CognitoIdentityPoolSpecCognitoIdentityProviders `json:"cognitoIdentityProviders,omitempty" tf:"cognito_identity_providers,omitempty"`
	// +optional
	DeveloperProviderName string `json:"developerProviderName,omitempty" tf:"developer_provider_name,omitempty"`
	IdentityPoolName      string `json:"identityPoolName" tf:"identity_pool_name"`
	// +optional
	OpenidConnectProviderArns []string `json:"openidConnectProviderArns,omitempty" tf:"openid_connect_provider_arns,omitempty"`
	// +optional
	SamlProviderArns []string `json:"samlProviderArns,omitempty" tf:"saml_provider_arns,omitempty"`
	// +optional
	SupportedLoginProviders map[string]string `json:"supportedLoginProviders,omitempty" tf:"supported_login_providers,omitempty"`
}

type CognitoIdentityPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CognitoIdentityPoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
