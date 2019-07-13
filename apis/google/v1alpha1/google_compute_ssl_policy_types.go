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

type GoogleComputeSslPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSslPolicySpec   `json:"spec,omitempty"`
	Status            GoogleComputeSslPolicyStatus `json:"status,omitempty"`
}

type GoogleComputeSslPolicySpec struct {
	MinTlsVersion     string   `json:"min_tls_version"`
	Profile           string   `json:"profile"`
	EnabledFeatures   []string `json:"enabled_features"`
	Project           string   `json:"project"`
	SelfLink          string   `json:"self_link"`
	Name              string   `json:"name"`
	CustomFeatures    []string `json:"custom_features"`
	Description       string   `json:"description"`
	CreationTimestamp string   `json:"creation_timestamp"`
	Fingerprint       string   `json:"fingerprint"`
}



type GoogleComputeSslPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeSslPolicyList is a list of GoogleComputeSslPolicys
type GoogleComputeSslPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSslPolicy CRD objects
	Items []GoogleComputeSslPolicy `json:"items,omitempty"`
}