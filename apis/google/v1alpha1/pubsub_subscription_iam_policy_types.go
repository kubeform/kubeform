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

type PubsubSubscriptionIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSubscriptionIamPolicySpec   `json:"spec,omitempty"`
	Status            PubsubSubscriptionIamPolicyStatus `json:"status,omitempty"`
}

type PubsubSubscriptionIamPolicySpec struct {
	PolicyData   string `json:"policy_data"`
	Subscription string `json:"subscription"`
}

type PubsubSubscriptionIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubSubscriptionIamPolicyList is a list of PubsubSubscriptionIamPolicys
type PubsubSubscriptionIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubSubscriptionIamPolicy CRD objects
	Items []PubsubSubscriptionIamPolicy `json:"items,omitempty"`
}
