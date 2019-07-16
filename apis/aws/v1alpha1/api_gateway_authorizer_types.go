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

type ApiGatewayAuthorizer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayAuthorizerSpec   `json:"spec,omitempty"`
	Status            ApiGatewayAuthorizerStatus `json:"status,omitempty"`
}

type ApiGatewayAuthorizerSpec struct {
	// +optional
	AuthorizerCredentials string `json:"authorizer_credentials,omitempty"`
	// +optional
	AuthorizerResultTtlInSeconds int `json:"authorizer_result_ttl_in_seconds,omitempty"`
	// +optional
	AuthorizerUri string `json:"authorizer_uri,omitempty"`
	// +optional
	IdentitySource string `json:"identity_source,omitempty"`
	// +optional
	IdentityValidationExpression string `json:"identity_validation_expression,omitempty"`
	Name                         string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ProviderArns []string `json:"provider_arns,omitempty"`
	RestApiId    string   `json:"rest_api_id"`
	// +optional
	Type string `json:"type,omitempty"`
}

type ApiGatewayAuthorizerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayAuthorizerList is a list of ApiGatewayAuthorizers
type ApiGatewayAuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayAuthorizer CRD objects
	Items []ApiGatewayAuthorizer `json:"items,omitempty"`
}
