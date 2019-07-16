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

type ApiManagementAuthorizationServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAuthorizationServerSpec   `json:"spec,omitempty"`
	Status            ApiManagementAuthorizationServerStatus `json:"status,omitempty"`
}

type ApiManagementAuthorizationServerSpecTokenBodyParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ApiManagementAuthorizationServerSpec struct {
	ApiManagementName     string `json:"api_management_name"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// +kubebuilder:validation:UniqueItems=true
	AuthorizationMethods []string `json:"authorization_methods"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BearerTokenSendingMethods []string `json:"bearer_token_sending_methods,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ClientAuthenticationMethod []string `json:"client_authentication_method,omitempty"`
	ClientId                   string   `json:"client_id"`
	ClientRegistrationEndpoint string   `json:"client_registration_endpoint"`
	// +optional
	ClientSecret string `json:"client_secret,omitempty"`
	// +optional
	DefaultScope string `json:"default_scope,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name"`
	// +kubebuilder:validation:UniqueItems=true
	GrantTypes        []string `json:"grant_types"`
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	// +optional
	ResourceOwnerPassword string `json:"resource_owner_password,omitempty"`
	// +optional
	ResourceOwnerUsername string `json:"resource_owner_username,omitempty"`
	// +optional
	SupportState bool `json:"support_state,omitempty"`
	// +optional
	TokenBodyParameter *[]ApiManagementAuthorizationServerSpec `json:"token_body_parameter,omitempty"`
	// +optional
	TokenEndpoint string `json:"token_endpoint,omitempty"`
}

type ApiManagementAuthorizationServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
