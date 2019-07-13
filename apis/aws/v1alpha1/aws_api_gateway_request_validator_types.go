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

type AwsApiGatewayRequestValidator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayRequestValidatorSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayRequestValidatorStatus `json:"status,omitempty"`
}

type AwsApiGatewayRequestValidatorSpec struct {
	Name                      string `json:"name"`
	ValidateRequestBody       bool   `json:"validate_request_body"`
	ValidateRequestParameters bool   `json:"validate_request_parameters"`
	RestApiId                 string `json:"rest_api_id"`
}



type AwsApiGatewayRequestValidatorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayRequestValidatorList is a list of AwsApiGatewayRequestValidators
type AwsApiGatewayRequestValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayRequestValidator CRD objects
	Items []AwsApiGatewayRequestValidator `json:"items,omitempty"`
}