package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
}

type SharedImageSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Eula        string `json:"eula,omitempty"`
	GalleryName string `json:"gallery_name"`
	// +kubebuilder:validation:MaxItems=1
	Identifier []SharedImageSpec `json:"identifier"`
	Location   string            `json:"location"`
	Name       string            `json:"name"`
	OsType     string            `json:"os_type"`
	// +optional
	PrivacyStatementUri string `json:"privacy_statement_uri,omitempty"`
	// +optional
	ReleaseNoteUri    string `json:"release_note_uri,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
}

type SharedImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
