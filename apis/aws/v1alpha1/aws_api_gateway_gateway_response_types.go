package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayGatewayResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayGatewayResponseSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayGatewayResponseStatus `json:"status,omitempty"`
}

type AwsApiGatewayGatewayResponseSpec struct {
	RestApiId          string            `json:"rest_api_id"`
	ResponseType       string            `json:"response_type"`
	StatusCode         string            `json:"status_code"`
	ResponseTemplates  map[string]string `json:"response_templates"`
	ResponseParameters map[string]string `json:"response_parameters"`
}

type AwsApiGatewayGatewayResponseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayGatewayResponseList is a list of AwsApiGatewayGatewayResponses
type AwsApiGatewayGatewayResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayGatewayResponse CRD objects
	Items []AwsApiGatewayGatewayResponse `json:"items,omitempty"`
}
