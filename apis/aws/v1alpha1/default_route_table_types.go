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

type DefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultRouteTableSpec   `json:"spec,omitempty"`
	Status            DefaultRouteTableStatus `json:"status,omitempty"`
}

type DefaultRouteTableSpec struct {
	DefaultRouteTableId string `json:"default_route_table_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PropagatingVgws []string `json:"propagating_vgws,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DefaultRouteTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultRouteTableList is a list of DefaultRouteTables
type DefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultRouteTable CRD objects
	Items []DefaultRouteTable `json:"items,omitempty"`
}
