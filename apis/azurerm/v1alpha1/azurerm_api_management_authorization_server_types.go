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
	Name                       string                                        `json:"name"`
	ResourceGroupName          string                                        `json:"resource_group_name"`
	AuthorizationMethods       []string                                      `json:"authorization_methods"`
	ClientId                   string                                        `json:"client_id"`
	SupportState               bool                                          `json:"support_state"`
	GrantTypes                 []string                                      `json:"grant_types"`
	BearerTokenSendingMethods  []string                                      `json:"bearer_token_sending_methods"`
	ClientAuthenticationMethod []string                                      `json:"client_authentication_method"`
	Description                string                                        `json:"description"`
	TokenBodyParameter         []AzurermApiManagementAuthorizationServerSpec `json:"token_body_parameter"`
	TokenEndpoint              string                                        `json:"token_endpoint"`
	AuthorizationEndpoint      string                                        `json:"authorization_endpoint"`
	ClientRegistrationEndpoint string                                        `json:"client_registration_endpoint"`
	DefaultScope               string                                        `json:"default_scope"`
	ResourceOwnerUsername      string                                        `json:"resource_owner_username"`
	ApiManagementName          string                                        `json:"api_management_name"`
	DisplayName                string                                        `json:"display_name"`
	ClientSecret               string                                        `json:"client_secret"`
	ResourceOwnerPassword      string                                        `json:"resource_owner_password"`
}

type AzurermApiManagementAuthorizationServerStatus struct {
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
