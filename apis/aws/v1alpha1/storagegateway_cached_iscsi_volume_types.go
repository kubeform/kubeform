package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	GatewayArn         string `json:"gatewayArn" tf:"gateway_arn"`
	NetworkInterfaceID string `json:"networkInterfaceID" tf:"network_interface_id"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	SourceVolumeArn   string `json:"sourceVolumeArn,omitempty" tf:"source_volume_arn,omitempty"`
	TargetName        string `json:"targetName" tf:"target_name"`
	VolumeSizeInBytes int    `json:"volumeSizeInBytes" tf:"volume_size_in_bytes"`
}

type StoragegatewayCachedIscsiVolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
