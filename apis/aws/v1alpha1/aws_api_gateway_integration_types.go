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

type AwsApiGatewayIntegration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayIntegrationSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayIntegrationStatus `json:"status,omitempty"`
}

type AwsApiGatewayIntegrationSpec struct {
	HttpMethod              string            `json:"http_method"`
	Credentials             string            `json:"credentials"`
	CacheKeyParameters      []string          `json:"cache_key_parameters"`
	RestApiId               string            `json:"rest_api_id"`
	Type                    string            `json:"type"`
	IntegrationHttpMethod   string            `json:"integration_http_method"`
	RequestParametersInJson string            `json:"request_parameters_in_json"`
	TimeoutMilliseconds     int               `json:"timeout_milliseconds"`
	ResourceId              string            `json:"resource_id"`
	ConnectionType          string            `json:"connection_type"`
	RequestTemplates        map[string]string `json:"request_templates"`
	ContentHandling         string            `json:"content_handling"`
	ConnectionId            string            `json:"connection_id"`
	Uri                     string            `json:"uri"`
	RequestParameters       map[string]string `json:"request_parameters"`
	PassthroughBehavior     string            `json:"passthrough_behavior"`
	CacheNamespace          string            `json:"cache_namespace"`
}



type AwsApiGatewayIntegrationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayIntegrationList is a list of AwsApiGatewayIntegrations
type AwsApiGatewayIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayIntegration CRD objects
	Items []AwsApiGatewayIntegration `json:"items,omitempty"`
}