package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	AcceleratorType      string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`
	InitialInstanceCount int    `json:"initialInstanceCount" tf:"initial_instance_count"`
	// +optional
	InitialVariantWeight json.Number `json:"initialVariantWeight,omitempty" tf:"initial_variant_weight,omitempty"`
	InstanceType         string      `json:"instanceType" tf:"instance_type"`
	ModelName            string      `json:"modelName" tf:"model_name"`
	// +optional
	VariantName string `json:"variantName,omitempty" tf:"variant_name,omitempty"`
}

type SagemakerEndpointConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Name               string                                                 `json:"name,omitempty" tf:"name,omitempty"`
	ProductionVariants []SagemakerEndpointConfigurationSpecProductionVariants `json:"productionVariants" tf:"production_variants"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SagemakerEndpointConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SagemakerEndpointConfigurationSpec `json:"state,omitempty"`
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
