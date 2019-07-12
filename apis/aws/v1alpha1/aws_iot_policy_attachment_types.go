package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsIotPolicyAttachmentStatus `json:"status,omitempty"`
}

type AwsIotPolicyAttachmentSpec struct {
	Policy string `json:"policy"`
	Target string `json:"target"`
}



type AwsIotPolicyAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotPolicyAttachmentList is a list of AwsIotPolicyAttachments
type AwsIotPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotPolicyAttachment CRD objects
	Items []AwsIotPolicyAttachment `json:"items,omitempty"`
}