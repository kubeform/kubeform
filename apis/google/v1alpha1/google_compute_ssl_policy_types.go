package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSslPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSslPolicySpec   `json:"spec,omitempty"`
	Status            GoogleComputeSslPolicyStatus `json:"status,omitempty"`
}

type GoogleComputeSslPolicySpec struct {
	MinTlsVersion     string   `json:"min_tls_version"`
	Fingerprint       string   `json:"fingerprint"`
	SelfLink          string   `json:"self_link"`
	Name              string   `json:"name"`
	CustomFeatures    []string `json:"custom_features"`
	CreationTimestamp string   `json:"creation_timestamp"`
	EnabledFeatures   []string `json:"enabled_features"`
	Project           string   `json:"project"`
	Description       string   `json:"description"`
	Profile           string   `json:"profile"`
}

type GoogleComputeSslPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSslPolicyList is a list of GoogleComputeSslPolicys
type GoogleComputeSslPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSslPolicy CRD objects
	Items []GoogleComputeSslPolicy `json:"items,omitempty"`
}
