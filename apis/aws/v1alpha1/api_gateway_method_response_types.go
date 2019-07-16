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

type ApiGatewayMethodResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodResponseSpec   `json:"spec,omitempty"`
	Status            ApiGatewayMethodResponseStatus `json:"status,omitempty"`
}

type ApiGatewayMethodResponseSpec struct {
	HttpMethod string `json:"http_method"`
	ResourceId string `json:"resource_id"`
	// +optional
	ResponseModels map[string]string `json:"response_models,omitempty"`
	// +optional
	ResponseParameters map[string]bool `json:"response_parameters,omitempty"`
	RestApiId          string          `json:"rest_api_id"`
	StatusCode         string          `json:"status_code"`
}

type ApiGatewayMethodResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayMethodResponseList is a list of ApiGatewayMethodResponses
type ApiGatewayMethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayMethodResponse CRD objects
	Items []ApiGatewayMethodResponse `json:"items,omitempty"`
}
