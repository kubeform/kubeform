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

type ApiGatewayResource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayResourceSpec   `json:"spec,omitempty"`
	Status            ApiGatewayResourceStatus `json:"status,omitempty"`
}

type ApiGatewayResourceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	ParentID    string                    `json:"parentID" tf:"parent_id"`
	PathPart    string                    `json:"pathPart" tf:"path_part"`
	RestAPIID   string                    `json:"restAPIID" tf:"rest_api_id"`
}

type ApiGatewayResourceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayResourceList is a list of ApiGatewayResources
type ApiGatewayResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayResource CRD objects
	Items []ApiGatewayResource `json:"items,omitempty"`
}
