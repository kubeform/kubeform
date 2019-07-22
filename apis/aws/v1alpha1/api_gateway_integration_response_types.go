package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	ContentHandling string `json:"contentHandling,omitempty" tf:"content_handling,omitempty"`
	HttpMethod      string `json:"httpMethod" tf:"http_method"`
	ResourceID      string `json:"resourceID" tf:"resource_id"`
	// +optional
	ResponseParameters map[string]string `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`
	// +optional
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`
	RestAPIID         string            `json:"restAPIID" tf:"rest_api_id"`
	// +optional
	SelectionPattern string `json:"selectionPattern,omitempty" tf:"selection_pattern,omitempty"`
	StatusCode       string `json:"statusCode" tf:"status_code"`
}

type ApiGatewayIntegrationResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
