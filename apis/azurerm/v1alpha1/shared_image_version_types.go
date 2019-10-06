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

type SharedImageVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedImageVersionSpec   `json:"spec,omitempty"`
	Status            SharedImageVersionStatus `json:"status,omitempty"`
}

type SharedImageVersionSpecTargetRegion struct {
	Name                 string `json:"name" tf:"name"`
	RegionalReplicaCount int    `json:"regionalReplicaCount" tf:"regional_replica_count"`
}

type SharedImageVersionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ExcludeFromLatest bool   `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest,omitempty"`
	GalleryName       string `json:"galleryName" tf:"gallery_name"`
	ImageName         string `json:"imageName" tf:"image_name"`
	Location          string `json:"location" tf:"location"`
	ManagedImageID    string `json:"managedImageID" tf:"managed_image_id"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags         map[string]string                    `json:"tags,omitempty" tf:"tags,omitempty"`
	TargetRegion []SharedImageVersionSpecTargetRegion `json:"targetRegion" tf:"target_region"`
}

type SharedImageVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SharedImageVersionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SharedImageVersionList is a list of SharedImageVersions
type SharedImageVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SharedImageVersion CRD objects
	Items []SharedImageVersion `json:"items,omitempty"`
}
