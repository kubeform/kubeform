package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AppCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCookieStickinessPolicySpec   `json:"spec,omitempty"`
	Status            AppCookieStickinessPolicyStatus `json:"status,omitempty"`
}

type AppCookieStickinessPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CookieName   string `json:"cookieName" tf:"cookie_name"`
	LbPort       int    `json:"lbPort" tf:"lb_port"`
	LoadBalancer string `json:"loadBalancer" tf:"load_balancer"`
	Name         string `json:"name" tf:"name"`
}

type AppCookieStickinessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppCookieStickinessPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppCookieStickinessPolicyList is a list of AppCookieStickinessPolicys
type AppCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppCookieStickinessPolicy CRD objects
	Items []AppCookieStickinessPolicy `json:"items,omitempty"`
}
