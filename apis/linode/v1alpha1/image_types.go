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
// +kubebuilder:subresource:status

type Image struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec,omitempty"`
	Status            ImageStatus `json:"status,omitempty"`
}

type ImageSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// When this Image was created.
	// +optional
	Created string `json:"created,omitempty" tf:"created,omitempty"`
	// The name of the User who created this Image.
	// +optional
	CreatedBy string `json:"createdBy,omitempty" tf:"created_by,omitempty"`
	// Whether or not this Image is deprecated. Will only be True for deprecated public Images.
	// +optional
	Deprecated bool `json:"deprecated,omitempty" tf:"deprecated,omitempty"`
	// A detailed description of this Image.
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// The ID of the Linode Disk that this Image will be created from.
	DiskID int `json:"diskID" tf:"disk_id"`
	// Only Images created automatically (from a deleted Linode; type=automatic) will expire.
	// +optional
	Expiry string `json:"expiry,omitempty" tf:"expiry,omitempty"`
	// True if the Image is public.
	// +optional
	IsPublic bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`
	// A short description of the Image. Labels cannot contain special characters.
	Label string `json:"label" tf:"label"`
	// The ID of the Linode that this Image will be created from.
	LinodeID int `json:"linodeID" tf:"linode_id"`
	// The minimum size this Image needs to deploy. Size is in MB.
	// +optional
	Size int `json:"size,omitempty" tf:"size,omitempty"`
	// How the Image was created. 'Manual' Images can be created at any time. 'Automatic' images are created automatically from a deleted Linode.
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// The upstream distribution vendor. Nil for private Images.
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
	State *base.State `json:"state,omitempty"`
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
