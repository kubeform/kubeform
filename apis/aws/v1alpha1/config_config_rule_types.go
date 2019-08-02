package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ConfigConfigRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigRuleSpec   `json:"spec,omitempty"`
	Status            ConfigConfigRuleStatus `json:"status,omitempty"`
}

type ConfigConfigRuleSpecScope struct {
	// +optional
	ComplianceResourceID string `json:"complianceResourceID,omitempty" tf:"compliance_resource_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:UniqueItems=true
	ComplianceResourceTypes []string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`
	// +optional
	TagKey string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`
	// +optional
	TagValue string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type ConfigConfigRuleSpecSourceSourceDetail struct {
	// +optional
	EventSource string `json:"eventSource,omitempty" tf:"event_source,omitempty"`
	// +optional
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`
	// +optional
	MessageType string `json:"messageType,omitempty" tf:"message_type,omitempty"`
}

type ConfigConfigRuleSpecSource struct {
	Owner string `json:"owner" tf:"owner"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	SourceDetail     []ConfigConfigRuleSpecSourceSourceDetail `json:"sourceDetail,omitempty" tf:"source_detail,omitempty"`
	SourceIdentifier string                                   `json:"sourceIdentifier" tf:"source_identifier"`
}

type ConfigConfigRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	InputParameters string `json:"inputParameters,omitempty" tf:"input_parameters,omitempty"`
	// +optional
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`
	Name                      string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scope []ConfigConfigRuleSpecScope `json:"scope,omitempty" tf:"scope,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Source []ConfigConfigRuleSpecSource `json:"source" tf:"source"`
}

type ConfigConfigRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *ConfigConfigRuleSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigRuleList is a list of ConfigConfigRules
type ConfigConfigRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigRule CRD objects
	Items []ConfigConfigRule `json:"items,omitempty"`
}
