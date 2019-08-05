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

type EndpointsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointsServiceSpec   `json:"spec,omitempty"`
	Status            EndpointsServiceStatus `json:"status,omitempty"`
}

type EndpointsServiceSpecApisMethods struct {
	Name         string `json:"name" tf:"name"`
	RequestType  string `json:"requestType" tf:"request_type"`
	ResponseType string `json:"responseType" tf:"response_type"`
	Syntax       string `json:"syntax" tf:"syntax"`
}

type EndpointsServiceSpecApis struct {
	Methods []EndpointsServiceSpecApisMethods `json:"methods" tf:"methods"`
	Name    string                            `json:"name" tf:"name"`
	Syntax  string                            `json:"syntax" tf:"syntax"`
	Version string                            `json:"version" tf:"version"`
}

type EndpointsServiceSpecEndpoints struct {
	Address string `json:"address" tf:"address"`
	Name    string `json:"name" tf:"name"`
}

type EndpointsServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Apis       []EndpointsServiceSpecApis      `json:"apis" tf:"apis"`
	ConfigID   string                          `json:"configID" tf:"config_id"`
	DnsAddress string                          `json:"dnsAddress" tf:"dns_address"`
	Endpoints  []EndpointsServiceSpecEndpoints `json:"endpoints" tf:"endpoints"`
	// +optional
	GrpcConfig string `json:"grpcConfig,omitempty" tf:"grpc_config,omitempty"`
	// +optional
	OpenapiConfig string `json:"openapiConfig,omitempty" tf:"openapi_config,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// Deprecated
	ProtocOutput string `json:"protocOutput,omitempty" tf:"protoc_output,omitempty"`
	// +optional
	ProtocOutputBase64 string `json:"protocOutputBase64,omitempty" tf:"protoc_output_base64,omitempty"`
	ServiceName        string `json:"serviceName" tf:"service_name"`
}

type EndpointsServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EndpointsServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EndpointsServiceList is a list of EndpointsServices
type EndpointsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EndpointsService CRD objects
	Items []EndpointsService `json:"items,omitempty"`
}
