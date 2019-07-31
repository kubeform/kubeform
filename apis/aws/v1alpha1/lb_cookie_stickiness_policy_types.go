package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LbCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbCookieStickinessPolicySpec   `json:"spec,omitempty"`
	Status            LbCookieStickinessPolicyStatus `json:"status,omitempty"`
}

type LbCookieStickinessPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	CookieExpirationPeriod int    `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period,omitempty"`
	LbPort                 int    `json:"lbPort" tf:"lb_port"`
	LoadBalancer           string `json:"loadBalancer" tf:"load_balancer"`
	Name                   string `json:"name" tf:"name"`
}

type LbCookieStickinessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbCookieStickinessPolicyList is a list of LbCookieStickinessPolicys
type LbCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbCookieStickinessPolicy CRD objects
	Items []LbCookieStickinessPolicy `json:"items,omitempty"`
}
