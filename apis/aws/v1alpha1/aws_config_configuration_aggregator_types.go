package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationAggregator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsConfigConfigurationAggregatorSpec   `json:"spec,omitempty"`
	Status            AwsConfigConfigurationAggregatorStatus `json:"status,omitempty"`
}

type AwsConfigConfigurationAggregatorSpecAccountAggregationSource struct {
	AccountIds []string `json:"account_ids"`
	AllRegions bool     `json:"all_regions"`
	Regions    []string `json:"regions"`
}

type AwsConfigConfigurationAggregatorSpecOrganizationAggregationSource struct {
	AllRegions bool     `json:"all_regions"`
	Regions    []string `json:"regions"`
	RoleArn    string   `json:"role_arn"`
}

type AwsConfigConfigurationAggregatorSpec struct {
	Arn                           string                                 `json:"arn"`
	Name                          string                                 `json:"name"`
	AccountAggregationSource      []AwsConfigConfigurationAggregatorSpec `json:"account_aggregation_source"`
	OrganizationAggregationSource []AwsConfigConfigurationAggregatorSpec `json:"organization_aggregation_source"`
}

type AwsConfigConfigurationAggregatorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigurationAggregatorList is a list of AwsConfigConfigurationAggregators
type AwsConfigConfigurationAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsConfigConfigurationAggregator CRD objects
	Items []AwsConfigConfigurationAggregator `json:"items,omitempty"`
}
