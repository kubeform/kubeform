package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLoadBalancerPolicySpec   `json:"spec,omitempty"`
	Status            AwsLoadBalancerPolicyStatus `json:"status,omitempty"`
}

type AwsLoadBalancerPolicySpecPolicyAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsLoadBalancerPolicySpec struct {
	PolicyTypeName   string                      `json:"policy_type_name"`
	PolicyAttribute  []AwsLoadBalancerPolicySpec `json:"policy_attribute"`
	LoadBalancerName string                      `json:"load_balancer_name"`
	PolicyName       string                      `json:"policy_name"`
}



type AwsLoadBalancerPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerPolicyList is a list of AwsLoadBalancerPolicys
type AwsLoadBalancerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLoadBalancerPolicy CRD objects
	Items []AwsLoadBalancerPolicy `json:"items,omitempty"`
}