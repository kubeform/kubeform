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

type AwsGuarddutyMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGuarddutyMemberSpec   `json:"spec,omitempty"`
	Status            AwsGuarddutyMemberStatus `json:"status,omitempty"`
}

type AwsGuarddutyMemberSpec struct {
	AccountId                string `json:"account_id"`
	DetectorId               string `json:"detector_id"`
	Email                    string `json:"email"`
	RelationshipStatus       string `json:"relationship_status"`
	Invite                   bool   `json:"invite"`
	DisableEmailNotification bool   `json:"disable_email_notification"`
	InvitationMessage        string `json:"invitation_message"`
}



type AwsGuarddutyMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGuarddutyMemberList is a list of AwsGuarddutyMembers
type AwsGuarddutyMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGuarddutyMember CRD objects
	Items []AwsGuarddutyMember `json:"items,omitempty"`
}