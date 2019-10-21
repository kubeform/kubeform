package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type EndpointsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointsServiceSpec   `json:"spec,omitempty"`
	Status            EndpointsServiceStatus `json:"status,omitempty"`
}

type EndpointsServiceSpecApisMethods struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	RequestType string `json:"requestType,omitempty" tf:"request_type,omitempty"`
	// +optional
	ResponseType string `json:"responseType,omitempty" tf:"response_type,omitempty"`
	// +optional
	Syntax string `json:"syntax,omitempty" tf:"syntax,omitempty"`
}

type EndpointsServiceSpecApis struct {
	// +optional
	Methods []EndpointsServiceSpecApisMethods `json:"methods,omitempty" tf:"methods,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Syntax string `json:"syntax,omitempty" tf:"syntax,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type EndpointsServiceSpecEndpoints struct {
	// +optional
	Address string `json:"address,omitempty" tf:"address,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type EndpointsServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Apis []EndpointsServiceSpecApis `json:"apis,omitempty" tf:"apis,omitempty"`
	// +optional
	ConfigID string `json:"configID,omitempty" tf:"config_id,omitempty"`
	// +optional
	DnsAddress string `json:"dnsAddress,omitempty" tf:"dns_address,omitempty"`
	// +optional
	Endpoints []EndpointsServiceSpecEndpoints `json:"endpoints,omitempty" tf:"endpoints,omitempty"`
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
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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