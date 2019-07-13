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

type AwsApiGatewayRestApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayRestApiSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayRestApiStatus `json:"status,omitempty"`
}

type AwsApiGatewayRestApiSpecEndpointConfiguration struct {
	Types []string `json:"types"`
}

type AwsApiGatewayRestApiSpec struct {
	BinaryMediaTypes       []string                   `json:"binary_media_types"`
	MinimumCompressionSize int                        `json:"minimum_compression_size"`
	RootResourceId         string                     `json:"root_resource_id"`
	ExecutionArn           string                     `json:"execution_arn"`
	Name                   string                     `json:"name"`
	Description            string                     `json:"description"`
	ApiKeySource           string                     `json:"api_key_source"`
	EndpointConfiguration  []AwsApiGatewayRestApiSpec `json:"endpoint_configuration"`
	Policy                 string                     `json:"policy"`
	Body                   string                     `json:"body"`
	CreatedDate            string                     `json:"created_date"`
}



type AwsApiGatewayRestApiStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayRestApiList is a list of AwsApiGatewayRestApis
type AwsApiGatewayRestApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayRestApi CRD objects
	Items []AwsApiGatewayRestApi `json:"items,omitempty"`
}