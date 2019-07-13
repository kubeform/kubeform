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

type AwsIamGroupPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamGroupPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsIamGroupPolicyAttachmentStatus `json:"status,omitempty"`
}

type AwsIamGroupPolicyAttachmentSpec struct {
	Group     string `json:"group"`
	PolicyArn string `json:"policy_arn"`
}



type AwsIamGroupPolicyAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamGroupPolicyAttachmentList is a list of AwsIamGroupPolicyAttachments
type AwsIamGroupPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamGroupPolicyAttachment CRD objects
	Items []AwsIamGroupPolicyAttachment `json:"items,omitempty"`
}