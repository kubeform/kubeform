package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	AuthorizerCredentials string `json:"authorizerCredentials,omitempty" tf:"authorizer_credentials,omitempty"`
	// +optional
	AuthorizerResultTtlInSeconds int `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`
	// +optional
	AuthorizerURI string `json:"authorizerURI,omitempty" tf:"authorizer_uri,omitempty"`
	// +optional
	IdentitySource string `json:"identitySource,omitempty" tf:"identity_source,omitempty"`
	// +optional
	IdentityValidationExpression string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
	Name                         string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ProviderArns []string `json:"providerArns,omitempty" tf:"provider_arns,omitempty"`
	RestAPIID    string   `json:"restAPIID" tf:"rest_api_id"`
	// +optional
	Type        string                    `json:"type,omitempty" tf:"type,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiGatewayAuthorizerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
