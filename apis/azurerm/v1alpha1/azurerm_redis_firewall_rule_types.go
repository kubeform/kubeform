package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRedisFirewallRuleSpec   `json:"spec,omitempty"`
	Status            AzurermRedisFirewallRuleStatus `json:"status,omitempty"`
}

type AzurermRedisFirewallRuleSpec struct {
	Name              string `json:"name"`
	RedisCacheName    string `json:"redis_cache_name"`
	ResourceGroupName string `json:"resource_group_name"`
	StartIp           string `json:"start_ip"`
	EndIp             string `json:"end_ip"`
}



type AzurermRedisFirewallRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRedisFirewallRuleList is a list of AzurermRedisFirewallRules
type AzurermRedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRedisFirewallRule CRD objects
	Items []AzurermRedisFirewallRule `json:"items,omitempty"`
}