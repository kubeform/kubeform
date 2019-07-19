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

type GuarddutyInviteAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyInviteAccepterSpec   `json:"spec,omitempty"`
	Status            GuarddutyInviteAccepterStatus `json:"status,omitempty"`
}

type GuarddutyInviteAccepterSpec struct {
	DetectorID      string                    `json:"detectorID" tf:"detector_id"`
	MasterAccountID string                    `json:"masterAccountID" tf:"master_account_id"`
	ProviderRef     core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GuarddutyInviteAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyInviteAccepterList is a list of GuarddutyInviteAccepters
type GuarddutyInviteAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyInviteAccepter CRD objects
	Items []GuarddutyInviteAccepter `json:"items,omitempty"`
}
