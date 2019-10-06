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

type SubnetRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            SubnetRouteTableAssociationStatus `json:"status,omitempty"`
}

type SubnetRouteTableAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RouteTableID string `json:"routeTableID" tf:"route_table_id"`
	SubnetID     string `json:"subnetID" tf:"subnet_id"`
}

type SubnetRouteTableAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SubnetRouteTableAssociationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubnetRouteTableAssociationList is a list of SubnetRouteTableAssociations
type SubnetRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubnetRouteTableAssociation CRD objects
	Items []SubnetRouteTableAssociation `json:"items,omitempty"`
}
