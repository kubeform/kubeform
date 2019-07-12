package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOrganizationsPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsOrganizationsPolicyAttachmentStatus `json:"status,omitempty"`
}

type AwsOrganizationsPolicyAttachmentSpec struct {
	PolicyId string `json:"policy_id"`
	TargetId string `json:"target_id"`
}



type AwsOrganizationsPolicyAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsPolicyAttachmentList is a list of AwsOrganizationsPolicyAttachments
type AwsOrganizationsPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOrganizationsPolicyAttachment CRD objects
	Items []AwsOrganizationsPolicyAttachment `json:"items,omitempty"`
}