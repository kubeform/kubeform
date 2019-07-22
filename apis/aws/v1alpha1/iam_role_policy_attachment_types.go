package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamRolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamRolePolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            IamRolePolicyAttachmentStatus `json:"status,omitempty"`
}

type IamRolePolicyAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	PolicyArn   string                    `json:"policyArn" tf:"policy_arn"`
	Role        string                    `json:"role" tf:"role"`
}

type IamRolePolicyAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamRolePolicyAttachmentList is a list of IamRolePolicyAttachments
type IamRolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamRolePolicyAttachment CRD objects
	Items []IamRolePolicyAttachment `json:"items,omitempty"`
}
