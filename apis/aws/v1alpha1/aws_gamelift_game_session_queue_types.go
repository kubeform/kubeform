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

type AwsGameliftGameSessionQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGameliftGameSessionQueueSpec   `json:"spec,omitempty"`
	Status            AwsGameliftGameSessionQueueStatus `json:"status,omitempty"`
}

type AwsGameliftGameSessionQueueSpecPlayerLatencyPolicy struct {
	MaximumIndividualPlayerLatencyMilliseconds int `json:"maximum_individual_player_latency_milliseconds"`
	PolicyDurationSeconds                      int `json:"policy_duration_seconds"`
}

type AwsGameliftGameSessionQueueSpec struct {
	Destinations        []string                          `json:"destinations"`
	Name                string                            `json:"name"`
	PlayerLatencyPolicy []AwsGameliftGameSessionQueueSpec `json:"player_latency_policy"`
	TimeoutInSeconds    int                               `json:"timeout_in_seconds"`
	Arn                 string                            `json:"arn"`
}

type AwsGameliftGameSessionQueueStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGameliftGameSessionQueueList is a list of AwsGameliftGameSessionQueues
type AwsGameliftGameSessionQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGameliftGameSessionQueue CRD objects
	Items []AwsGameliftGameSessionQueue `json:"items,omitempty"`
}
