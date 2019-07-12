package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayMethodResponseSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayMethodResponseStatus `json:"status,omitempty"`
}

type AwsApiGatewayMethodResponseSpec struct {
	RestApiId                string            `json:"rest_api_id"`
	ResourceId               string            `json:"resource_id"`
	HttpMethod               string            `json:"http_method"`
	StatusCode               string            `json:"status_code"`
	ResponseModels           map[string]string `json:"response_models"`
	ResponseParameters       map[string]bool   `json:"response_parameters"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
}



type AwsApiGatewayMethodResponseStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodResponseList is a list of AwsApiGatewayMethodResponses
type AwsApiGatewayMethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayMethodResponse CRD objects
	Items []AwsApiGatewayMethodResponse `json:"items,omitempty"`
}