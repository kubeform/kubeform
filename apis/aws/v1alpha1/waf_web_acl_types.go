package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type WafWebACL struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafWebACLSpec   `json:"spec,omitempty"`
	Status            WafWebACLStatus `json:"status,omitempty"`
}

type WafWebACLSpecDefaultAction struct {
	Type string `json:"type" tf:"type"`
}

type WafWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafWebACLSpecLoggingConfigurationRedactedFields struct {
	FieldToMatch []WafWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
}

type WafWebACLSpecLoggingConfiguration struct {
	LogDestination string `json:"logDestination" tf:"log_destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedactedFields []WafWebACLSpecLoggingConfigurationRedactedFields `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`
}

type WafWebACLSpecRulesAction struct {
	Type string `json:"type" tf:"type"`
}

type WafWebACLSpecRulesOverrideAction struct {
	Type string `json:"type" tf:"type"`
}

type WafWebACLSpecRules struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Action []WafWebACLSpecRulesAction `json:"action,omitempty" tf:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OverrideAction []WafWebACLSpecRulesOverrideAction `json:"overrideAction,omitempty" tf:"override_action,omitempty"`
	Priority       int                                `json:"priority" tf:"priority"`
	RuleID         string                             `json:"ruleID" tf:"rule_id"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type WafWebACLSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	DefaultAction []WafWebACLSpecDefaultAction `json:"defaultAction" tf:"default_action"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfiguration []WafWebACLSpecLoggingConfiguration `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`
	MetricName           string                              `json:"metricName" tf:"metric_name"`
	Name                 string                              `json:"name" tf:"name"`
	// +optional
	Rules []WafWebACLSpecRules `json:"rules,omitempty" tf:"rules,omitempty"`
}

type WafWebACLStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafWebACLSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafWebACLList is a list of WafWebACLs
type WafWebACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafWebACL CRD objects
	Items []WafWebACL `json:"items,omitempty"`
}
