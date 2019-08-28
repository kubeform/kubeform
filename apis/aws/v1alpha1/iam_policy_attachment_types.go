package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            IamPolicyAttachmentStatus `json:"status,omitempty"`
}

type IamPolicyAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Groups    []string `json:"groups,omitempty" tf:"groups,omitempty"`
	Name      string   `json:"name" tf:"name"`
	PolicyArn string   `json:"policyArn" tf:"policy_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Roles []string `json:"roles,omitempty" tf:"roles,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Users []string `json:"users,omitempty" tf:"users,omitempty"`
}



type IamPolicyAttachmentStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *IamPolicyAttachmentSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamPolicyAttachmentList is a list of IamPolicyAttachments
type IamPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamPolicyAttachment CRD objects
	Items []IamPolicyAttachment `json:"items,omitempty"`
}