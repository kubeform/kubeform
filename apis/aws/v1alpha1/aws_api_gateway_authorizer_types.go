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

type AwsApiGatewayAuthorizer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayAuthorizerSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayAuthorizerStatus `json:"status,omitempty"`
}

type AwsApiGatewayAuthorizerSpec struct {
	AuthorizerResultTtlInSeconds int      `json:"authorizer_result_ttl_in_seconds"`
	ProviderArns                 []string `json:"provider_arns"`
	RestApiId                    string   `json:"rest_api_id"`
	Type                         string   `json:"type"`
	AuthorizerCredentials        string   `json:"authorizer_credentials"`
	IdentityValidationExpression string   `json:"identity_validation_expression"`
	AuthorizerUri                string   `json:"authorizer_uri"`
	IdentitySource               string   `json:"identity_source"`
	Name                         string   `json:"name"`
}

type AwsApiGatewayAuthorizerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayAuthorizerList is a list of AwsApiGatewayAuthorizers
type AwsApiGatewayAuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayAuthorizer CRD objects
	Items []AwsApiGatewayAuthorizer `json:"items,omitempty"`
}
