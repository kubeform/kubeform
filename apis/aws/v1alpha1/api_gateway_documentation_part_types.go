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

type ApiGatewayDocumentationPart struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDocumentationPartSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDocumentationPartStatus `json:"status,omitempty"`
}

type ApiGatewayDocumentationPartSpecLocation struct {
	// +optional
	Method string `json:"method,omitempty" tf:"method,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	StatusCode string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
	Type       string `json:"type" tf:"type"`
}

type ApiGatewayDocumentationPartSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:MaxItems=1
	Location   []ApiGatewayDocumentationPartSpecLocation `json:"location" tf:"location"`
	Properties string                                    `json:"properties" tf:"properties"`
	RestAPIID  string                                    `json:"restAPIID" tf:"rest_api_id"`
}

type ApiGatewayDocumentationPartStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayDocumentationPartSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayDocumentationPartList is a list of ApiGatewayDocumentationParts
type ApiGatewayDocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayDocumentationPart CRD objects
	Items []ApiGatewayDocumentationPart `json:"items,omitempty"`
}
