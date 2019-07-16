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

type WafregionalWebAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalWebAclSpec   `json:"spec,omitempty"`
	Status            WafregionalWebAclStatus `json:"status,omitempty"`
}

type WafregionalWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type WafregionalWebAclSpecLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

type WafregionalWebAclSpecLoggingConfigurationRedactedFields struct {
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch []WafregionalWebAclSpecLoggingConfigurationRedactedFields `json:"field_to_match"`
}

type WafregionalWebAclSpecLoggingConfiguration struct {
	LogDestination string `json:"log_destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedactedFields *[]WafregionalWebAclSpecLoggingConfiguration `json:"redacted_fields,omitempty"`
}

type WafregionalWebAclSpecRuleAction struct {
	Type string `json:"type"`
}

type WafregionalWebAclSpecRuleOverrideAction struct {
	Type string `json:"type"`
}

type WafregionalWebAclSpecRule struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Action *[]WafregionalWebAclSpecRule `json:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OverrideAction *[]WafregionalWebAclSpecRule `json:"override_action,omitempty"`
	Priority       int                          `json:"priority"`
	RuleId         string                       `json:"rule_id"`
	// +optional
	Type string `json:"type,omitempty"`
}

type WafregionalWebAclSpec struct {
	// +kubebuilder:validation:MaxItems=1
	DefaultAction []WafregionalWebAclSpec `json:"default_action"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfiguration *[]WafregionalWebAclSpec `json:"logging_configuration,omitempty"`
	MetricName           string                   `json:"metric_name"`
	Name                 string                   `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Rule *[]WafregionalWebAclSpec `json:"rule,omitempty"`
}

type WafregionalWebAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalWebAclList is a list of WafregionalWebAcls
type WafregionalWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalWebAcl CRD objects
	Items []WafregionalWebAcl `json:"items,omitempty"`
}
