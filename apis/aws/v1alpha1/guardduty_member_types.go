package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type GuarddutyMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyMemberSpec   `json:"spec,omitempty"`
	Status            GuarddutyMemberStatus `json:"status,omitempty"`
}

type GuarddutyMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountID  string `json:"accountID" tf:"account_id"`
	DetectorID string `json:"detectorID" tf:"detector_id"`
	// +optional
	DisableEmailNotification bool   `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`
	Email                    string `json:"email" tf:"email"`
	// +optional
	InvitationMessage string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`
	// +optional
	Invite bool `json:"invite,omitempty" tf:"invite,omitempty"`
	// +optional
	RelationshipStatus string `json:"relationshipStatus,omitempty" tf:"relationship_status,omitempty"`
}

type GuarddutyMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GuarddutyMemberSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyMemberList is a list of GuarddutyMembers
type GuarddutyMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyMember CRD objects
	Items []GuarddutyMember `json:"items,omitempty"`
}
