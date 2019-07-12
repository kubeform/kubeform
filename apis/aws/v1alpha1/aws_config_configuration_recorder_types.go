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

type AwsConfigConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsConfigConfigurationRecorderSpec    `json:"spec,omitempty"`
	Status            AwsConfigConfigurationRecorderStatus2 `json:"status,omitempty"`
}

type AwsConfigConfigurationRecorderSpecRecordingGroup struct {
	AllSupported               bool     `json:"all_supported"`
	IncludeGlobalResourceTypes bool     `json:"include_global_resource_types"`
	ResourceTypes              []string `json:"resource_types"`
}

type AwsConfigConfigurationRecorderSpec struct {
	Name           string                               `json:"name"`
	RoleArn        string                               `json:"role_arn"`
	RecordingGroup []AwsConfigConfigurationRecorderSpec `json:"recording_group"`
}

type AwsConfigConfigurationRecorderStatus2 struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsConfigConfigurationRecorderList is a list of AwsConfigConfigurationRecorders
type AwsConfigConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsConfigConfigurationRecorder CRD objects
	Items []AwsConfigConfigurationRecorder `json:"items,omitempty"`
}
