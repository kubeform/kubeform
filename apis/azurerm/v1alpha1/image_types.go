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

type ImageSpecDataDisk struct {
	// +optional
	BlobURI string `json:"blobURI,omitempty" tf:"blob_uri,omitempty"`
	// +optional
	Caching string `json:"caching,omitempty" tf:"caching,omitempty"`
	// +optional
	Lun int `json:"lun,omitempty" tf:"lun,omitempty"`
	// +optional
	ManagedDiskID string `json:"managedDiskID,omitempty" tf:"managed_disk_id,omitempty"`
	// +optional
	SizeGb int `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type ImageSpecOsDisk struct {
	// +optional
	BlobURI string `json:"blobURI,omitempty" tf:"blob_uri,omitempty"`
	// +optional
	Caching string `json:"caching,omitempty" tf:"caching,omitempty"`
	// +optional
	ManagedDiskID string `json:"managedDiskID,omitempty" tf:"managed_disk_id,omitempty"`
	// +optional
	OsState string `json:"osState,omitempty" tf:"os_state,omitempty"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	SizeGb int `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type ImageSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DataDisk []ImageSpecDataDisk `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`
	Location string              `json:"location" tf:"location"`
	Name     string              `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OsDisk            []ImageSpecOsDisk `json:"osDisk,omitempty" tf:"os_disk,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SourceVirtualMachineID string `json:"sourceVirtualMachineID,omitempty" tf:"source_virtual_machine_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ZoneResilient bool `json:"zoneResilient,omitempty" tf:"zone_resilient,omitempty"`
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
