package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanVolumeAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanVolumeAttachmentSpec   `json:"spec,omitempty"`
	Status            DigitaloceanVolumeAttachmentStatus `json:"status,omitempty"`
}

type DigitaloceanVolumeAttachmentSpec struct {
	DropletId int    `json:"droplet_id"`
	VolumeId  string `json:"volume_id"`
}

type DigitaloceanVolumeAttachmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanVolumeAttachmentList is a list of DigitaloceanVolumeAttachments
type DigitaloceanVolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanVolumeAttachment CRD objects
	Items []DigitaloceanVolumeAttachment `json:"items,omitempty"`
}
