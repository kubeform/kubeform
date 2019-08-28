package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LoadBalancerListenerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerListenerPolicySpec   `json:"spec,omitempty"`
	Status            LoadBalancerListenerPolicyStatus `json:"status,omitempty"`
}

type LoadBalancerListenerPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	LoadBalancerName string `json:"loadBalancerName" tf:"load_balancer_name"`
	LoadBalancerPort int    `json:"loadBalancerPort" tf:"load_balancer_port"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PolicyNames []string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}



type LoadBalancerListenerPolicyStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *LoadBalancerListenerPolicySpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoadBalancerListenerPolicyList is a list of LoadBalancerListenerPolicys
type LoadBalancerListenerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancerListenerPolicy CRD objects
	Items []LoadBalancerListenerPolicy `json:"items,omitempty"`
}