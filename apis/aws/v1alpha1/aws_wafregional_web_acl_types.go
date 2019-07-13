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

type AwsWafregionalWebAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalWebAclSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalWebAclStatus `json:"status,omitempty"`
}

type AwsWafregionalWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type AwsWafregionalWebAclSpecLoggingConfigurationRedactedFieldsFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafregionalWebAclSpecLoggingConfigurationRedactedFields struct {
	FieldToMatch []AwsWafregionalWebAclSpecLoggingConfigurationRedactedFields `json:"field_to_match"`
}

type AwsWafregionalWebAclSpecLoggingConfiguration struct {
	LogDestination string                                         `json:"log_destination"`
	RedactedFields []AwsWafregionalWebAclSpecLoggingConfiguration `json:"redacted_fields"`
}

type AwsWafregionalWebAclSpecRuleAction struct {
	Type string `json:"type"`
}

type AwsWafregionalWebAclSpecRuleOverrideAction struct {
	Type string `json:"type"`
}

type AwsWafregionalWebAclSpecRule struct {
	Action         []AwsWafregionalWebAclSpecRule `json:"action"`
	OverrideAction []AwsWafregionalWebAclSpecRule `json:"override_action"`
	Priority       int                            `json:"priority"`
	Type           string                         `json:"type"`
	RuleId         string                         `json:"rule_id"`
}

type AwsWafregionalWebAclSpec struct {
	Arn                  string                     `json:"arn"`
	Name                 string                     `json:"name"`
	DefaultAction        []AwsWafregionalWebAclSpec `json:"default_action"`
	LoggingConfiguration []AwsWafregionalWebAclSpec `json:"logging_configuration"`
	MetricName           string                     `json:"metric_name"`
	Rule                 []AwsWafregionalWebAclSpec `json:"rule"`
}



type AwsWafregionalWebAclStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalWebAclList is a list of AwsWafregionalWebAcls
type AwsWafregionalWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalWebAcl CRD objects
	Items []AwsWafregionalWebAcl `json:"items,omitempty"`
}