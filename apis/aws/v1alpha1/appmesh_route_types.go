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

type AppmeshRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshRouteSpec   `json:"spec,omitempty"`
	Status            AppmeshRouteStatus `json:"status,omitempty"`
}

type AppmeshRouteSpecSpecHttpRouteActionWeightedTarget struct {
	VirtualNode string `json:"virtualNode" tf:"virtual_node"`
	Weight      int    `json:"weight" tf:"weight"`
}

type AppmeshRouteSpecSpecHttpRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	WeightedTarget []AppmeshRouteSpecSpecHttpRouteActionWeightedTarget `json:"weightedTarget" tf:"weighted_target"`
}

type AppmeshRouteSpecSpecHttpRouteMatch struct {
	Prefix string `json:"prefix" tf:"prefix"`
}

type AppmeshRouteSpecSpecHttpRoute struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Action []AppmeshRouteSpecSpecHttpRouteAction `json:"action" tf:"action"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Match []AppmeshRouteSpecSpecHttpRouteMatch `json:"match" tf:"match"`
}

type AppmeshRouteSpecSpecTcpRouteActionWeightedTarget struct {
	VirtualNode string `json:"virtualNode" tf:"virtual_node"`
	Weight      int    `json:"weight" tf:"weight"`
}

type AppmeshRouteSpecSpecTcpRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	WeightedTarget []AppmeshRouteSpecSpecTcpRouteActionWeightedTarget `json:"weightedTarget" tf:"weighted_target"`
}

type AppmeshRouteSpecSpecTcpRoute struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Action []AppmeshRouteSpecSpecTcpRouteAction `json:"action" tf:"action"`
}

type AppmeshRouteSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpRoute []AppmeshRouteSpecSpecHttpRoute `json:"httpRoute,omitempty" tf:"http_route,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpRoute []AppmeshRouteSpecSpecTcpRoute `json:"tcpRoute,omitempty" tf:"tcp_route,omitempty"`
}

type AppmeshRouteSpec struct {
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
	Spec []AppmeshRouteSpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags              map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VirtualRouterName string            `json:"virtualRouterName" tf:"virtual_router_name"`
}

type AppmeshRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppmeshRouteSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
