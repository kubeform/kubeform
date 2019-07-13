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

type AwsGuarddutyInviteAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGuarddutyInviteAccepterSpec   `json:"spec,omitempty"`
	Status            AwsGuarddutyInviteAccepterStatus `json:"status,omitempty"`
}

type AwsGuarddutyInviteAccepterSpec struct {
	DetectorId      string `json:"detector_id"`
	MasterAccountId string `json:"master_account_id"`
}



type AwsGuarddutyInviteAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGuarddutyInviteAccepterList is a list of AwsGuarddutyInviteAccepters
type AwsGuarddutyInviteAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGuarddutyInviteAccepter CRD objects
	Items []AwsGuarddutyInviteAccepter `json:"items,omitempty"`
}