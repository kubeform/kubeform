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

type Project struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec,omitempty"`
	Status            ProjectStatus `json:"status,omitempty"`
}

type ProjectSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Environment string `json:"environment,omitempty" tf:"environment,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	OwnerID int `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	OwnerUUID string `json:"ownerUUID,omitempty" tf:"owner_uuid,omitempty"`
	// +optional
	Purpose string `json:"purpose,omitempty" tf:"purpose,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Resources []string `json:"resources,omitempty" tf:"resources,omitempty"`
	// +optional
	UpdatedAt string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}



type ProjectStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ProjectSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectList is a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Project CRD objects
	Items []Project `json:"items,omitempty"`
}