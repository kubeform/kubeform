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

type LoadBalancerBackendServerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerBackendServerPolicySpec   `json:"spec,omitempty"`
	Status            LoadBalancerBackendServerPolicyStatus `json:"status,omitempty"`
}

type LoadBalancerBackendServerPolicySpec struct {
	InstancePort     int    `json:"instance_port"`
	LoadBalancerName string `json:"load_balancer_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PolicyNames []string `json:"policy_names,omitempty"`
}

type LoadBalancerBackendServerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoadBalancerBackendServerPolicyList is a list of LoadBalancerBackendServerPolicys
type LoadBalancerBackendServerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancerBackendServerPolicy CRD objects
	Items []LoadBalancerBackendServerPolicy `json:"items,omitempty"`
}
