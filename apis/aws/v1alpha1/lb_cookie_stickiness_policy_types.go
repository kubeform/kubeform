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

type LbCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbCookieStickinessPolicySpec   `json:"spec,omitempty"`
	Status            LbCookieStickinessPolicyStatus `json:"status,omitempty"`
}

type LbCookieStickinessPolicySpec struct {
	// +optional
	CookieExpirationPeriod int    `json:"cookie_expiration_period,omitempty"`
	LbPort                 int    `json:"lb_port"`
	LoadBalancer           string `json:"load_balancer"`
	Name                   string `json:"name"`
}

type LbCookieStickinessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
