package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SesEventDestination struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesEventDestinationSpec   `json:"spec,omitempty"`
	Status            SesEventDestinationStatus `json:"status,omitempty"`
}

type SesEventDestinationSpecCloudwatchDestination struct {
	DefaultValue  string `json:"defaultValue" tf:"default_value"`
	DimensionName string `json:"dimensionName" tf:"dimension_name"`
	ValueSource   string `json:"valueSource" tf:"value_source"`
}

type SesEventDestinationSpecKinesisDestination struct {
	RoleArn   string `json:"roleArn" tf:"role_arn"`
	StreamArn string `json:"streamArn" tf:"stream_arn"`
}

type SesEventDestinationSpecSnsDestination struct {
	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

type SesEventDestinationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchDestination []SesEventDestinationSpecCloudwatchDestination `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination,omitempty"`
	ConfigurationSetName  string                                         `json:"configurationSetName" tf:"configuration_set_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	KinesisDestination []SesEventDestinationSpecKinesisDestination `json:"kinesisDestination,omitempty" tf:"kinesis_destination,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	MatchingTypes []string `json:"matchingTypes" tf:"matching_types"`
	Name          string   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	SnsDestination []SesEventDestinationSpecSnsDestination `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type SesEventDestinationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SesEventDestinationSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesEventDestinationList is a list of SesEventDestinations
type SesEventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesEventDestination CRD objects
	Items []SesEventDestination `json:"items,omitempty"`
}
