package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsXraySamplingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsXraySamplingRuleSpec   `json:"spec,omitempty"`
	Status            AwsXraySamplingRuleStatus `json:"status,omitempty"`
}

type AwsXraySamplingRuleSpec struct {
	RuleName      string            `json:"rule_name"`
	ServiceName   string            `json:"service_name"`
	Host          string            `json:"host"`
	UrlPath       string            `json:"url_path"`
	Attributes    map[string]string `json:"attributes"`
	Arn           string            `json:"arn"`
	ResourceArn   string            `json:"resource_arn"`
	Priority      int               `json:"priority"`
	FixedRate     float64           `json:"fixed_rate"`
	ReservoirSize int               `json:"reservoir_size"`
	ServiceType   string            `json:"service_type"`
	HttpMethod    string            `json:"http_method"`
	Version       int               `json:"version"`
}



type AwsXraySamplingRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsXraySamplingRuleList is a list of AwsXraySamplingRules
type AwsXraySamplingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsXraySamplingRule CRD objects
	Items []AwsXraySamplingRule `json:"items,omitempty"`
}