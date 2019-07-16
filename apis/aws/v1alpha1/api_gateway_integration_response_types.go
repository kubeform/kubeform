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

type ApiGatewayIntegrationResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayIntegrationResponseSpec   `json:"spec,omitempty"`
	Status            ApiGatewayIntegrationResponseStatus `json:"status,omitempty"`
}

type ApiGatewayIntegrationResponseSpec struct {
	// +optional
	ContentHandling string `json:"content_handling,omitempty"`
	HttpMethod      string `json:"http_method"`
	ResourceId      string `json:"resource_id"`
	// +optional
	ResponseParameters map[string]string `json:"response_parameters,omitempty"`
	// +optional
	ResponseTemplates map[string]string `json:"response_templates,omitempty"`
	RestApiId         string            `json:"rest_api_id"`
	// +optional
	SelectionPattern string `json:"selection_pattern,omitempty"`
	StatusCode       string `json:"status_code"`
}

type ApiGatewayIntegrationResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayIntegrationResponseList is a list of ApiGatewayIntegrationResponses
type ApiGatewayIntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayIntegrationResponse CRD objects
	Items []ApiGatewayIntegrationResponse `json:"items,omitempty"`
}
