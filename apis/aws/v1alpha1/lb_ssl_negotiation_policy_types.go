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

type LbSslNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSslNegotiationPolicySpec   `json:"spec,omitempty"`
	Status            LbSslNegotiationPolicyStatus `json:"status,omitempty"`
}

type LbSslNegotiationPolicySpecAttribute struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type LbSslNegotiationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Attribute    []LbSslNegotiationPolicySpecAttribute `json:"attribute,omitempty" tf:"attribute,omitempty"`
	LbPort       int                                   `json:"lbPort" tf:"lb_port"`
	LoadBalancer string                                `json:"loadBalancer" tf:"load_balancer"`
	Name         string                                `json:"name" tf:"name"`
}

type LbSslNegotiationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbSslNegotiationPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
