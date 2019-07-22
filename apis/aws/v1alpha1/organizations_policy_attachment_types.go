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

type OrganizationsPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            OrganizationsPolicyAttachmentStatus `json:"status,omitempty"`
}

type OrganizationsPolicyAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	PolicyID    string                    `json:"policyID" tf:"policy_id"`
	TargetID    string                    `json:"targetID" tf:"target_id"`
}

type OrganizationsPolicyAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationsPolicyAttachmentList is a list of OrganizationsPolicyAttachments
type OrganizationsPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationsPolicyAttachment CRD objects
	Items []OrganizationsPolicyAttachment `json:"items,omitempty"`
}
