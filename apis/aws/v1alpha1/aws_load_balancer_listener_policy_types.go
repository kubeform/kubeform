package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerListenerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLoadBalancerListenerPolicySpec   `json:"spec,omitempty"`
	Status            AwsLoadBalancerListenerPolicyStatus `json:"status,omitempty"`
}

type AwsLoadBalancerListenerPolicySpec struct {
	LoadBalancerPort int      `json:"load_balancer_port"`
	LoadBalancerName string   `json:"load_balancer_name"`
	PolicyNames      []string `json:"policy_names"`
}



type AwsLoadBalancerListenerPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerListenerPolicyList is a list of AwsLoadBalancerListenerPolicys
type AwsLoadBalancerListenerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLoadBalancerListenerPolicy CRD objects
	Items []AwsLoadBalancerListenerPolicy `json:"items,omitempty"`
}