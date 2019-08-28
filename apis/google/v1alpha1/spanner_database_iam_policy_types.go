package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SpannerDatabaseIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerDatabaseIamPolicySpec   `json:"spec,omitempty"`
	Status            SpannerDatabaseIamPolicyStatus `json:"status,omitempty"`
}

type SpannerDatabaseIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Database string `json:"database" tf:"database"`
	// +optional
	Etag       string `json:"etag,omitempty" tf:"etag,omitempty"`
	Instance   string `json:"instance" tf:"instance"`
	PolicyData string `json:"policyData" tf:"policy_data"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}



type SpannerDatabaseIamPolicyStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *SpannerDatabaseIamPolicySpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerDatabaseIamPolicyList is a list of SpannerDatabaseIamPolicys
type SpannerDatabaseIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerDatabaseIamPolicy CRD objects
	Items []SpannerDatabaseIamPolicy `json:"items,omitempty"`
}