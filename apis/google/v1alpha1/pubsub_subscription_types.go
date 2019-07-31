package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	Attributes   map[string]string `json:"attributes,omitempty" tf:"attributes,omitempty"`
	PushEndpoint string            `json:"pushEndpoint" tf:"push_endpoint"`
}

type PubsubSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AckDeadlineSeconds int    `json:"ackDeadlineSeconds,omitempty" tf:"ack_deadline_seconds,omitempty"`
	Name               string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PushConfig []PubsubSubscriptionSpecPushConfig `json:"pushConfig,omitempty" tf:"push_config,omitempty"`
	Topic      string                             `json:"topic" tf:"topic"`
}

type PubsubSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
