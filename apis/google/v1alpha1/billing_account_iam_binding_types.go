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

type BillingAccountIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingAccountIamBindingSpec   `json:"spec,omitempty"`
	Status            BillingAccountIamBindingStatus `json:"status,omitempty"`
}

type BillingAccountIamBindingSpec struct {
	BillingAccountId string `json:"billing_account_id"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members"`
	Role    string   `json:"role"`
}

type BillingAccountIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BillingAccountIamBindingList is a list of BillingAccountIamBindings
type BillingAccountIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BillingAccountIamBinding CRD objects
	Items []BillingAccountIamBinding `json:"items,omitempty"`
}
