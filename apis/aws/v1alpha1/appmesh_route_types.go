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

type AppmeshRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshRouteSpec   `json:"spec,omitempty"`
	Status            AppmeshRouteStatus `json:"status,omitempty"`
}

type AppmeshRouteSpecSpecHttpRouteActionWeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int    `json:"weight"`
}

type AppmeshRouteSpecSpecHttpRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	WeightedTarget []AppmeshRouteSpecSpecHttpRouteAction `json:"weighted_target"`
}

type AppmeshRouteSpecSpecHttpRouteMatch struct {
	Prefix string `json:"prefix"`
}

type AppmeshRouteSpecSpecHttpRoute struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Action []AppmeshRouteSpecSpecHttpRoute `json:"action"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Match []AppmeshRouteSpecSpecHttpRoute `json:"match"`
}

type AppmeshRouteSpecSpecTcpRouteActionWeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int    `json:"weight"`
}

type AppmeshRouteSpecSpecTcpRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	WeightedTarget []AppmeshRouteSpecSpecTcpRouteAction `json:"weighted_target"`
}

type AppmeshRouteSpecSpecTcpRoute struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Action []AppmeshRouteSpecSpecTcpRoute `json:"action"`
}

type AppmeshRouteSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpRoute *[]AppmeshRouteSpecSpec `json:"http_route,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpRoute *[]AppmeshRouteSpecSpec `json:"tcp_route,omitempty"`
}

type AppmeshRouteSpec struct {
	MeshName string `json:"mesh_name"`
	Name     string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshRouteSpec `json:"spec"`
	// +optional
	Tags              map[string]string `json:"tags,omitempty"`
	VirtualRouterName string            `json:"virtual_router_name"`
}

type AppmeshRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppmeshRouteList is a list of AppmeshRoutes
type AppmeshRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppmeshRoute CRD objects
	Items []AppmeshRoute `json:"items,omitempty"`
}
