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

type AwsAppmeshRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppmeshRouteSpec   `json:"spec,omitempty"`
	Status            AwsAppmeshRouteStatus `json:"status,omitempty"`
}

type AwsAppmeshRouteSpecSpecHttpRouteMatch struct {
	Prefix string `json:"prefix"`
}

type AwsAppmeshRouteSpecSpecHttpRouteActionWeightedTarget struct {
	Weight      int    `json:"weight"`
	VirtualNode string `json:"virtual_node"`
}

type AwsAppmeshRouteSpecSpecHttpRouteAction struct {
	WeightedTarget []AwsAppmeshRouteSpecSpecHttpRouteAction `json:"weighted_target"`
}

type AwsAppmeshRouteSpecSpecHttpRoute struct {
	Match  []AwsAppmeshRouteSpecSpecHttpRoute `json:"match"`
	Action []AwsAppmeshRouteSpecSpecHttpRoute `json:"action"`
}

type AwsAppmeshRouteSpecSpecTcpRouteActionWeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int    `json:"weight"`
}

type AwsAppmeshRouteSpecSpecTcpRouteAction struct {
	WeightedTarget []AwsAppmeshRouteSpecSpecTcpRouteAction `json:"weighted_target"`
}

type AwsAppmeshRouteSpecSpecTcpRoute struct {
	Action []AwsAppmeshRouteSpecSpecTcpRoute `json:"action"`
}

type AwsAppmeshRouteSpecSpec struct {
	HttpRoute []AwsAppmeshRouteSpecSpec `json:"http_route"`
	TcpRoute  []AwsAppmeshRouteSpecSpec `json:"tcp_route"`
}

type AwsAppmeshRouteSpec struct {
	Arn               string                `json:"arn"`
	CreatedDate       string                `json:"created_date"`
	LastUpdatedDate   string                `json:"last_updated_date"`
	Tags              map[string]string     `json:"tags"`
	Name              string                `json:"name"`
	MeshName          string                `json:"mesh_name"`
	VirtualRouterName string                `json:"virtual_router_name"`
	Spec              []AwsAppmeshRouteSpec `json:"spec"`
}



type AwsAppmeshRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppmeshRouteList is a list of AwsAppmeshRoutes
type AwsAppmeshRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppmeshRoute CRD objects
	Items []AwsAppmeshRoute `json:"items,omitempty"`
}