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

type BillingAccountIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingAccountIamMemberSpec   `json:"spec,omitempty"`
	Status            BillingAccountIamMemberStatus `json:"status,omitempty"`
}

type BillingAccountIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BillingAccountID string `json:"billingAccountID" tf:"billing_account_id"`
	Member           string `json:"member" tf:"member"`
	Role             string `json:"role" tf:"role"`
}

type BillingAccountIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BillingAccountIamMemberList is a list of BillingAccountIamMembers
type BillingAccountIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BillingAccountIamMember CRD objects
	Items []BillingAccountIamMember `json:"items,omitempty"`
}
