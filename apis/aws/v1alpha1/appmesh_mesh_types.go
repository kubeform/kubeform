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
// +kubebuilder:subresource:status

type AppmeshMesh struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshMeshSpec   `json:"spec,omitempty"`
	Status            AppmeshMeshStatus `json:"status,omitempty"`
}

type AppmeshMeshSpecSpecEgressFilter struct {
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type AppmeshMeshSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EgressFilter []AppmeshMeshSpecSpecEgressFilter `json:"egressFilter,omitempty" tf:"egress_filter,omitempty"`
}

type AppmeshMeshSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	LastUpdatedDate string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Spec []AppmeshMeshSpecSpec `json:"spec,omitempty" tf:"spec,omitempty"`
}

type AppmeshMeshStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppmeshMeshSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppmeshMeshList is a list of AppmeshMeshs
type AppmeshMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppmeshMesh CRD objects
	Items []AppmeshMesh `json:"items,omitempty"`
}
