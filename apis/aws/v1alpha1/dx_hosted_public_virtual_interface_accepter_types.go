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
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DxHostedPublicVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPublicVirtualInterfaceAccepterSpec   `json:"spec,omitempty"`
	Status            DxHostedPublicVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

type DxHostedPublicVirtualInterfaceAccepterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VirtualInterfaceID string            `json:"virtualInterfaceID" tf:"virtual_interface_id"`
}

type DxHostedPublicVirtualInterfaceAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxHostedPublicVirtualInterfaceAccepterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterfaceAccepterList is a list of DxHostedPublicVirtualInterfaceAccepters
type DxHostedPublicVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxHostedPublicVirtualInterfaceAccepter CRD objects
	Items []DxHostedPublicVirtualInterfaceAccepter `json:"items,omitempty"`
}
