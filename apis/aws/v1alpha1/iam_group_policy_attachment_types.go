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

type IamGroupPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamGroupPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            IamGroupPolicyAttachmentStatus `json:"status,omitempty"`
}

type IamGroupPolicyAttachmentSpec struct {
	Group     string `json:"group"`
	PolicyArn string `json:"policy_arn"`
}

type IamGroupPolicyAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamGroupPolicyAttachmentList is a list of IamGroupPolicyAttachments
type IamGroupPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamGroupPolicyAttachment CRD objects
	Items []IamGroupPolicyAttachment `json:"items,omitempty"`
}
