package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LbSSLNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSSLNegotiationPolicySpec   `json:"spec,omitempty"`
	Status            LbSSLNegotiationPolicyStatus `json:"status,omitempty"`
}

type LbSSLNegotiationPolicySpecAttribute struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type LbSSLNegotiationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Attribute    []LbSSLNegotiationPolicySpecAttribute `json:"attribute,omitempty" tf:"attribute,omitempty"`
	LbPort       int                                   `json:"lbPort" tf:"lb_port"`
	LoadBalancer string                                `json:"loadBalancer" tf:"load_balancer"`
	Name         string                                `json:"name" tf:"name"`
}

type LbSSLNegotiationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbSSLNegotiationPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbSSLNegotiationPolicyList is a list of LbSSLNegotiationPolicys
type LbSSLNegotiationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbSSLNegotiationPolicy CRD objects
	Items []LbSSLNegotiationPolicy `json:"items,omitempty"`
}
