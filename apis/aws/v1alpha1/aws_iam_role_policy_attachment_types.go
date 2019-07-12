package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamRolePolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsIamRolePolicyAttachmentStatus `json:"status,omitempty"`
}

type AwsIamRolePolicyAttachmentSpec struct {
	Role      string `json:"role"`
	PolicyArn string `json:"policy_arn"`
}



type AwsIamRolePolicyAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRolePolicyAttachmentList is a list of AwsIamRolePolicyAttachments
type AwsIamRolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamRolePolicyAttachment CRD objects
	Items []AwsIamRolePolicyAttachment `json:"items,omitempty"`
}