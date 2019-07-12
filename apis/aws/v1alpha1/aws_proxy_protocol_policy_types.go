package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsProxyProtocolPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsProxyProtocolPolicySpec   `json:"spec,omitempty"`
	Status            AwsProxyProtocolPolicyStatus `json:"status,omitempty"`
}

type AwsProxyProtocolPolicySpec struct {
	LoadBalancer  string   `json:"load_balancer"`
	InstancePorts []string `json:"instance_ports"`
}

type AwsProxyProtocolPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsProxyProtocolPolicyList is a list of AwsProxyProtocolPolicys
type AwsProxyProtocolPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsProxyProtocolPolicy CRD objects
	Items []AwsProxyProtocolPolicy `json:"items,omitempty"`
}
