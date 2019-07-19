package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	InstancePort     int    `json:"instancePort" tf:"instance_port"`
	LoadBalancerName string `json:"loadBalancerName" tf:"load_balancer_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PolicyNames []string                  `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LoadBalancerBackendServerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
