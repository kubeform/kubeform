package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementAuthorizationServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAuthorizationServerSpec   `json:"spec,omitempty"`
	Status            ApiManagementAuthorizationServerStatus `json:"status,omitempty"`
}

type ApiManagementAuthorizationServerSpecTokenBodyParameter struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type ApiManagementAuthorizationServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ApiManagementName     string `json:"apiManagementName" tf:"api_management_name"`
	AuthorizationEndpoint string `json:"authorizationEndpoint" tf:"authorization_endpoint"`
	// +kubebuilder:validation:UniqueItems=true
	AuthorizationMethods []string `json:"authorizationMethods" tf:"authorization_methods"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BearerTokenSendingMethods []string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ClientAuthenticationMethod []string `json:"clientAuthenticationMethod,omitempty" tf:"client_authentication_method,omitempty"`
	ClientID                   string   `json:"clientID" tf:"client_id"`
	ClientRegistrationEndpoint string   `json:"clientRegistrationEndpoint" tf:"client_registration_endpoint"`
	// +optional
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret,omitempty"`
	// +optional
	DefaultScope string `json:"defaultScope,omitempty" tf:"default_scope,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	// +kubebuilder:validation:UniqueItems=true
	GrantTypes        []string `json:"grantTypes" tf:"grant_types"`
	Name              string   `json:"name" tf:"name"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceOwnerPassword string `json:"-" sensitive:"true" tf:"resource_owner_password,omitempty"`
	// +optional
	ResourceOwnerUsername string `json:"resourceOwnerUsername,omitempty" tf:"resource_owner_username,omitempty"`
	// +optional
	SupportState bool `json:"supportState,omitempty" tf:"support_state,omitempty"`
	// +optional
	TokenBodyParameter []ApiManagementAuthorizationServerSpecTokenBodyParameter `json:"tokenBodyParameter,omitempty" tf:"token_body_parameter,omitempty"`
	// +optional
	TokenEndpoint string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}



type ApiManagementAuthorizationServerStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ApiManagementAuthorizationServerSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementAuthorizationServerList is a list of ApiManagementAuthorizationServers
type ApiManagementAuthorizationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementAuthorizationServer CRD objects
	Items []ApiManagementAuthorizationServer `json:"items,omitempty"`
}