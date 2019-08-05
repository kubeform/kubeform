package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LambdaEventSourceMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaEventSourceMappingSpec   `json:"spec,omitempty"`
	Status            LambdaEventSourceMappingStatus `json:"status,omitempty"`
}

type LambdaEventSourceMappingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	BatchSize int `json:"batchSize,omitempty" tf:"batch_size,omitempty"`
	// +optional
	Enabled              bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	EventSourceArn       string `json:"eventSourceArn" tf:"event_source_arn"`
	FunctionArn          string `json:"functionArn" tf:"function_arn"`
	FunctionName         string `json:"functionName" tf:"function_name"`
	LastModified         string `json:"lastModified" tf:"last_modified"`
	LastProcessingResult string `json:"lastProcessingResult" tf:"last_processing_result"`
	// +optional
	StartingPosition string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`
	// +optional
	StartingPositionTimestamp string `json:"startingPositionTimestamp,omitempty" tf:"starting_position_timestamp,omitempty"`
	State                     string `json:"state" tf:"state"`
	StateTransitionReason     string `json:"stateTransitionReason" tf:"state_transition_reason"`
	Uuid                      string `json:"uuid" tf:"uuid"`
}

type LambdaEventSourceMappingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LambdaEventSourceMappingSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaEventSourceMappingList is a list of LambdaEventSourceMappings
type LambdaEventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaEventSourceMapping CRD objects
	Items []LambdaEventSourceMapping `json:"items,omitempty"`
}
