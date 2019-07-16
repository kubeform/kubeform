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

type BatchJobDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchJobDefinitionSpec   `json:"spec,omitempty"`
	Status            BatchJobDefinitionStatus `json:"status,omitempty"`
}

type BatchJobDefinitionSpecRetryStrategy struct {
	// +optional
	Attempts int `json:"attempts,omitempty"`
}

type BatchJobDefinitionSpecTimeout struct {
	// +optional
	AttemptDurationSeconds int `json:"attempt_duration_seconds,omitempty"`
}

type BatchJobDefinitionSpec struct {
	// +optional
	ContainerProperties string `json:"container_properties,omitempty"`
	Name                string `json:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetryStrategy *[]BatchJobDefinitionSpec `json:"retry_strategy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Timeout *[]BatchJobDefinitionSpec `json:"timeout,omitempty"`
	Type    string                    `json:"type"`
}

type BatchJobDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchJobDefinitionList is a list of BatchJobDefinitions
type BatchJobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchJobDefinition CRD objects
	Items []BatchJobDefinition `json:"items,omitempty"`
}
