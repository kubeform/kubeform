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

type AwsApiGatewayIntegrationResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayIntegrationResponseSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayIntegrationResponseStatus `json:"status,omitempty"`
}

type AwsApiGatewayIntegrationResponseSpec struct {
	SelectionPattern         string            `json:"selection_pattern"`
	ResponseTemplates        map[string]string `json:"response_templates"`
	ResponseParameters       map[string]string `json:"response_parameters"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
	ContentHandling          string            `json:"content_handling"`
	HttpMethod               string            `json:"http_method"`
	StatusCode               string            `json:"status_code"`
	RestApiId                string            `json:"rest_api_id"`
	ResourceId               string            `json:"resource_id"`
}



type AwsApiGatewayIntegrationResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayIntegrationResponseList is a list of AwsApiGatewayIntegrationResponses
type AwsApiGatewayIntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayIntegrationResponse CRD objects
	Items []AwsApiGatewayIntegrationResponse `json:"items,omitempty"`
}