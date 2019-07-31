package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LbTargetGroupAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbTargetGroupAttachmentSpec   `json:"spec,omitempty"`
	Status            LbTargetGroupAttachmentStatus `json:"status,omitempty"`
}

type LbTargetGroupAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	Port           int    `json:"port,omitempty" tf:"port,omitempty"`
	TargetGroupArn string `json:"targetGroupArn" tf:"target_group_arn"`
	TargetID       string `json:"targetID" tf:"target_id"`
}

type LbTargetGroupAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbTargetGroupAttachmentList is a list of LbTargetGroupAttachments
type LbTargetGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbTargetGroupAttachment CRD objects
	Items []LbTargetGroupAttachment `json:"items,omitempty"`
}
