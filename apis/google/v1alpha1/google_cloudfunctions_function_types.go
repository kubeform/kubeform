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

type GoogleCloudfunctionsFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleCloudfunctionsFunctionSpec   `json:"spec,omitempty"`
	Status            GoogleCloudfunctionsFunctionStatus `json:"status,omitempty"`
}

type GoogleCloudfunctionsFunctionSpecEventTriggerFailurePolicy struct {
	Retry bool `json:"retry"`
}

type GoogleCloudfunctionsFunctionSpecEventTrigger struct {
	EventType     string                                         `json:"event_type"`
	Resource      string                                         `json:"resource"`
	FailurePolicy []GoogleCloudfunctionsFunctionSpecEventTrigger `json:"failure_policy"`
}

type GoogleCloudfunctionsFunctionSpec struct {
	TriggerBucket        string                             `json:"trigger_bucket"`
	TriggerHttp          bool                               `json:"trigger_http"`
	Project              string                             `json:"project"`
	AvailableMemoryMb    int                                `json:"available_memory_mb"`
	Timeout              int                                `json:"timeout"`
	Labels               map[string]string                  `json:"labels"`
	Runtime              string                             `json:"runtime"`
	Name                 string                             `json:"name"`
	Description          string                             `json:"description"`
	HttpsTriggerUrl      string                             `json:"https_trigger_url"`
	EntryPoint           string                             `json:"entry_point"`
	TriggerTopic         string                             `json:"trigger_topic"`
	Region               string                             `json:"region"`
	RetryOnFailure       bool                               `json:"retry_on_failure"`
	SourceArchiveBucket  string                             `json:"source_archive_bucket"`
	SourceArchiveObject  string                             `json:"source_archive_object"`
	EnvironmentVariables map[string]string                  `json:"environment_variables"`
	EventTrigger         []GoogleCloudfunctionsFunctionSpec `json:"event_trigger"`
}



type GoogleCloudfunctionsFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleCloudfunctionsFunctionList is a list of GoogleCloudfunctionsFunctions
type GoogleCloudfunctionsFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleCloudfunctionsFunction CRD objects
	Items []GoogleCloudfunctionsFunction `json:"items,omitempty"`
}