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

type AwsLoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLoadBalancerPolicySpec   `json:"spec,omitempty"`
	Status            AwsLoadBalancerPolicyStatus `json:"status,omitempty"`
}

type AwsLoadBalancerPolicySpecPolicyAttribute struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type AwsLoadBalancerPolicySpec struct {
	PolicyName       string                      `json:"policy_name"`
	PolicyTypeName   string                      `json:"policy_type_name"`
	PolicyAttribute  []AwsLoadBalancerPolicySpec `json:"policy_attribute"`
	LoadBalancerName string                      `json:"load_balancer_name"`
}



type AwsLoadBalancerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLoadBalancerPolicyList is a list of AwsLoadBalancerPolicys
type AwsLoadBalancerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLoadBalancerPolicy CRD objects
	Items []AwsLoadBalancerPolicy `json:"items,omitempty"`
}