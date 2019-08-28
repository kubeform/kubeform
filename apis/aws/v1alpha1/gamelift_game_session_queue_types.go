package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GameliftGameSessionQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftGameSessionQueueSpec   `json:"spec,omitempty"`
	Status            GameliftGameSessionQueueStatus `json:"status,omitempty"`
}

type GameliftGameSessionQueueSpecPlayerLatencyPolicy struct {
	MaximumIndividualPlayerLatencyMilliseconds int `json:"maximumIndividualPlayerLatencyMilliseconds" tf:"maximum_individual_player_latency_milliseconds"`
	// +optional
	PolicyDurationSeconds int `json:"policyDurationSeconds,omitempty" tf:"policy_duration_seconds,omitempty"`
}

type GameliftGameSessionQueueSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Destinations []string `json:"destinations,omitempty" tf:"destinations,omitempty"`
	Name         string   `json:"name" tf:"name"`
	// +optional
	PlayerLatencyPolicy []GameliftGameSessionQueueSpecPlayerLatencyPolicy `json:"playerLatencyPolicy,omitempty" tf:"player_latency_policy,omitempty"`
	// +optional
	TimeoutInSeconds int `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`
}



type GameliftGameSessionQueueStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *GameliftGameSessionQueueSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GameliftGameSessionQueueList is a list of GameliftGameSessionQueues
type GameliftGameSessionQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GameliftGameSessionQueue CRD objects
	Items []GameliftGameSessionQueue `json:"items,omitempty"`
}