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

type ApiGatewayBasePathMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayBasePathMappingSpec   `json:"spec,omitempty"`
	Status            ApiGatewayBasePathMappingStatus `json:"status,omitempty"`
}

type ApiGatewayBasePathMappingSpec struct {
	ApiID string `json:"apiID" tf:"api_id"`
	// +optional
	BasePath   string `json:"basePath,omitempty" tf:"base_path,omitempty"`
	DomainName string `json:"domainName" tf:"domain_name"`
	// +optional
	StageName   string                    `json:"stageName,omitempty" tf:"stage_name,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiGatewayBasePathMappingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayBasePathMappingList is a list of ApiGatewayBasePathMappings
type ApiGatewayBasePathMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayBasePathMapping CRD objects
	Items []ApiGatewayBasePathMapping `json:"items,omitempty"`
}
