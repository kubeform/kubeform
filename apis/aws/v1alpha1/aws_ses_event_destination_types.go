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

type AwsSesEventDestination struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesEventDestinationSpec   `json:"spec,omitempty"`
	Status            AwsSesEventDestinationStatus `json:"status,omitempty"`
}

type AwsSesEventDestinationSpecCloudwatchDestination struct {
	DefaultValue  string `json:"default_value"`
	DimensionName string `json:"dimension_name"`
	ValueSource   string `json:"value_source"`
}

type AwsSesEventDestinationSpecKinesisDestination struct {
	StreamArn string `json:"stream_arn"`
	RoleArn   string `json:"role_arn"`
}

type AwsSesEventDestinationSpecSnsDestination struct {
	TopicArn string `json:"topic_arn"`
}

type AwsSesEventDestinationSpec struct {
	CloudwatchDestination []AwsSesEventDestinationSpec `json:"cloudwatch_destination"`
	KinesisDestination    []AwsSesEventDestinationSpec `json:"kinesis_destination"`
	SnsDestination        []AwsSesEventDestinationSpec `json:"sns_destination"`
	Name                  string                       `json:"name"`
	ConfigurationSetName  string                       `json:"configuration_set_name"`
	Enabled               bool                         `json:"enabled"`
	MatchingTypes         []string                     `json:"matching_types"`
}



type AwsSesEventDestinationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesEventDestinationList is a list of AwsSesEventDestinations
type AwsSesEventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesEventDestination CRD objects
	Items []AwsSesEventDestination `json:"items,omitempty"`
}