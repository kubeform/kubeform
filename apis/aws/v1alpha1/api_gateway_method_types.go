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

type ApiGatewayMethod struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodSpec   `json:"spec,omitempty"`
	Status            ApiGatewayMethodStatus `json:"status,omitempty"`
}

type ApiGatewayMethodSpec struct {
	// +optional
	ApiKeyRequired bool   `json:"api_key_required,omitempty"`
	Authorization  string `json:"authorization"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AuthorizationScopes []string `json:"authorization_scopes,omitempty"`
	// +optional
	AuthorizerId string `json:"authorizer_id,omitempty"`
	HttpMethod   string `json:"http_method"`
	// +optional
	RequestModels map[string]string `json:"request_models,omitempty"`
	// +optional
	RequestParameters map[string]bool `json:"request_parameters,omitempty"`
	// +optional
	RequestValidatorId string `json:"request_validator_id,omitempty"`
	ResourceId         string `json:"resource_id"`
	RestApiId          string `json:"rest_api_id"`
}

type ApiGatewayMethodStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
