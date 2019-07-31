package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type TransferServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferServerSpec   `json:"spec,omitempty"`
	Status            TransferServerStatus `json:"status,omitempty"`
}

type TransferServerSpecEndpointDetails struct {
	VpcEndpointID string `json:"vpcEndpointID" tf:"vpc_endpoint_id"`
}

type TransferServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	EndpointDetails []TransferServerSpecEndpointDetails `json:"endpointDetails,omitempty" tf:"endpoint_details,omitempty"`
	// +optional
	EndpointType string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// +optional
	IdentityProviderType string `json:"identityProviderType,omitempty" tf:"identity_provider_type,omitempty"`
	// +optional
	InvocationRole string `json:"invocationRole,omitempty" tf:"invocation_role,omitempty"`
	// +optional
	LoggingRole string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
}

type TransferServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TransferServerList is a list of TransferServers
type TransferServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TransferServer CRD objects
	Items []TransferServer `json:"items,omitempty"`
}
