package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBatchJobDefinitionSpec   `json:"spec,omitempty"`
	Status            AwsBatchJobDefinitionStatus `json:"status,omitempty"`
}

type AwsBatchJobDefinitionSpecRetryStrategy struct {
	Attempts int `json:"attempts"`
}

type AwsBatchJobDefinitionSpecTimeout struct {
	AttemptDurationSeconds int `json:"attempt_duration_seconds"`
}

type AwsBatchJobDefinitionSpec struct {
	Parameters          map[string]string           `json:"parameters"`
	RetryStrategy       []AwsBatchJobDefinitionSpec `json:"retry_strategy"`
	Timeout             []AwsBatchJobDefinitionSpec `json:"timeout"`
	Type                string                      `json:"type"`
	Revision            int                         `json:"revision"`
	Arn                 string                      `json:"arn"`
	Name                string                      `json:"name"`
	ContainerProperties string                      `json:"container_properties"`
}



type AwsBatchJobDefinitionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobDefinitionList is a list of AwsBatchJobDefinitions
type AwsBatchJobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBatchJobDefinition CRD objects
	Items []AwsBatchJobDefinition `json:"items,omitempty"`
}