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

type SharedImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedImageSpec   `json:"spec,omitempty"`
	Status            SharedImageStatus `json:"status,omitempty"`
}

type SharedImageSpecIdentifier struct {
	Offer     string `json:"offer" tf:"offer"`
	Publisher string `json:"publisher" tf:"publisher"`
	Sku       string `json:"sku" tf:"sku"`
}

type SharedImageSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Eula        string `json:"eula,omitempty" tf:"eula,omitempty"`
	GalleryName string `json:"galleryName" tf:"gallery_name"`
	// +kubebuilder:validation:MaxItems=1
	Identifier []SharedImageSpecIdentifier `json:"identifier" tf:"identifier"`
	Location   string                      `json:"location" tf:"location"`
	Name       string                      `json:"name" tf:"name"`
	OsType     string                      `json:"osType" tf:"os_type"`
	// +optional
	PrivacyStatementURI string `json:"privacyStatementURI,omitempty" tf:"privacy_statement_uri,omitempty"`
	// +optional
	ReleaseNoteURI    string `json:"releaseNoteURI,omitempty" tf:"release_note_uri,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SharedImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SharedImageSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SharedImageList is a list of SharedImages
type SharedImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SharedImage CRD objects
	Items []SharedImage `json:"items,omitempty"`
}
