package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ResponseParameters map[string]string `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`
	// +optional
	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`
	ResponseType      string            `json:"responseType" tf:"response_type"`
	RestAPIID         string            `json:"restAPIID" tf:"rest_api_id"`
	// +optional
	StatusCode string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ApiGatewayGatewayResponseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayGatewayResponseSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
