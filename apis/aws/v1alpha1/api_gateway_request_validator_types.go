package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiGatewayRequestValidator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayRequestValidatorSpec   `json:"spec,omitempty"`
	Status            ApiGatewayRequestValidatorStatus `json:"status,omitempty"`
}

type ApiGatewayRequestValidatorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name      string `json:"name" tf:"name"`
	RestAPIID string `json:"restAPIID" tf:"rest_api_id"`
	// +optional
	ValidateRequestBody bool `json:"validateRequestBody,omitempty" tf:"validate_request_body,omitempty"`
	// +optional
	ValidateRequestParameters bool `json:"validateRequestParameters,omitempty" tf:"validate_request_parameters,omitempty"`
}

type ApiGatewayRequestValidatorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayRequestValidatorSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayRequestValidatorList is a list of ApiGatewayRequestValidators
type ApiGatewayRequestValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayRequestValidator CRD objects
	Items []ApiGatewayRequestValidator `json:"items,omitempty"`
}
