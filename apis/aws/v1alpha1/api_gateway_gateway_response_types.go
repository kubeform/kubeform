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

type ApiGatewayGatewayResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayGatewayResponseSpec   `json:"spec,omitempty"`
	Status            ApiGatewayGatewayResponseStatus `json:"status,omitempty"`
}

type ApiGatewayGatewayResponseSpec struct {
	// +optional
	ResponseParameters map[string]string `json:"response_parameters,omitempty"`
	// +optional
	ResponseTemplates map[string]string `json:"response_templates,omitempty"`
	ResponseType      string            `json:"response_type"`
	RestApiId         string            `json:"rest_api_id"`
	// +optional
	StatusCode string `json:"status_code,omitempty"`
}

type ApiGatewayGatewayResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayGatewayResponseList is a list of ApiGatewayGatewayResponses
type ApiGatewayGatewayResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayGatewayResponse CRD objects
	Items []ApiGatewayGatewayResponse `json:"items,omitempty"`
}
