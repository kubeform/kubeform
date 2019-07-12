package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRequestValidator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayRequestValidatorSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayRequestValidatorStatus `json:"status,omitempty"`
}

type AwsApiGatewayRequestValidatorSpec struct {
	RestApiId                 string `json:"rest_api_id"`
	Name                      string `json:"name"`
	ValidateRequestBody       bool   `json:"validate_request_body"`
	ValidateRequestParameters bool   `json:"validate_request_parameters"`
}

type AwsApiGatewayRequestValidatorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayRequestValidatorList is a list of AwsApiGatewayRequestValidators
type AwsApiGatewayRequestValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayRequestValidator CRD objects
	Items []AwsApiGatewayRequestValidator `json:"items,omitempty"`
}
