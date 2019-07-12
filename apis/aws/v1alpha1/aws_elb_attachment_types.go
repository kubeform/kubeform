package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElbAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElbAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsElbAttachmentStatus `json:"status,omitempty"`
}

type AwsElbAttachmentSpec struct {
	Elb      string `json:"elb"`
	Instance string `json:"instance"`
}



type AwsElbAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbAttachmentList is a list of AwsElbAttachments
type AwsElbAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElbAttachment CRD objects
	Items []AwsElbAttachment `json:"items,omitempty"`
}