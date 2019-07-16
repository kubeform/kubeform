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

type PubsubSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSubscriptionSpec   `json:"spec,omitempty"`
	Status            PubsubSubscriptionStatus `json:"status,omitempty"`
}

type PubsubSubscriptionSpecPushConfig struct {
	// +optional
	Attributes   map[string]string `json:"attributes,omitempty"`
	PushEndpoint string            `json:"push_endpoint"`
}

type PubsubSubscriptionSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PushConfig *[]PubsubSubscriptionSpec `json:"push_config,omitempty"`
	Topic      string                    `json:"topic"`
}

type PubsubSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubSubscriptionList is a list of PubsubSubscriptions
type PubsubSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubSubscription CRD objects
	Items []PubsubSubscription `json:"items,omitempty"`
}
