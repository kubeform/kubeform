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

type DxLag struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxLagSpec   `json:"spec,omitempty"`
	Status            DxLagStatus `json:"status,omitempty"`
}

type DxLagSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                  string `json:"arn,omitempty" tf:"arn,omitempty"`
	ConnectionsBandwidth string `json:"connectionsBandwidth" tf:"connections_bandwidth"`
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// +optional
	HasLogicalRedundancy string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy,omitempty"`
	// +optional
	JumboFrameCapable bool   `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DxLagStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxLagSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxLagList is a list of DxLags
type DxLagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxLag CRD objects
	Items []DxLag `json:"items,omitempty"`
}
