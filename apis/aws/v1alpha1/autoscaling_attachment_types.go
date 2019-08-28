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

type AutoscalingAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingAttachmentSpec   `json:"spec,omitempty"`
	Status            AutoscalingAttachmentStatus `json:"status,omitempty"`
}

type AutoscalingAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AlbTargetGroupArn    string `json:"albTargetGroupArn,omitempty" tf:"alb_target_group_arn,omitempty"`
	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`
	// +optional
	Elb string `json:"elb,omitempty" tf:"elb,omitempty"`
}



type AutoscalingAttachmentStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AutoscalingAttachmentSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingAttachmentList is a list of AutoscalingAttachments
type AutoscalingAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingAttachment CRD objects
	Items []AutoscalingAttachment `json:"items,omitempty"`
}