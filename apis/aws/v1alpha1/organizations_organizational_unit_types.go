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

type OrganizationsOrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsOrganizationalUnitSpec   `json:"spec,omitempty"`
	Status            OrganizationsOrganizationalUnitStatus `json:"status,omitempty"`
}

type OrganizationsOrganizationalUnitSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name     string `json:"name" tf:"name"`
	ParentID string `json:"parentID" tf:"parent_id"`
}

type OrganizationsOrganizationalUnitStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OrganizationsOrganizationalUnitSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationsOrganizationalUnitList is a list of OrganizationsOrganizationalUnits
type OrganizationsOrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationsOrganizationalUnit CRD objects
	Items []OrganizationsOrganizationalUnit `json:"items,omitempty"`
}
