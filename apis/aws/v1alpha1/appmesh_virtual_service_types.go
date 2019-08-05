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

type AppmeshVirtualService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualServiceSpec   `json:"spec,omitempty"`
	Status            AppmeshVirtualServiceStatus `json:"status,omitempty"`
}

type AppmeshVirtualServiceSpecSpecProviderVirtualNode struct {
	VirtualNodeName string `json:"virtualNodeName" tf:"virtual_node_name"`
}

type AppmeshVirtualServiceSpecSpecProviderVirtualRouter struct {
	VirtualRouterName string `json:"virtualRouterName" tf:"virtual_router_name"`
}

type AppmeshVirtualServiceSpecSpecProvider struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualNode []AppmeshVirtualServiceSpecSpecProviderVirtualNode `json:"virtualNode,omitempty" tf:"virtual_node,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualRouter []AppmeshVirtualServiceSpecSpecProviderVirtualRouter `json:"virtualRouter,omitempty" tf:"virtual_router,omitempty"`
}

type AppmeshVirtualServiceSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Provider []AppmeshVirtualServiceSpecSpecProvider `json:"provider,omitempty" tf:"provider,omitempty"`
}

type AppmeshVirtualServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	MeshName string `json:"meshName" tf:"mesh_name"`
	Name     string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshVirtualServiceSpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppmeshVirtualServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppmeshVirtualServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppmeshVirtualServiceList is a list of AppmeshVirtualServices
type AppmeshVirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppmeshVirtualService CRD objects
	Items []AppmeshVirtualService `json:"items,omitempty"`
}
