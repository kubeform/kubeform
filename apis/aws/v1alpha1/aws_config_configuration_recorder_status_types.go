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

type AwsConfigConfigurationRecorderStatus struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsConfigConfigurationRecorderStatusSpec   `json:"spec,omitempty"`
	Status            AwsConfigConfigurationRecorderStatusStatus `json:"status,omitempty"`
}

type AwsConfigConfigurationRecorderStatusSpec struct {
	Name      string `json:"name"`
	IsEnabled bool   `json:"is_enabled"`
}

type AwsConfigConfigurationRecorderStatusStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsConfigConfigurationRecorderStatusList is a list of AwsConfigConfigurationRecorderStatuss
type AwsConfigConfigurationRecorderStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsConfigConfigurationRecorderStatus CRD objects
	Items []AwsConfigConfigurationRecorderStatus `json:"items,omitempty"`
}
