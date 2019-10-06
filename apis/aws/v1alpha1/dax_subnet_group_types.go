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

type DaxSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DaxSubnetGroupSpec   `json:"spec,omitempty"`
	Status            DaxSubnetGroupStatus `json:"status,omitempty"`
}

type DaxSubnetGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string   `json:"description,omitempty" tf:"description,omitempty"`
	Name        string   `json:"name" tf:"name"`
	SubnetIDS   []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type DaxSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DaxSubnetGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DaxSubnetGroupList is a list of DaxSubnetGroups
type DaxSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DaxSubnetGroup CRD objects
	Items []DaxSubnetGroup `json:"items,omitempty"`
}
