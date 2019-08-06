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

type MskConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskConfigurationSpec   `json:"spec,omitempty"`
	Status            MskConfigurationStatus `json:"status,omitempty"`
}

type MskConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	KafkaVersions []string `json:"kafkaVersions" tf:"kafka_versions"`
	// +optional
	LatestRevision   int    `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`
	Name             string `json:"name" tf:"name"`
	ServerProperties string `json:"serverProperties" tf:"server_properties"`
}

type MskConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MskConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
