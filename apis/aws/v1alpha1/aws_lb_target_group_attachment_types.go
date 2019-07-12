package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroupAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbTargetGroupAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsLbTargetGroupAttachmentStatus `json:"status,omitempty"`
}

type AwsLbTargetGroupAttachmentSpec struct {
	TargetGroupArn   string `json:"target_group_arn"`
	TargetId         string `json:"target_id"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

type AwsLbTargetGroupAttachmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupAttachmentList is a list of AwsLbTargetGroupAttachments
type AwsLbTargetGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbTargetGroupAttachment CRD objects
	Items []AwsLbTargetGroupAttachment `json:"items,omitempty"`
}
