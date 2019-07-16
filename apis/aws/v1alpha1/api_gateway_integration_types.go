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

type ApiGatewayIntegration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayIntegrationSpec   `json:"spec,omitempty"`
	Status            ApiGatewayIntegrationStatus `json:"status,omitempty"`
}

type ApiGatewayIntegrationSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CacheKeyParameters []string `json:"cache_key_parameters,omitempty"`
	// +optional
	ConnectionId string `json:"connection_id,omitempty"`
	// +optional
	ConnectionType string `json:"connection_type,omitempty"`
	// +optional
	ContentHandling string `json:"content_handling,omitempty"`
	// +optional
	Credentials string `json:"credentials,omitempty"`
	HttpMethod  string `json:"http_method"`
	// +optional
	IntegrationHttpMethod string `json:"integration_http_method,omitempty"`
	// +optional
	RequestParameters map[string]string `json:"request_parameters,omitempty"`
	// +optional
	RequestTemplates map[string]string `json:"request_templates,omitempty"`
	ResourceId       string            `json:"resource_id"`
	RestApiId        string            `json:"rest_api_id"`
	// +optional
	TimeoutMilliseconds int    `json:"timeout_milliseconds,omitempty"`
	Type                string `json:"type"`
	// +optional
	Uri string `json:"uri,omitempty"`
}

type ApiGatewayIntegrationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayIntegrationList is a list of ApiGatewayIntegrations
type ApiGatewayIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayIntegration CRD objects
	Items []ApiGatewayIntegration `json:"items,omitempty"`
}
