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

type MacieMemberAccountAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MacieMemberAccountAssociationSpec   `json:"spec,omitempty"`
	Status            MacieMemberAccountAssociationStatus `json:"status,omitempty"`
}

type MacieMemberAccountAssociationSpec struct {
	MemberAccountId string `json:"member_account_id"`
}

type MacieMemberAccountAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MacieMemberAccountAssociationList is a list of MacieMemberAccountAssociations
type MacieMemberAccountAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MacieMemberAccountAssociation CRD objects
	Items []MacieMemberAccountAssociation `json:"items,omitempty"`
}
