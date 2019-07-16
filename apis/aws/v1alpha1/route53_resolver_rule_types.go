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

type Route53ResolverRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverRuleSpec   `json:"spec,omitempty"`
	Status            Route53ResolverRuleStatus `json:"status,omitempty"`
}

type Route53ResolverRuleSpecTargetIp struct {
	Ip string `json:"ip"`
	// +optional
	Port int `json:"port,omitempty"`
}

type Route53ResolverRuleSpec struct {
	DomainName string `json:"domain_name"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	ResolverEndpointId string `json:"resolver_endpoint_id,omitempty"`
	RuleType           string `json:"rule_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetIp *[]Route53ResolverRuleSpec `json:"target_ip,omitempty"`
}

type Route53ResolverRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ResolverRuleList is a list of Route53ResolverRules
type Route53ResolverRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53ResolverRule CRD objects
	Items []Route53ResolverRule `json:"items,omitempty"`
}
