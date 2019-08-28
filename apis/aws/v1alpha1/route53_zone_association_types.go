package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Route53ZoneAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ZoneAssociationSpec   `json:"spec,omitempty"`
	Status            Route53ZoneAssociationStatus `json:"status,omitempty"`
}

type Route53ZoneAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	VpcID string `json:"vpcID" tf:"vpc_id"`
	// +optional
	VpcRegion string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`
	ZoneID    string `json:"zoneID" tf:"zone_id"`
}



type Route53ZoneAssociationStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *Route53ZoneAssociationSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ZoneAssociationList is a list of Route53ZoneAssociations
type Route53ZoneAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53ZoneAssociation CRD objects
	Items []Route53ZoneAssociation `json:"items,omitempty"`
}