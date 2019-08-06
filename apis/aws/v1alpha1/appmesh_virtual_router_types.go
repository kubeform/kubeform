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

type AppmeshVirtualRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualRouterSpec   `json:"spec,omitempty"`
	Status            AppmeshVirtualRouterStatus `json:"status,omitempty"`
}

type AppmeshVirtualRouterSpecSpecListenerPortMapping struct {
	Port     int    `json:"port" tf:"port"`
	Protocol string `json:"protocol" tf:"protocol"`
}

type AppmeshVirtualRouterSpecSpecListener struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	PortMapping []AppmeshVirtualRouterSpecSpecListenerPortMapping `json:"portMapping" tf:"port_mapping"`
}

type AppmeshVirtualRouterSpecSpec struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Listener []AppmeshVirtualRouterSpecSpecListener `json:"listener" tf:"listener"`
}

type AppmeshVirtualRouterSpec struct {
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
	Spec []AppmeshVirtualRouterSpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppmeshVirtualRouterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppmeshVirtualRouterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
