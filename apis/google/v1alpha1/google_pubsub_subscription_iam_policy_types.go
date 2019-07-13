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

type GooglePubsubSubscriptionIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubSubscriptionIamPolicySpec   `json:"spec,omitempty"`
	Status            GooglePubsubSubscriptionIamPolicyStatus `json:"status,omitempty"`
}

type GooglePubsubSubscriptionIamPolicySpec struct {
	PolicyData   string `json:"policy_data"`
	Etag         string `json:"etag"`
	Subscription string `json:"subscription"`
	Project      string `json:"project"`
}



type GooglePubsubSubscriptionIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GooglePubsubSubscriptionIamPolicyList is a list of GooglePubsubSubscriptionIamPolicys
type GooglePubsubSubscriptionIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubSubscriptionIamPolicy CRD objects
	Items []GooglePubsubSubscriptionIamPolicy `json:"items,omitempty"`
}