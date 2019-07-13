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

type AwsLbCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbCookieStickinessPolicySpec   `json:"spec,omitempty"`
	Status            AwsLbCookieStickinessPolicyStatus `json:"status,omitempty"`
}

type AwsLbCookieStickinessPolicySpec struct {
	Name                   string `json:"name"`
	LoadBalancer           string `json:"load_balancer"`
	LbPort                 int    `json:"lb_port"`
	CookieExpirationPeriod int    `json:"cookie_expiration_period"`
}



type AwsLbCookieStickinessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLbCookieStickinessPolicyList is a list of AwsLbCookieStickinessPolicys
type AwsLbCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbCookieStickinessPolicy CRD objects
	Items []AwsLbCookieStickinessPolicy `json:"items,omitempty"`
}