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

type ConfigConfigRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigRuleSpec   `json:"spec,omitempty"`
	Status            ConfigConfigRuleStatus `json:"status,omitempty"`
}

type ConfigConfigRuleSpecScope struct {
	// +optional
	ComplianceResourceId string `json:"compliance_resource_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:UniqueItems=true
	ComplianceResourceTypes []string `json:"compliance_resource_types,omitempty"`
	// +optional
	TagKey string `json:"tag_key,omitempty"`
	// +optional
	TagValue string `json:"tag_value,omitempty"`
}

type ConfigConfigRuleSpecSourceSourceDetail struct {
	// +optional
	EventSource string `json:"event_source,omitempty"`
	// +optional
	MaximumExecutionFrequency string `json:"maximum_execution_frequency,omitempty"`
	// +optional
	MessageType string `json:"message_type,omitempty"`
}

type ConfigConfigRuleSpecSource struct {
	Owner string `json:"owner"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	SourceDetail     *[]ConfigConfigRuleSpecSource `json:"source_detail,omitempty"`
	SourceIdentifier string                        `json:"source_identifier"`
}

type ConfigConfigRuleSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	InputParameters string `json:"input_parameters,omitempty"`
	// +optional
	MaximumExecutionFrequency string `json:"maximum_execution_frequency,omitempty"`
	Name                      string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scope *[]ConfigConfigRuleSpec `json:"scope,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Source []ConfigConfigRuleSpec `json:"source"`
}

type ConfigConfigRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
