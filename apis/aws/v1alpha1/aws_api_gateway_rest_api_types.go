package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Description            string                     `json:"description"`
	ApiKeySource           string                     `json:"api_key_source"`
	BinaryMediaTypes       []string                   `json:"binary_media_types"`
	CreatedDate            string                     `json:"created_date"`
	EndpointConfiguration  []AwsApiGatewayRestApiSpec `json:"endpoint_configuration"`
	ExecutionArn           string                     `json:"execution_arn"`
	Name                   string                     `json:"name"`
	Policy                 string                     `json:"policy"`
	Body                   string                     `json:"body"`
	MinimumCompressionSize int                        `json:"minimum_compression_size"`
	RootResourceId         string                     `json:"root_resource_id"`
}



type AwsApiGatewayRestApiStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayRestApiList is a list of AwsApiGatewayRestApis
type AwsApiGatewayRestApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayRestApi CRD objects
	Items []AwsApiGatewayRestApi `json:"items,omitempty"`
}