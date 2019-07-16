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

type VpcEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSpec   `json:"spec,omitempty"`
	Status            VpcEndpointStatus `json:"status,omitempty"`
}

type VpcEndpointSpec struct {
	// +optional
	AutoAccept bool `json:"auto_accept,omitempty"`
	// +optional
	PrivateDnsEnabled bool   `json:"private_dns_enabled,omitempty"`
	ServiceName       string `json:"service_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	VpcEndpointType string `json:"vpc_endpoint_type,omitempty"`
	VpcId           string `json:"vpc_id"`
}

type VpcEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointList is a list of VpcEndpoints
type VpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpoint CRD objects
	Items []VpcEndpoint `json:"items,omitempty"`
}
