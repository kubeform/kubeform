package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ConfigConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationRecorderSpec   `json:"spec,omitempty"`
	Status            ConfigConfigurationRecorderStatus `json:"status,omitempty"`
}

type ConfigConfigurationRecorderSpecRecordingGroup struct {
	// +optional
	AllSupported bool `json:"allSupported,omitempty" tf:"all_supported,omitempty"`
	// +optional
	IncludeGlobalResourceTypes bool `json:"includeGlobalResourceTypes,omitempty" tf:"include_global_resource_types,omitempty"`
	// +optional
	ResourceTypes []string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type ConfigConfigurationRecorderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RecordingGroup []ConfigConfigurationRecorderSpecRecordingGroup `json:"recordingGroup,omitempty" tf:"recording_group,omitempty"`
	RoleArn        string                                          `json:"roleArn" tf:"role_arn"`
}

type ConfigConfigurationRecorderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigConfigurationRecorderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationRecorderList is a list of ConfigConfigurationRecorders
type ConfigConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationRecorder CRD objects
	Items []ConfigConfigurationRecorder `json:"items,omitempty"`
}
