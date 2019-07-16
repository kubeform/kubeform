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

type AlbTargetGroupAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbTargetGroupAttachmentSpec   `json:"spec,omitempty"`
	Status            AlbTargetGroupAttachmentStatus `json:"status,omitempty"`
}

type AlbTargetGroupAttachmentSpec struct {
	// +optional
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// +optional
	Port           int    `json:"port,omitempty"`
	TargetGroupArn string `json:"target_group_arn"`
	TargetId       string `json:"target_id"`
}

type AlbTargetGroupAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AlbTargetGroupAttachmentList is a list of AlbTargetGroupAttachments
type AlbTargetGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlbTargetGroupAttachment CRD objects
	Items []AlbTargetGroupAttachment `json:"items,omitempty"`
}
