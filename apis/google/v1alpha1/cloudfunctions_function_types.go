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

type CloudfunctionsFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfunctionsFunctionSpec   `json:"spec,omitempty"`
	Status            CloudfunctionsFunctionStatus `json:"status,omitempty"`
}

type CloudfunctionsFunctionSpec struct {
	// +optional
	AvailableMemoryMb int `json:"available_memory_mb,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EntryPoint string `json:"entry_point,omitempty"`
	// +optional
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
	// +optional
	Labels              map[string]string `json:"labels,omitempty"`
	Name                string            `json:"name"`
	SourceArchiveBucket string            `json:"source_archive_bucket"`
	SourceArchiveObject string            `json:"source_archive_object"`
	// +optional
	Timeout int `json:"timeout,omitempty"`
	// +optional
	TriggerHttp bool `json:"trigger_http,omitempty"`
}

type CloudfunctionsFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
