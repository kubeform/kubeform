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

type GooglePubsubSubscriptionIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubSubscriptionIamBindingSpec   `json:"spec,omitempty"`
	Status            GooglePubsubSubscriptionIamBindingStatus `json:"status,omitempty"`
}

type GooglePubsubSubscriptionIamBindingSpec struct {
	Subscription string   `json:"subscription"`
	Project      string   `json:"project"`
	Etag         string   `json:"etag"`
	Role         string   `json:"role"`
	Members      []string `json:"members"`
}

type GooglePubsubSubscriptionIamBindingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GooglePubsubSubscriptionIamBindingList is a list of GooglePubsubSubscriptionIamBindings
type GooglePubsubSubscriptionIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubSubscriptionIamBinding CRD objects
	Items []GooglePubsubSubscriptionIamBinding `json:"items,omitempty"`
}
