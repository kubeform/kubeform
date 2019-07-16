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

type LbSslNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSslNegotiationPolicySpec   `json:"spec,omitempty"`
	Status            LbSslNegotiationPolicyStatus `json:"status,omitempty"`
}

type LbSslNegotiationPolicySpecAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LbSslNegotiationPolicySpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Attribute    *[]LbSslNegotiationPolicySpec `json:"attribute,omitempty"`
	LbPort       int                           `json:"lb_port"`
	LoadBalancer string                        `json:"load_balancer"`
	Name         string                        `json:"name"`
}

type LbSslNegotiationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbSslNegotiationPolicyList is a list of LbSslNegotiationPolicys
type LbSslNegotiationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbSslNegotiationPolicy CRD objects
	Items []LbSslNegotiationPolicy `json:"items,omitempty"`
}
