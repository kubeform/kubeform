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
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ApiGatewayMethod struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodSpec   `json:"spec,omitempty"`
	Status            ApiGatewayMethodStatus `json:"status,omitempty"`
}

type ApiGatewayMethodSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApiKeyRequired bool   `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`
	Authorization  string `json:"authorization" tf:"authorization"`
	// +optional
	AuthorizationScopes []string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`
	// +optional
	AuthorizerID string `json:"authorizerID,omitempty" tf:"authorizer_id,omitempty"`
	HttpMethod   string `json:"httpMethod" tf:"http_method"`
	// +optional
	RequestModels map[string]string `json:"requestModels,omitempty" tf:"request_models,omitempty"`
	// +optional
	RequestParameters map[string]bool `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`
	// +optional
	RequestValidatorID string `json:"requestValidatorID,omitempty" tf:"request_validator_id,omitempty"`
	ResourceID         string `json:"resourceID" tf:"resource_id"`
	RestAPIID          string `json:"restAPIID" tf:"rest_api_id"`
}

type ApiGatewayMethodStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayMethodSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayMethodList is a list of ApiGatewayMethods
type ApiGatewayMethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayMethod CRD objects
	Items []ApiGatewayMethod `json:"items,omitempty"`
}
