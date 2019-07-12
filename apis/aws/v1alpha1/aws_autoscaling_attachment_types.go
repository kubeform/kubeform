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

type AwsAutoscalingAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingAttachmentStatus `json:"status,omitempty"`
}

type AwsAutoscalingAttachmentSpec struct {
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	Elb                  string `json:"elb"`
	AlbTargetGroupArn    string `json:"alb_target_group_arn"`
}

type AwsAutoscalingAttachmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAutoscalingAttachmentList is a list of AwsAutoscalingAttachments
type AwsAutoscalingAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingAttachment CRD objects
	Items []AwsAutoscalingAttachment `json:"items,omitempty"`
}
