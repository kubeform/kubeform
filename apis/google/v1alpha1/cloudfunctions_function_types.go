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

type CloudfunctionsFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfunctionsFunctionSpec   `json:"spec,omitempty"`
	Status            CloudfunctionsFunctionStatus `json:"status,omitempty"`
}

type CloudfunctionsFunctionSpecEventTriggerFailurePolicy struct {
	Retry bool `json:"retry" tf:"retry"`
}

type CloudfunctionsFunctionSpecEventTrigger struct {
	EventType string `json:"eventType" tf:"event_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FailurePolicy []CloudfunctionsFunctionSpecEventTriggerFailurePolicy `json:"failurePolicy,omitempty" tf:"failure_policy,omitempty"`
	Resource      string                                                `json:"resource" tf:"resource"`
}

type CloudfunctionsFunctionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AvailableMemoryMb int `json:"availableMemoryMb,omitempty" tf:"available_memory_mb,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EntryPoint string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`
	// +optional
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EventTrigger []CloudfunctionsFunctionSpecEventTrigger `json:"eventTrigger,omitempty" tf:"event_trigger,omitempty"`
	// +optional
	HttpsTriggerURL string `json:"httpsTriggerURL,omitempty" tf:"https_trigger_url,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	// Deprecated
	RetryOnFailure bool `json:"retryOnFailure,omitempty" tf:"retry_on_failure,omitempty"`
	// +optional
	Runtime             string `json:"runtime,omitempty" tf:"runtime,omitempty"`
	SourceArchiveBucket string `json:"sourceArchiveBucket" tf:"source_archive_bucket"`
	SourceArchiveObject string `json:"sourceArchiveObject" tf:"source_archive_object"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
	// +optional
	// Deprecated
	TriggerBucket string `json:"triggerBucket,omitempty" tf:"trigger_bucket,omitempty"`
	// +optional
	TriggerHTTP bool `json:"triggerHTTP,omitempty" tf:"trigger_http,omitempty"`
	// +optional
	// Deprecated
	TriggerTopic string `json:"triggerTopic,omitempty" tf:"trigger_topic,omitempty"`
}

type CloudfunctionsFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudfunctionsFunctionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudfunctionsFunctionList is a list of CloudfunctionsFunctions
type CloudfunctionsFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudfunctionsFunction CRD objects
	Items []CloudfunctionsFunction `json:"items,omitempty"`
}
