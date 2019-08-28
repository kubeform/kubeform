package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StoragegatewayCachedIscsiVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayCachedIscsiVolumeSpec   `json:"spec,omitempty"`
	Status            StoragegatewayCachedIscsiVolumeStatus `json:"status,omitempty"`
}

type StoragegatewayCachedIscsiVolumeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	ChapEnabled bool   `json:"chapEnabled,omitempty" tf:"chap_enabled,omitempty"`
	GatewayArn  string `json:"gatewayArn" tf:"gateway_arn"`
	// +optional
	LunNumber          int    `json:"lunNumber,omitempty" tf:"lun_number,omitempty"`
	NetworkInterfaceID string `json:"networkInterfaceID" tf:"network_interface_id"`
	// +optional
	NetworkInterfacePort int `json:"networkInterfacePort,omitempty" tf:"network_interface_port,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	SourceVolumeArn string `json:"sourceVolumeArn,omitempty" tf:"source_volume_arn,omitempty"`
	// +optional
	TargetArn  string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
	TargetName string `json:"targetName" tf:"target_name"`
	// +optional
	VolumeArn string `json:"volumeArn,omitempty" tf:"volume_arn,omitempty"`
	// +optional
	VolumeID          string `json:"volumeID,omitempty" tf:"volume_id,omitempty"`
	VolumeSizeInBytes int    `json:"volumeSizeInBytes" tf:"volume_size_in_bytes"`
}



type StoragegatewayCachedIscsiVolumeStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *StoragegatewayCachedIscsiVolumeSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StoragegatewayCachedIscsiVolumeList is a list of StoragegatewayCachedIscsiVolumes
type StoragegatewayCachedIscsiVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StoragegatewayCachedIscsiVolume CRD objects
	Items []StoragegatewayCachedIscsiVolume `json:"items,omitempty"`
}