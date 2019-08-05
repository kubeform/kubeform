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

type StorageObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageObjectAccessControlSpec   `json:"spec,omitempty"`
	Status            StorageObjectAccessControlStatus `json:"status,omitempty"`
}

type StorageObjectAccessControlSpecProjectTeam struct {
	// +optional
	ProjectNumber string `json:"projectNumber,omitempty" tf:"project_number,omitempty"`
	// +optional
	Team string `json:"team,omitempty" tf:"team,omitempty"`
}

type StorageObjectAccessControlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Bucket     string `json:"bucket" tf:"bucket"`
	Domain     string `json:"domain" tf:"domain"`
	Email      string `json:"email" tf:"email"`
	Entity     string `json:"entity" tf:"entity"`
	EntityID   string `json:"entityID" tf:"entity_id"`
	Generation int    `json:"generation" tf:"generation"`
	Object     string `json:"object" tf:"object"`
	// +kubebuilder:validation:MaxItems=1
	ProjectTeam []StorageObjectAccessControlSpecProjectTeam `json:"projectTeam" tf:"project_team"`
	Role        string                                      `json:"role" tf:"role"`
}

type StorageObjectAccessControlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageObjectAccessControlSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageObjectAccessControlList is a list of StorageObjectAccessControls
type StorageObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageObjectAccessControl CRD objects
	Items []StorageObjectAccessControl `json:"items,omitempty"`
}
