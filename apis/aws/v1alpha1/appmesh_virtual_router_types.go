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

type AppmeshVirtualRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualRouterSpec   `json:"spec,omitempty"`
	Status            AppmeshVirtualRouterStatus `json:"status,omitempty"`
}

type AppmeshVirtualRouterSpecSpecListenerPortMapping struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AppmeshVirtualRouterSpecSpecListener struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	PortMapping []AppmeshVirtualRouterSpecSpecListener `json:"port_mapping"`
}

type AppmeshVirtualRouterSpecSpec struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Listener []AppmeshVirtualRouterSpecSpec `json:"listener"`
}

type AppmeshVirtualRouterSpec struct {
	MeshName string `json:"mesh_name"`
	Name     string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshVirtualRouterSpec `json:"spec"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AppmeshVirtualRouterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppmeshVirtualRouterList is a list of AppmeshVirtualRouters
type AppmeshVirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppmeshVirtualRouter CRD objects
	Items []AppmeshVirtualRouter `json:"items,omitempty"`
}
