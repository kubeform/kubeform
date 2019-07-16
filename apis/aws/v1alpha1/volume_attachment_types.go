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

type VolumeAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeAttachmentSpec   `json:"spec,omitempty"`
	Status            VolumeAttachmentStatus `json:"status,omitempty"`
}

type VolumeAttachmentSpec struct {
	DeviceName string `json:"device_name"`
	// +optional
	ForceDetach bool   `json:"force_detach,omitempty"`
	InstanceId  string `json:"instance_id"`
	// +optional
	SkipDestroy bool   `json:"skip_destroy,omitempty"`
	VolumeId    string `json:"volume_id"`
}

type VolumeAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VolumeAttachmentList is a list of VolumeAttachments
type VolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeAttachment CRD objects
	Items []VolumeAttachment `json:"items,omitempty"`
}
