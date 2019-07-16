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

type SesEventDestination struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesEventDestinationSpec   `json:"spec,omitempty"`
	Status            SesEventDestinationStatus `json:"status,omitempty"`
}

type SesEventDestinationSpecCloudwatchDestination struct {
	DefaultValue  string `json:"default_value"`
	DimensionName string `json:"dimension_name"`
	ValueSource   string `json:"value_source"`
}

type SesEventDestinationSpecKinesisDestination struct {
	RoleArn   string `json:"role_arn"`
	StreamArn string `json:"stream_arn"`
}

type SesEventDestinationSpecSnsDestination struct {
	TopicArn string `json:"topic_arn"`
}

type SesEventDestinationSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CloudwatchDestination *[]SesEventDestinationSpec `json:"cloudwatch_destination,omitempty"`
	ConfigurationSetName  string                     `json:"configuration_set_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	KinesisDestination *[]SesEventDestinationSpec `json:"kinesis_destination,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	MatchingTypes []string `json:"matching_types"`
	Name          string   `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	SnsDestination *[]SesEventDestinationSpec `json:"sns_destination,omitempty"`
}

type SesEventDestinationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
