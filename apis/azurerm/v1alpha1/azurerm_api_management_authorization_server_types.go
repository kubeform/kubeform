package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	ResourceGroupName          string                                        `json:"resource_group_name"`
	AuthorizationMethods       []string                                      `json:"authorization_methods"`
	DisplayName                string                                        `json:"display_name"`
	ResourceOwnerUsername      string                                        `json:"resource_owner_username"`
	GrantTypes                 []string                                      `json:"grant_types"`
	Description                string                                        `json:"description"`
	SupportState               bool                                          `json:"support_state"`
	Name                       string                                        `json:"name"`
	ApiManagementName          string                                        `json:"api_management_name"`
	AuthorizationEndpoint      string                                        `json:"authorization_endpoint"`
	ClientRegistrationEndpoint string                                        `json:"client_registration_endpoint"`
	TokenEndpoint              string                                        `json:"token_endpoint"`
	ClientId                   string                                        `json:"client_id"`
	BearerTokenSendingMethods  []string                                      `json:"bearer_token_sending_methods"`
	ClientSecret               string                                        `json:"client_secret"`
	ResourceOwnerPassword      string                                        `json:"resource_owner_password"`
	ClientAuthenticationMethod []string                                      `json:"client_authentication_method"`
	DefaultScope               string                                        `json:"default_scope"`
	TokenBodyParameter         []AzurermApiManagementAuthorizationServerSpec `json:"token_body_parameter"`
}



type AzurermApiManagementAuthorizationServerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementAuthorizationServerList is a list of AzurermApiManagementAuthorizationServers
type AzurermApiManagementAuthorizationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementAuthorizationServer CRD objects
	Items []AzurermApiManagementAuthorizationServer `json:"items,omitempty"`
}