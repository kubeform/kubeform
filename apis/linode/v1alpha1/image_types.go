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

type Image struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec,omitempty"`
	Status            ImageStatus `json:"status,omitempty"`
}

type ImageSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Created string `json:"created,omitempty" tf:"created,omitempty"`
	// +optional
	CreatedBy string `json:"createdBy,omitempty" tf:"created_by,omitempty"`
	// +optional
	Deprecated bool `json:"deprecated,omitempty" tf:"deprecated,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DiskID      int    `json:"diskID" tf:"disk_id"`
	// +optional
	Expiry string `json:"expiry,omitempty" tf:"expiry,omitempty"`
	// +optional
	IsPublic bool   `json:"isPublic,omitempty" tf:"is_public,omitempty"`
	Label    string `json:"label" tf:"label"`
	LinodeID int    `json:"linodeID" tf:"linode_id"`
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Vendor string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type ImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ImageSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ImageList is a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Image CRD objects
	Items []Image `json:"items,omitempty"`
}
