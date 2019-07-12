package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	SourceArchiveObject  string                             `json:"source_archive_object"`
	Description          string                             `json:"description"`
	Timeout              int                                `json:"timeout"`
	TriggerTopic         string                             `json:"trigger_topic"`
	Project              string                             `json:"project"`
	EventTrigger         []GoogleCloudfunctionsFunctionSpec `json:"event_trigger"`
	HttpsTriggerUrl      string                             `json:"https_trigger_url"`
	Name                 string                             `json:"name"`
	SourceArchiveBucket  string                             `json:"source_archive_bucket"`
	AvailableMemoryMb    int                                `json:"available_memory_mb"`
	Runtime              string                             `json:"runtime"`
	TriggerBucket        string                             `json:"trigger_bucket"`
	TriggerHttp          bool                               `json:"trigger_http"`
	EntryPoint           string                             `json:"entry_point"`
	EnvironmentVariables map[string]string                  `json:"environment_variables"`
	RetryOnFailure       bool                               `json:"retry_on_failure"`
	Labels               map[string]string                  `json:"labels"`
	Region               string                             `json:"region"`
}



type GoogleCloudfunctionsFunctionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudfunctionsFunctionList is a list of GoogleCloudfunctionsFunctions
type GoogleCloudfunctionsFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleCloudfunctionsFunction CRD objects
	Items []GoogleCloudfunctionsFunction `json:"items,omitempty"`
}