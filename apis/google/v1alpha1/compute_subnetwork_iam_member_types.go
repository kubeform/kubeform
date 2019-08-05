package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeSubnetworkIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSubnetworkIamMemberSpec   `json:"spec,omitempty"`
	Status            ComputeSubnetworkIamMemberStatus `json:"status,omitempty"`
}

type ComputeSubnetworkIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Member string `json:"member" tf:"member"`
	// +optional
	// Deprecated
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// Deprecated
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	Role   string `json:"role" tf:"role"`
	// Deprecated
	Subnetwork string `json:"subnetwork" tf:"subnetwork"`
}

type ComputeSubnetworkIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeSubnetworkIamMemberSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkIamMemberList is a list of ComputeSubnetworkIamMembers
type ComputeSubnetworkIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSubnetworkIamMember CRD objects
	Items []ComputeSubnetworkIamMember `json:"items,omitempty"`
}
