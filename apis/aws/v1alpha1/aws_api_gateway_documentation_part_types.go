package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationPart struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayDocumentationPartSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayDocumentationPartStatus `json:"status,omitempty"`
}

type AwsApiGatewayDocumentationPartSpecLocation struct {
	Method     string `json:"method"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	StatusCode string `json:"status_code"`
	Type       string `json:"type"`
}

type AwsApiGatewayDocumentationPartSpec struct {
	Location   []AwsApiGatewayDocumentationPartSpec `json:"location"`
	Properties string                               `json:"properties"`
	RestApiId  string                               `json:"rest_api_id"`
}



type AwsApiGatewayDocumentationPartStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationPartList is a list of AwsApiGatewayDocumentationParts
type AwsApiGatewayDocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayDocumentationPart CRD objects
	Items []AwsApiGatewayDocumentationPart `json:"items,omitempty"`
}