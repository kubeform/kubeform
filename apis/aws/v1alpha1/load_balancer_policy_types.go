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

type LoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerPolicySpec   `json:"spec,omitempty"`
	Status            LoadBalancerPolicyStatus `json:"status,omitempty"`
}

type LoadBalancerPolicySpecPolicyAttribute struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type LoadBalancerPolicySpec struct {
	LoadBalancerName string `json:"loadBalancerName" tf:"load_balancer_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PolicyAttribute []LoadBalancerPolicySpecPolicyAttribute `json:"policyAttribute,omitempty" tf:"policy_attribute,omitempty"`
	PolicyName      string                                  `json:"policyName" tf:"policy_name"`
	PolicyTypeName  string                                  `json:"policyTypeName" tf:"policy_type_name"`
	ProviderRef     core.LocalObjectReference               `json:"providerRef" tf:"-"`
}

type LoadBalancerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoadBalancerPolicyList is a list of LoadBalancerPolicys
type LoadBalancerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancerPolicy CRD objects
	Items []LoadBalancerPolicy `json:"items,omitempty"`
}
