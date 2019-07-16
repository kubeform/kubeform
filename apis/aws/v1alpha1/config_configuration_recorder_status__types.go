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

type ConfigConfigurationRecorderStatus_ struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationRecorderStatus_Spec   `json:"spec,omitempty"`
	Status            ConfigConfigurationRecorderStatus_Status `json:"status,omitempty"`
}

type ConfigConfigurationRecorderStatus_Spec struct {
	IsEnabled bool   `json:"is_enabled"`
	Name      string `json:"name"`
}

type ConfigConfigurationRecorderStatus_Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationRecorderStatus_List is a list of ConfigConfigurationRecorderStatus_s
type ConfigConfigurationRecorderStatus_List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationRecorderStatus_ CRD objects
	Items []ConfigConfigurationRecorderStatus_ `json:"items,omitempty"`
}
