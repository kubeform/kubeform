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

type GoogleLoggingBillingAccountExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingBillingAccountExclusionSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingBillingAccountExclusionStatus `json:"status,omitempty"`
}

type GoogleLoggingBillingAccountExclusionSpec struct {
	Name           string `json:"name"`
	BillingAccount string `json:"billing_account"`
	Description    string `json:"description"`
	Disabled       bool   `json:"disabled"`
	Filter         string `json:"filter"`
}



type GoogleLoggingBillingAccountExclusionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingBillingAccountExclusionList is a list of GoogleLoggingBillingAccountExclusions
type GoogleLoggingBillingAccountExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingBillingAccountExclusion CRD objects
	Items []GoogleLoggingBillingAccountExclusion `json:"items,omitempty"`
}