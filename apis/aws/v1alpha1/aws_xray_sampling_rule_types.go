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

type AwsXraySamplingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsXraySamplingRuleSpec   `json:"spec,omitempty"`
	Status            AwsXraySamplingRuleStatus `json:"status,omitempty"`
}

type AwsXraySamplingRuleSpec struct {
	Attributes    map[string]string `json:"attributes"`
	FixedRate     float64           `json:"fixed_rate"`
	ServiceName   string            `json:"service_name"`
	ServiceType   string            `json:"service_type"`
	Version       int               `json:"version"`
	Host          string            `json:"host"`
	HttpMethod    string            `json:"http_method"`
	UrlPath       string            `json:"url_path"`
	Arn           string            `json:"arn"`
	RuleName      string            `json:"rule_name"`
	ResourceArn   string            `json:"resource_arn"`
	Priority      int               `json:"priority"`
	ReservoirSize int               `json:"reservoir_size"`
}

type AwsXraySamplingRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsXraySamplingRuleList is a list of AwsXraySamplingRules
type AwsXraySamplingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsXraySamplingRule CRD objects
	Items []AwsXraySamplingRule `json:"items,omitempty"`
}
