package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayIntegrationSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayIntegrationStatus `json:"status,omitempty"`
}

type AwsApiGatewayIntegrationSpec struct {
	ConnectionType          string            `json:"connection_type"`
	CacheNamespace          string            `json:"cache_namespace"`
	TimeoutMilliseconds     int               `json:"timeout_milliseconds"`
	ResourceId              string            `json:"resource_id"`
	HttpMethod              string            `json:"http_method"`
	PassthroughBehavior     string            `json:"passthrough_behavior"`
	ContentHandling         string            `json:"content_handling"`
	CacheKeyParameters      []string          `json:"cache_key_parameters"`
	Type                    string            `json:"type"`
	ConnectionId            string            `json:"connection_id"`
	Uri                     string            `json:"uri"`
	RequestParameters       map[string]string `json:"request_parameters"`
	RequestParametersInJson string            `json:"request_parameters_in_json"`
	RestApiId               string            `json:"rest_api_id"`
	Credentials             string            `json:"credentials"`
	IntegrationHttpMethod   string            `json:"integration_http_method"`
	RequestTemplates        map[string]string `json:"request_templates"`
}

type AwsApiGatewayIntegrationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationList is a list of AwsApiGatewayIntegrations
type AwsApiGatewayIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayIntegration CRD objects
	Items []AwsApiGatewayIntegration `json:"items,omitempty"`
}
