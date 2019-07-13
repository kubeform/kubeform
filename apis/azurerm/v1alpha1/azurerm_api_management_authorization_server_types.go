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

type AzurermApiManagementAuthorizationServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementAuthorizationServerSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementAuthorizationServerStatus `json:"status,omitempty"`
}

type AzurermApiManagementAuthorizationServerSpecTokenBodyParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AzurermApiManagementAuthorizationServerSpec struct {
	DisplayName                string                                        `json:"display_name"`
	ClientSecret               string                                        `json:"client_secret"`
	ResourceOwnerUsername      string                                        `json:"resource_owner_username"`
	ResourceOwnerPassword      string                                        `json:"resource_owner_password"`
	ApiManagementName          string                                        `json:"api_management_name"`
	ResourceGroupName          string                                        `json:"resource_group_name"`
	AuthorizationMethods       []string                                      `json:"authorization_methods"`
	ClientId                   string                                        `json:"client_id"`
	ClientRegistrationEndpoint string                                        `json:"client_registration_endpoint"`
	Description                string                                        `json:"description"`
	SupportState               bool                                          `json:"support_state"`
	TokenBodyParameter         []AzurermApiManagementAuthorizationServerSpec `json:"token_body_parameter"`
	AuthorizationEndpoint      string                                        `json:"authorization_endpoint"`
	ClientAuthenticationMethod []string                                      `json:"client_authentication_method"`
	DefaultScope               string                                        `json:"default_scope"`
	TokenEndpoint              string                                        `json:"token_endpoint"`
	Name                       string                                        `json:"name"`
	GrantTypes                 []string                                      `json:"grant_types"`
	BearerTokenSendingMethods  []string                                      `json:"bearer_token_sending_methods"`
}



type AzurermApiManagementAuthorizationServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementAuthorizationServerList is a list of AzurermApiManagementAuthorizationServers
type AzurermApiManagementAuthorizationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementAuthorizationServer CRD objects
	Items []AzurermApiManagementAuthorizationServer `json:"items,omitempty"`
}