package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSharedImageGallery struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSharedImageGallerySpec   `json:"spec,omitempty"`
	Status            AzurermSharedImageGalleryStatus `json:"status,omitempty"`
}

type AzurermSharedImageGallerySpec struct {
	Tags              map[string]string `json:"tags"`
	UniqueName        string            `json:"unique_name"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Location          string            `json:"location"`
	Description       string            `json:"description"`
}

type AzurermSharedImageGalleryStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSharedImageGalleryList is a list of AzurermSharedImageGallerys
type AzurermSharedImageGalleryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSharedImageGallery CRD objects
	Items []AzurermSharedImageGallery `json:"items,omitempty"`
}
