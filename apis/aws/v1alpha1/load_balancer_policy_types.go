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

type LoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerPolicySpec   `json:"spec,omitempty"`
	Status            LoadBalancerPolicyStatus `json:"status,omitempty"`
}

type LoadBalancerPolicySpecPolicyAttribute struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Value string `json:"value,omitempty"`
}

type LoadBalancerPolicySpec struct {
	LoadBalancerName string `json:"load_balancer_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PolicyAttribute *[]LoadBalancerPolicySpec `json:"policy_attribute,omitempty"`
	PolicyName      string                    `json:"policy_name"`
	PolicyTypeName  string                    `json:"policy_type_name"`
}

type LoadBalancerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
