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

type AzurermSharedImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSharedImageSpec   `json:"spec,omitempty"`
	Status            AzurermSharedImageStatus `json:"status,omitempty"`
}

type AzurermSharedImageSpecIdentifier struct {
	Publisher string `json:"publisher"`
	Offer     string `json:"offer"`
	Sku       string `json:"sku"`
}

type AzurermSharedImageSpec struct {
	GalleryName         string                   `json:"gallery_name"`
	ResourceGroupName   string                   `json:"resource_group_name"`
	Identifier          []AzurermSharedImageSpec `json:"identifier"`
	ReleaseNoteUri      string                   `json:"release_note_uri"`
	Tags                map[string]string        `json:"tags"`
	Name                string                   `json:"name"`
	Location            string                   `json:"location"`
	OsType              string                   `json:"os_type"`
	Description         string                   `json:"description"`
	Eula                string                   `json:"eula"`
	PrivacyStatementUri string                   `json:"privacy_statement_uri"`
}

type AzurermSharedImageStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSharedImageList is a list of AzurermSharedImages
type AzurermSharedImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSharedImage CRD objects
	Items []AzurermSharedImage `json:"items,omitempty"`
}
