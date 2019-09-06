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

type Project struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec,omitempty"`
	Status            ProjectStatus `json:"status,omitempty"`
}

type ProjectSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// the date and time when the project was created, (ISO8601)
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// the descirption of the project
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// the environment of the project's resources
	// +optional
	Environment string `json:"environment,omitempty" tf:"environment,omitempty"`
	// the human-readable name for the project
	Name string `json:"name" tf:"name"`
	// the id of the project owner.
	// +optional
	OwnerID int `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// the unique universal identifier of the project owner.
	// +optional
	OwnerUUID string `json:"ownerUUID,omitempty" tf:"owner_uuid,omitempty"`
	// the purpose of the project
	// +optional
	Purpose string `json:"purpose,omitempty" tf:"purpose,omitempty"`
	// the resources associated with the project
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Resources []string `json:"resources,omitempty" tf:"resources,omitempty"`
	// the date and time when the project was last updated, (ISO8601)
	// +optional
	UpdatedAt string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
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
