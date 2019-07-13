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

type AwsRoute53ResolverRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53ResolverRuleSpec   `json:"spec,omitempty"`
	Status            AwsRoute53ResolverRuleStatus `json:"status,omitempty"`
}

type AwsRoute53ResolverRuleSpecTargetIp struct {
	Port int    `json:"port"`
	Ip   string `json:"ip"`
}

type AwsRoute53ResolverRuleSpec struct {
	Name               string                       `json:"name"`
	ResolverEndpointId string                       `json:"resolver_endpoint_id"`
	TargetIp           []AwsRoute53ResolverRuleSpec `json:"target_ip"`
	Arn                string                       `json:"arn"`
	DomainName         string                       `json:"domain_name"`
	RuleType           string                       `json:"rule_type"`
	Tags               map[string]string            `json:"tags"`
	OwnerId            string                       `json:"owner_id"`
	ShareStatus        string                       `json:"share_status"`
}



type AwsRoute53ResolverRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRoute53ResolverRuleList is a list of AwsRoute53ResolverRules
type AwsRoute53ResolverRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53ResolverRule CRD objects
	Items []AwsRoute53ResolverRule `json:"items,omitempty"`
}