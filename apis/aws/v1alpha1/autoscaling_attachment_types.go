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

type AutoscalingAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingAttachmentSpec   `json:"spec,omitempty"`
	Status            AutoscalingAttachmentStatus `json:"status,omitempty"`
}

type AutoscalingAttachmentSpec struct {
	// +optional
	AlbTargetGroupArn    string `json:"alb_target_group_arn,omitempty"`
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	// +optional
	Elb string `json:"elb,omitempty"`
}

type AutoscalingAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
