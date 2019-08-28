package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn  string `json:"arn,omitempty" tf:"arn,omitempty"`
	Data string `json:"data" tf:"data"`
	// +optional
	Description   string `json:"description,omitempty" tf:"description,omitempty"`
	EngineType    string `json:"engineType" tf:"engine_type"`
	EngineVersion string `json:"engineVersion" tf:"engine_version"`
	// +optional
	LatestRevision int    `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`
	Name           string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type MqConfigurationStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *MqConfigurationSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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