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

type AwsVolumeAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVolumeAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsVolumeAttachmentStatus `json:"status,omitempty"`
}

type AwsVolumeAttachmentSpec struct {
	DeviceName  string `json:"device_name"`
	InstanceId  string `json:"instance_id"`
	VolumeId    string `json:"volume_id"`
	ForceDetach bool   `json:"force_detach"`
	SkipDestroy bool   `json:"skip_destroy"`
}



type AwsVolumeAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVolumeAttachmentList is a list of AwsVolumeAttachments
type AwsVolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVolumeAttachment CRD objects
	Items []AwsVolumeAttachment `json:"items,omitempty"`
}