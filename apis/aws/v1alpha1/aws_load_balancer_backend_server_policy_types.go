package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerBackendServerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLoadBalancerBackendServerPolicySpec   `json:"spec,omitempty"`
	Status            AwsLoadBalancerBackendServerPolicyStatus `json:"status,omitempty"`
}

type AwsLoadBalancerBackendServerPolicySpec struct {
	LoadBalancerName string   `json:"load_balancer_name"`
	PolicyNames      []string `json:"policy_names"`
	InstancePort     int      `json:"instance_port"`
}

type AwsLoadBalancerBackendServerPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerBackendServerPolicyList is a list of AwsLoadBalancerBackendServerPolicys
type AwsLoadBalancerBackendServerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLoadBalancerBackendServerPolicy CRD objects
	Items []AwsLoadBalancerBackendServerPolicy `json:"items,omitempty"`
}
