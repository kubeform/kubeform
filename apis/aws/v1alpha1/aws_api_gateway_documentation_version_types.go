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

type AwsApiGatewayDocumentationVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayDocumentationVersionSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayDocumentationVersionStatus `json:"status,omitempty"`
}

type AwsApiGatewayDocumentationVersionSpec struct {
	Version     string `json:"version"`
	RestApiId   string `json:"rest_api_id"`
	Description string `json:"description"`
}



type AwsApiGatewayDocumentationVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayDocumentationVersionList is a list of AwsApiGatewayDocumentationVersions
type AwsApiGatewayDocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayDocumentationVersion CRD objects
	Items []AwsApiGatewayDocumentationVersion `json:"items,omitempty"`
}