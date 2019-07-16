package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SagemakerEndpointConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerEndpointConfigurationSpec   `json:"spec,omitempty"`
	Status            SagemakerEndpointConfigurationStatus `json:"status,omitempty"`
}

type SagemakerEndpointConfigurationSpecProductionVariants struct {
	// +optional
	AcceleratorType      string `json:"accelerator_type,omitempty"`
	InitialInstanceCount int    `json:"initial_instance_count"`
	// +optional
	InitialVariantWeight json.Number `json:"initial_variant_weight,omitempty"`
	InstanceType         string      `json:"instance_type"`
	ModelName            string      `json:"model_name"`
}

type SagemakerEndpointConfigurationSpec struct {
	// +optional
	KmsKeyArn          string                               `json:"kms_key_arn,omitempty"`
	ProductionVariants []SagemakerEndpointConfigurationSpec `json:"production_variants"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SagemakerEndpointConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerEndpointConfigurationList is a list of SagemakerEndpointConfigurations
type SagemakerEndpointConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerEndpointConfiguration CRD objects
	Items []SagemakerEndpointConfiguration `json:"items,omitempty"`
}
