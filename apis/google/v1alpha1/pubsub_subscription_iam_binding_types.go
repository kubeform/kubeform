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

type PubsubSubscriptionIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSubscriptionIamBindingSpec   `json:"spec,omitempty"`
	Status            PubsubSubscriptionIamBindingStatus `json:"status,omitempty"`
}

type PubsubSubscriptionIamBindingSpec struct {
	// +kubebuilder:validation:UniqueItems=true
	Members      []string `json:"members"`
	Role         string   `json:"role"`
	Subscription string   `json:"subscription"`
}

type PubsubSubscriptionIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubSubscriptionIamBindingList is a list of PubsubSubscriptionIamBindings
type PubsubSubscriptionIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubSubscriptionIamBinding CRD objects
	Items []PubsubSubscriptionIamBinding `json:"items,omitempty"`
}
