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

type AzurermSharedImageVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSharedImageVersionSpec   `json:"spec,omitempty"`
	Status            AzurermSharedImageVersionStatus `json:"status,omitempty"`
}

type AzurermSharedImageVersionSpecTargetRegion struct {
	Name                 string `json:"name"`
	RegionalReplicaCount int    `json:"regional_replica_count"`
}

type AzurermSharedImageVersionSpec struct {
	TargetRegion      []AzurermSharedImageVersionSpec `json:"target_region"`
	Location          string                          `json:"location"`
	GalleryName       string                          `json:"gallery_name"`
	ImageName         string                          `json:"image_name"`
	ResourceGroupName string                          `json:"resource_group_name"`
	ManagedImageId    string                          `json:"managed_image_id"`
	ExcludeFromLatest bool                            `json:"exclude_from_latest"`
	Tags              map[string]string               `json:"tags"`
	Name              string                          `json:"name"`
}



type AzurermSharedImageVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSharedImageVersionList is a list of AzurermSharedImageVersions
type AzurermSharedImageVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSharedImageVersion CRD objects
	Items []AzurermSharedImageVersion `json:"items,omitempty"`
}