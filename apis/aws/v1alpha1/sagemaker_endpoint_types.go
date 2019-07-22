package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
	EndpointConfigName string                    `json:"endpointConfigName" tf:"endpoint_config_name"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SagemakerEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
