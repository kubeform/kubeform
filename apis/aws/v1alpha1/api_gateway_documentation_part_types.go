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

type ApiGatewayDocumentationPart struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDocumentationPartSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDocumentationPartStatus `json:"status,omitempty"`
}

type ApiGatewayDocumentationPartSpecLocation struct {
	// +optional
	Method string `json:"method,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	StatusCode string `json:"status_code,omitempty"`
	Type       string `json:"type"`
}

type ApiGatewayDocumentationPartSpec struct {
	// +kubebuilder:validation:MaxItems=1
	Location   []ApiGatewayDocumentationPartSpec `json:"location"`
	Properties string                            `json:"properties"`
	RestApiId  string                            `json:"rest_api_id"`
}

type ApiGatewayDocumentationPartStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
