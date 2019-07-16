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

type ProxyProtocolPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxyProtocolPolicySpec   `json:"spec,omitempty"`
	Status            ProxyProtocolPolicyStatus `json:"status,omitempty"`
}

type ProxyProtocolPolicySpec struct {
	// +kubebuilder:validation:UniqueItems=true
	InstancePorts []string `json:"instance_ports"`
	LoadBalancer  string   `json:"load_balancer"`
}

type ProxyProtocolPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProxyProtocolPolicyList is a list of ProxyProtocolPolicys
type ProxyProtocolPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProxyProtocolPolicy CRD objects
	Items []ProxyProtocolPolicy `json:"items,omitempty"`
}
