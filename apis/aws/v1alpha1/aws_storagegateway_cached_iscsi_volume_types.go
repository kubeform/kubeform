package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewayCachedIscsiVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayCachedIscsiVolumeSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayCachedIscsiVolumeStatus `json:"status,omitempty"`
}

type AwsStoragegatewayCachedIscsiVolumeSpec struct {
	ChapEnabled          bool   `json:"chap_enabled"`
	GatewayArn           string `json:"gateway_arn"`
	LunNumber            int    `json:"lun_number"`
	NetworkInterfaceId   string `json:"network_interface_id"`
	NetworkInterfacePort int    `json:"network_interface_port"`
	SourceVolumeArn      string `json:"source_volume_arn"`
	TargetArn            string `json:"target_arn"`
	Arn                  string `json:"arn"`
	VolumeArn            string `json:"volume_arn"`
	VolumeId             string `json:"volume_id"`
	VolumeSizeInBytes    int    `json:"volume_size_in_bytes"`
	TargetName           string `json:"target_name"`
	SnapshotId           string `json:"snapshot_id"`
}

type AwsStoragegatewayCachedIscsiVolumeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewayCachedIscsiVolumeList is a list of AwsStoragegatewayCachedIscsiVolumes
type AwsStoragegatewayCachedIscsiVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayCachedIscsiVolume CRD objects
	Items []AwsStoragegatewayCachedIscsiVolume `json:"items,omitempty"`
}
