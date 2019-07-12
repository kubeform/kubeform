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

type AwsStoragegatewayCachedIscsiVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayCachedIscsiVolumeSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayCachedIscsiVolumeStatus `json:"status,omitempty"`
}

type AwsStoragegatewayCachedIscsiVolumeSpec struct {
	SourceVolumeArn      string `json:"source_volume_arn"`
	TargetName           string `json:"target_name"`
	GatewayArn           string `json:"gateway_arn"`
	NetworkInterfacePort int    `json:"network_interface_port"`
	SnapshotId           string `json:"snapshot_id"`
	NetworkInterfaceId   string `json:"network_interface_id"`
	TargetArn            string `json:"target_arn"`
	VolumeArn            string `json:"volume_arn"`
	VolumeId             string `json:"volume_id"`
	VolumeSizeInBytes    int    `json:"volume_size_in_bytes"`
	Arn                  string `json:"arn"`
	ChapEnabled          bool   `json:"chap_enabled"`
	LunNumber            int    `json:"lun_number"`
}

type AwsStoragegatewayCachedIscsiVolumeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsStoragegatewayCachedIscsiVolumeList is a list of AwsStoragegatewayCachedIscsiVolumes
type AwsStoragegatewayCachedIscsiVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayCachedIscsiVolume CRD objects
	Items []AwsStoragegatewayCachedIscsiVolume `json:"items,omitempty"`
}
