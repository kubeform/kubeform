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

type ApiGatewayRestAPI struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayRestAPISpec   `json:"spec,omitempty"`
	Status            ApiGatewayRestAPIStatus `json:"status,omitempty"`
}

type ApiGatewayRestAPISpecEndpointConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Types []string `json:"types" tf:"types"`
}

type ApiGatewayRestAPISpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApiKeySource string `json:"apiKeySource,omitempty" tf:"api_key_source,omitempty"`
	// +optional
	BinaryMediaTypes []string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types,omitempty"`
	// +optional
	Body string `json:"body,omitempty" tf:"body,omitempty"`
	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	EndpointConfiguration []ApiGatewayRestAPISpecEndpointConfiguration `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`
	// +optional
	ExecutionArn string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`
	// +optional
	MinimumCompressionSize int    `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	RootResourceID string `json:"rootResourceID,omitempty" tf:"root_resource_id,omitempty"`
}

type ApiGatewayRestAPIStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayRestAPISpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayRestAPIList is a list of ApiGatewayRestAPIs
type ApiGatewayRestAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayRestAPI CRD objects
	Items []ApiGatewayRestAPI `json:"items,omitempty"`
}
