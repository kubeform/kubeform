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

type AwsCognitoIdentityProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoIdentityProviderSpec   `json:"spec,omitempty"`
	Status            AwsCognitoIdentityProviderStatus `json:"status,omitempty"`
}

type AwsCognitoIdentityProviderSpec struct {
	ProviderDetails  map[string]string `json:"provider_details"`
	ProviderName     string            `json:"provider_name"`
	ProviderType     string            `json:"provider_type"`
	UserPoolId       string            `json:"user_pool_id"`
	AttributeMapping map[string]string `json:"attribute_mapping"`
	IdpIdentifiers   []string          `json:"idp_identifiers"`
}

type AwsCognitoIdentityProviderStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCognitoIdentityProviderList is a list of AwsCognitoIdentityProviders
type AwsCognitoIdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoIdentityProvider CRD objects
	Items []AwsCognitoIdentityProvider `json:"items,omitempty"`
}
