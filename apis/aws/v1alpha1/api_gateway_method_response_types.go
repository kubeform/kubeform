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

type ApiGatewayMethodResponse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodResponseSpec   `json:"spec,omitempty"`
	Status            ApiGatewayMethodResponseStatus `json:"status,omitempty"`
}

type ApiGatewayMethodResponseSpec struct {
	HttpMethod string `json:"httpMethod" tf:"http_method"`
	ResourceID string `json:"resourceID" tf:"resource_id"`
	// +optional
	ResponseModels map[string]string `json:"responseModels,omitempty" tf:"response_models,omitempty"`
	// +optional
	ResponseParameters map[string]bool           `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`
	RestAPIID          string                    `json:"restAPIID" tf:"rest_api_id"`
	StatusCode         string                    `json:"statusCode" tf:"status_code"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiGatewayMethodResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
