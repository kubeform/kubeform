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

type AwsApiGatewayMethod struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayMethodSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayMethodStatus `json:"status,omitempty"`
}

type AwsApiGatewayMethodSpec struct {
	RestApiId               string            `json:"rest_api_id"`
	ResourceId              string            `json:"resource_id"`
	HttpMethod              string            `json:"http_method"`
	Authorization           string            `json:"authorization"`
	AuthorizerId            string            `json:"authorizer_id"`
	RequestModels           map[string]string `json:"request_models"`
	RequestParameters       map[string]bool   `json:"request_parameters"`
	AuthorizationScopes     []string          `json:"authorization_scopes"`
	ApiKeyRequired          bool              `json:"api_key_required"`
	RequestParametersInJson string            `json:"request_parameters_in_json"`
	RequestValidatorId      string            `json:"request_validator_id"`
}



type AwsApiGatewayMethodStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayMethodList is a list of AwsApiGatewayMethods
type AwsApiGatewayMethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayMethod CRD objects
	Items []AwsApiGatewayMethod `json:"items,omitempty"`
}