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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	LastUpdatedDate string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`
	MeshName        string `json:"meshName" tf:"mesh_name"`
	Name            string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshVirtualServiceSpecSpec `json:"spec" tf:"spec"`
}

type AppmeshVirtualServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppmeshVirtualServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
