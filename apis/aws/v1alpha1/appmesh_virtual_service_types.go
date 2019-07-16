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

type AppmeshVirtualService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualServiceSpec   `json:"spec,omitempty"`
	Status            AppmeshVirtualServiceStatus `json:"status,omitempty"`
}

type AppmeshVirtualServiceSpecSpecProviderVirtualNode struct {
	VirtualNodeName string `json:"virtual_node_name"`
}

type AppmeshVirtualServiceSpecSpecProviderVirtualRouter struct {
	VirtualRouterName string `json:"virtual_router_name"`
}

type AppmeshVirtualServiceSpecSpecProvider struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualNode *[]AppmeshVirtualServiceSpecSpecProvider `json:"virtual_node,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualRouter *[]AppmeshVirtualServiceSpecSpecProvider `json:"virtual_router,omitempty"`
}

type AppmeshVirtualServiceSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Provider *[]AppmeshVirtualServiceSpecSpec `json:"provider,omitempty"`
}

type AppmeshVirtualServiceSpec struct {
	MeshName string `json:"mesh_name"`
	Name     string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshVirtualServiceSpec `json:"spec"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AppmeshVirtualServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
