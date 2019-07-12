package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayBasePathMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayBasePathMappingSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayBasePathMappingStatus `json:"status,omitempty"`
}

type AwsApiGatewayBasePathMappingSpec struct {
	ApiId      string `json:"api_id"`
	BasePath   string `json:"base_path"`
	StageName  string `json:"stage_name"`
	DomainName string `json:"domain_name"`
}



type AwsApiGatewayBasePathMappingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayBasePathMappingList is a list of AwsApiGatewayBasePathMappings
type AwsApiGatewayBasePathMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayBasePathMapping CRD objects
	Items []AwsApiGatewayBasePathMapping `json:"items,omitempty"`
}