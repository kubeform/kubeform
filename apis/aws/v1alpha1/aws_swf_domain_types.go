package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSwfDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSwfDomainSpec   `json:"spec,omitempty"`
	Status            AwsSwfDomainStatus `json:"status,omitempty"`
}

type AwsSwfDomainSpec struct {
	NamePrefix                             string `json:"name_prefix"`
	Description                            string `json:"description"`
	WorkflowExecutionRetentionPeriodInDays string `json:"workflow_execution_retention_period_in_days"`
	Name                                   string `json:"name"`
}

type AwsSwfDomainStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSwfDomainList is a list of AwsSwfDomains
type AwsSwfDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSwfDomain CRD objects
	Items []AwsSwfDomain `json:"items,omitempty"`
}
