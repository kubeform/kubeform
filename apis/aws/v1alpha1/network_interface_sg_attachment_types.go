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

type NetworkInterfaceSgAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSgAttachmentSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceSgAttachmentStatus `json:"status,omitempty"`
}

type NetworkInterfaceSgAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkInterfaceID string `json:"networkInterfaceID" tf:"network_interface_id"`
	SecurityGroupID    string `json:"securityGroupID" tf:"security_group_id"`
}

type NetworkInterfaceSgAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkInterfaceSgAttachmentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceSgAttachmentList is a list of NetworkInterfaceSgAttachments
type NetworkInterfaceSgAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceSgAttachment CRD objects
	Items []NetworkInterfaceSgAttachment `json:"items,omitempty"`
}
