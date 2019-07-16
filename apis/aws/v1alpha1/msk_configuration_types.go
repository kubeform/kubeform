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

type MskConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskConfigurationSpec   `json:"spec,omitempty"`
	Status            MskConfigurationStatus `json:"status,omitempty"`
}

type MskConfigurationSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	KafkaVersions    []string `json:"kafka_versions"`
	Name             string   `json:"name"`
	ServerProperties string   `json:"server_properties"`
}

type MskConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MskConfigurationList is a list of MskConfigurations
type MskConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MskConfiguration CRD objects
	Items []MskConfiguration `json:"items,omitempty"`
}
