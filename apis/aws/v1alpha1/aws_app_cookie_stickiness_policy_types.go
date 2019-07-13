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

type AwsAppCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppCookieStickinessPolicySpec   `json:"spec,omitempty"`
	Status            AwsAppCookieStickinessPolicyStatus `json:"status,omitempty"`
}

type AwsAppCookieStickinessPolicySpec struct {
	Name         string `json:"name"`
	LoadBalancer string `json:"load_balancer"`
	LbPort       int    `json:"lb_port"`
	CookieName   string `json:"cookie_name"`
}



type AwsAppCookieStickinessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppCookieStickinessPolicyList is a list of AwsAppCookieStickinessPolicys
type AwsAppCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppCookieStickinessPolicy CRD objects
	Items []AwsAppCookieStickinessPolicy `json:"items,omitempty"`
}