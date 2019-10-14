package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFirewallRuleSpec   `json:"spec,omitempty"`
	Status            RedisFirewallRuleStatus `json:"status,omitempty"`
}

type RedisFirewallRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	EndIP             string `json:"endIP" tf:"end_ip"`
	Name              string `json:"name" tf:"name"`
	RedisCacheName    string `json:"redisCacheName" tf:"redis_cache_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	StartIP           string `json:"startIP" tf:"start_ip"`
}

type RedisFirewallRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RedisFirewallRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisFirewallRuleList is a list of RedisFirewallRules
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedisFirewallRule CRD objects
	Items []RedisFirewallRule `json:"items,omitempty"`
}
