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

type AwsConfigConfigRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsConfigConfigRuleSpec   `json:"spec,omitempty"`
	Status            AwsConfigConfigRuleStatus `json:"status,omitempty"`
}

type AwsConfigConfigRuleSpecScope struct {
	TagValue                string   `json:"tag_value"`
	ComplianceResourceId    string   `json:"compliance_resource_id"`
	ComplianceResourceTypes []string `json:"compliance_resource_types"`
	TagKey                  string   `json:"tag_key"`
}

type AwsConfigConfigRuleSpecSourceSourceDetail struct {
	EventSource               string `json:"event_source"`
	MaximumExecutionFrequency string `json:"maximum_execution_frequency"`
	MessageType               string `json:"message_type"`
}

type AwsConfigConfigRuleSpecSource struct {
	Owner            string                          `json:"owner"`
	SourceDetail     []AwsConfigConfigRuleSpecSource `json:"source_detail"`
	SourceIdentifier string                          `json:"source_identifier"`
}

type AwsConfigConfigRuleSpec struct {
	Arn                       string                    `json:"arn"`
	Description               string                    `json:"description"`
	InputParameters           string                    `json:"input_parameters"`
	MaximumExecutionFrequency string                    `json:"maximum_execution_frequency"`
	Scope                     []AwsConfigConfigRuleSpec `json:"scope"`
	Source                    []AwsConfigConfigRuleSpec `json:"source"`
	Name                      string                    `json:"name"`
	RuleId                    string                    `json:"rule_id"`
}



type AwsConfigConfigRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsConfigConfigRuleList is a list of AwsConfigConfigRules
type AwsConfigConfigRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsConfigConfigRule CRD objects
	Items []AwsConfigConfigRule `json:"items,omitempty"`
}