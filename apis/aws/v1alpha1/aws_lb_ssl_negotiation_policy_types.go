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

type AwsLbSslNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbSslNegotiationPolicySpec   `json:"spec,omitempty"`
	Status            AwsLbSslNegotiationPolicyStatus `json:"status,omitempty"`
}

type AwsLbSslNegotiationPolicySpecAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsLbSslNegotiationPolicySpec struct {
	Name         string                          `json:"name"`
	LoadBalancer string                          `json:"load_balancer"`
	LbPort       int                             `json:"lb_port"`
	Attribute    []AwsLbSslNegotiationPolicySpec `json:"attribute"`
}



type AwsLbSslNegotiationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLbSslNegotiationPolicyList is a list of AwsLbSslNegotiationPolicys
type AwsLbSslNegotiationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbSslNegotiationPolicy CRD objects
	Items []AwsLbSslNegotiationPolicy `json:"items,omitempty"`
}