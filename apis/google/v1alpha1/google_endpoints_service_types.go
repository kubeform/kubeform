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

type GoogleEndpointsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleEndpointsServiceSpec   `json:"spec,omitempty"`
	Status            GoogleEndpointsServiceStatus `json:"status,omitempty"`
}

type GoogleEndpointsServiceSpecApisMethods struct {
	Syntax       string `json:"syntax"`
	RequestType  string `json:"request_type"`
	ResponseType string `json:"response_type"`
	Name         string `json:"name"`
}

type GoogleEndpointsServiceSpecApis struct {
	Name    string                           `json:"name"`
	Syntax  string                           `json:"syntax"`
	Version string                           `json:"version"`
	Methods []GoogleEndpointsServiceSpecApis `json:"methods"`
}

type GoogleEndpointsServiceSpecEndpoints struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type GoogleEndpointsServiceSpec struct {
	ServiceName        string                       `json:"service_name"`
	GrpcConfig         string                       `json:"grpc_config"`
	ProtocOutputBase64 string                       `json:"protoc_output_base64"`
	Project            string                       `json:"project"`
	Apis               []GoogleEndpointsServiceSpec `json:"apis"`
	Endpoints          []GoogleEndpointsServiceSpec `json:"endpoints"`
	OpenapiConfig      string                       `json:"openapi_config"`
	ProtocOutput       string                       `json:"protoc_output"`
	ConfigId           string                       `json:"config_id"`
	DnsAddress         string                       `json:"dns_address"`
}

type GoogleEndpointsServiceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleEndpointsServiceList is a list of GoogleEndpointsServices
type GoogleEndpointsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleEndpointsService CRD objects
	Items []GoogleEndpointsService `json:"items,omitempty"`
}
