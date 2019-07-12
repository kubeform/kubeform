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

type AwsSagemakerEndpointConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSagemakerEndpointConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsSagemakerEndpointConfigurationStatus `json:"status,omitempty"`
}

type AwsSagemakerEndpointConfigurationSpecProductionVariants struct {
	AcceleratorType      string  `json:"accelerator_type"`
	VariantName          string  `json:"variant_name"`
	ModelName            string  `json:"model_name"`
	InitialInstanceCount int     `json:"initial_instance_count"`
	InstanceType         string  `json:"instance_type"`
	InitialVariantWeight float64 `json:"initial_variant_weight"`
}

type AwsSagemakerEndpointConfigurationSpec struct {
	Tags               map[string]string                       `json:"tags"`
	Arn                string                                  `json:"arn"`
	Name               string                                  `json:"name"`
	ProductionVariants []AwsSagemakerEndpointConfigurationSpec `json:"production_variants"`
	KmsKeyArn          string                                  `json:"kms_key_arn"`
}

type AwsSagemakerEndpointConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSagemakerEndpointConfigurationList is a list of AwsSagemakerEndpointConfigurations
type AwsSagemakerEndpointConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSagemakerEndpointConfiguration CRD objects
	Items []AwsSagemakerEndpointConfiguration `json:"items,omitempty"`
}
