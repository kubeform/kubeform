package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafWebAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafWebAclSpec   `json:"spec,omitempty"`
	Status            AwsWafWebAclStatus `json:"status,omitempty"`
}

type AwsWafWebAclSpecLoggingConfigurationRedactedFieldsFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type AwsWafWebAclSpecLoggingConfigurationRedactedFields struct {
	FieldToMatch []AwsWafWebAclSpecLoggingConfigurationRedactedFields `json:"field_to_match"`
}

type AwsWafWebAclSpecLoggingConfiguration struct {
	LogDestination string                                 `json:"log_destination"`
	RedactedFields []AwsWafWebAclSpecLoggingConfiguration `json:"redacted_fields"`
}

type AwsWafWebAclSpecRulesAction struct {
	Type string `json:"type"`
}

type AwsWafWebAclSpecRulesOverrideAction struct {
	Type string `json:"type"`
}

type AwsWafWebAclSpecRules struct {
	Action         []AwsWafWebAclSpecRules `json:"action"`
	OverrideAction []AwsWafWebAclSpecRules `json:"override_action"`
	Priority       int                     `json:"priority"`
	Type           string                  `json:"type"`
	RuleId         string                  `json:"rule_id"`
}

type AwsWafWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type AwsWafWebAclSpec struct {
	MetricName           string             `json:"metric_name"`
	LoggingConfiguration []AwsWafWebAclSpec `json:"logging_configuration"`
	Rules                []AwsWafWebAclSpec `json:"rules"`
	Arn                  string             `json:"arn"`
	Name                 string             `json:"name"`
	DefaultAction        []AwsWafWebAclSpec `json:"default_action"`
}

type AwsWafWebAclStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAclList is a list of AwsWafWebAcls
type AwsWafWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafWebAcl CRD objects
	Items []AwsWafWebAcl `json:"items,omitempty"`
}
