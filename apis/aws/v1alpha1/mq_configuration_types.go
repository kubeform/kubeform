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

type MqConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MqConfigurationSpec   `json:"spec,omitempty"`
	Status            MqConfigurationStatus `json:"status,omitempty"`
}

type MqConfigurationSpec struct {
	Data string `json:"data"`
	// +optional
	Description   string `json:"description,omitempty"`
	EngineType    string `json:"engine_type"`
	EngineVersion string `json:"engine_version"`
	Name          string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type MqConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MqConfigurationList is a list of MqConfigurations
type MqConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MqConfiguration CRD objects
	Items []MqConfiguration `json:"items,omitempty"`
}
