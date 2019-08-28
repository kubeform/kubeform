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

type LambdaEventSourceMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaEventSourceMappingSpec   `json:"spec,omitempty"`
	Status            LambdaEventSourceMappingStatus `json:"status,omitempty"`
}

type LambdaEventSourceMappingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BatchSize int `json:"batchSize,omitempty" tf:"batch_size,omitempty"`
	// +optional
	Enabled        bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	EventSourceArn string `json:"eventSourceArn" tf:"event_source_arn"`
	// +optional
	FunctionArn  string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`
	FunctionName string `json:"functionName" tf:"function_name"`
	// +optional
	LastModified string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`
	// +optional
	LastProcessingResult string `json:"lastProcessingResult,omitempty" tf:"last_processing_result,omitempty"`
	// +optional
	StartingPosition string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`
	// +optional
	StartingPositionTimestamp string `json:"startingPositionTimestamp,omitempty" tf:"starting_position_timestamp,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	StateTransitionReason string `json:"stateTransitionReason,omitempty" tf:"state_transition_reason,omitempty"`
	// +optional
	Uuid string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}



type LambdaEventSourceMappingStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
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