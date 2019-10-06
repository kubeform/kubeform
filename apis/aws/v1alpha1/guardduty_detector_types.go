package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GuarddutyDetector struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyDetectorSpec   `json:"spec,omitempty"`
	Status            GuarddutyDetectorStatus `json:"status,omitempty"`
}

type GuarddutyDetectorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID string `json:"accountID,omitempty" tf:"account_id,omitempty"`
	// +optional
	Enable bool `json:"enable,omitempty" tf:"enable,omitempty"`
	// +optional
	FindingPublishingFrequency string `json:"findingPublishingFrequency,omitempty" tf:"finding_publishing_frequency,omitempty"`
}

type GuarddutyDetectorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GuarddutyDetectorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyDetectorList is a list of GuarddutyDetectors
type GuarddutyDetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyDetector CRD objects
	Items []GuarddutyDetector `json:"items,omitempty"`
}
