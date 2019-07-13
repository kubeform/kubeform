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

type GoogleBillingAccountIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBillingAccountIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleBillingAccountIamBindingStatus `json:"status,omitempty"`
}

type GoogleBillingAccountIamBindingSpec struct {
	BillingAccountId string   `json:"billing_account_id"`
	Role             string   `json:"role"`
	Members          []string `json:"members"`
	Etag             string   `json:"etag"`
}



type GoogleBillingAccountIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBillingAccountIamBindingList is a list of GoogleBillingAccountIamBindings
type GoogleBillingAccountIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBillingAccountIamBinding CRD objects
	Items []GoogleBillingAccountIamBinding `json:"items,omitempty"`
}