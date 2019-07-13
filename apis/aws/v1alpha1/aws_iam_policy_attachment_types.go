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

type AwsIamPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsIamPolicyAttachmentStatus `json:"status,omitempty"`
}

type AwsIamPolicyAttachmentSpec struct {
	Name      string   `json:"name"`
	Users     []string `json:"users"`
	Roles     []string `json:"roles"`
	Groups    []string `json:"groups"`
	PolicyArn string   `json:"policy_arn"`
}



type AwsIamPolicyAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamPolicyAttachmentList is a list of AwsIamPolicyAttachments
type AwsIamPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamPolicyAttachment CRD objects
	Items []AwsIamPolicyAttachment `json:"items,omitempty"`
}