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

type GuarddutyMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyMemberSpec   `json:"spec,omitempty"`
	Status            GuarddutyMemberStatus `json:"status,omitempty"`
}

type GuarddutyMemberSpec struct {
	AccountId  string `json:"account_id"`
	DetectorId string `json:"detector_id"`
	// +optional
	DisableEmailNotification bool   `json:"disable_email_notification,omitempty"`
	Email                    string `json:"email"`
	// +optional
	InvitationMessage string `json:"invitation_message,omitempty"`
	// +optional
	Invite bool `json:"invite,omitempty"`
}

type GuarddutyMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
