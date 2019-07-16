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

type WafWebAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafWebAclSpec   `json:"spec,omitempty"`
	Status            WafWebAclStatus `json:"status,omitempty"`
}

type WafWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type WafWebAclSpecLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafWebAclSpecLoggingConfigurationRedactedFields struct {
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch []WafWebAclSpecLoggingConfigurationRedactedFields `json:"field_to_match"`
}

type WafWebAclSpecLoggingConfiguration struct {
	LogDestination string `json:"log_destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedactedFields *[]WafWebAclSpecLoggingConfiguration `json:"redacted_fields,omitempty"`
}

type WafWebAclSpecRulesAction struct {
	Type string `json:"type"`
}

type WafWebAclSpecRulesOverrideAction struct {
	Type string `json:"type"`
}

type WafWebAclSpecRules struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Action *[]WafWebAclSpecRules `json:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OverrideAction *[]WafWebAclSpecRules `json:"override_action,omitempty"`
	Priority       int                   `json:"priority"`
	RuleId         string                `json:"rule_id"`
	// +optional
	Type string `json:"type,omitempty"`
}

type WafWebAclSpec struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	DefaultAction []WafWebAclSpec `json:"default_action"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfiguration *[]WafWebAclSpec `json:"logging_configuration,omitempty"`
	MetricName           string           `json:"metric_name"`
	Name                 string           `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Rules *[]WafWebAclSpec `json:"rules,omitempty"`
}

type WafWebAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafWebAclList is a list of WafWebAcls
type WafWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafWebAcl CRD objects
	Items []WafWebAcl `json:"items,omitempty"`
}
