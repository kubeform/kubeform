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

type SagemakerEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerEndpointSpec   `json:"spec,omitempty"`
	Status            SagemakerEndpointStatus `json:"status,omitempty"`
}

type SagemakerEndpointSpec struct {
	EndpointConfigName string `json:"endpoint_config_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SagemakerEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerEndpointList is a list of SagemakerEndpoints
type SagemakerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerEndpoint CRD objects
	Items []SagemakerEndpoint `json:"items,omitempty"`
}
