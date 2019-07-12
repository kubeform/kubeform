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

type AwsIamUserPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamUserPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsIamUserPolicyAttachmentStatus `json:"status,omitempty"`
}

type AwsIamUserPolicyAttachmentSpec struct {
	User      string `json:"user"`
	PolicyArn string `json:"policy_arn"`
}

type AwsIamUserPolicyAttachmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamUserPolicyAttachmentList is a list of AwsIamUserPolicyAttachments
type AwsIamUserPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamUserPolicyAttachment CRD objects
	Items []AwsIamUserPolicyAttachment `json:"items,omitempty"`
}
