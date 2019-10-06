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

type ComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status            ComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

type ComputeInterconnectAttachmentSpecPrivateInterconnectInfo struct {
	// +optional
	Tag8021q int `json:"tag8021q,omitempty" tf:"tag8021q,omitempty"`
}

type ComputeInterconnectAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CloudRouterIPAddress string `json:"cloudRouterIPAddress,omitempty" tf:"cloud_router_ip_address,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	CustomerRouterIPAddress string `json:"customerRouterIPAddress,omitempty" tf:"customer_router_ip_address,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	GoogleReferenceID string `json:"googleReferenceID,omitempty" tf:"google_reference_id,omitempty"`
	Interconnect      string `json:"interconnect" tf:"interconnect"`
	Name              string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PrivateInterconnectInfo []ComputeInterconnectAttachmentSpecPrivateInterconnectInfo `json:"privateInterconnectInfo,omitempty" tf:"private_interconnect_info,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	Router string `json:"router" tf:"router"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeInterconnectAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeInterconnectAttachmentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInterconnectAttachmentList is a list of ComputeInterconnectAttachments
type ComputeInterconnectAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInterconnectAttachment CRD objects
	Items []ComputeInterconnectAttachment `json:"items,omitempty"`
}
