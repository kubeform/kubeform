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

type ComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status            ComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

type ComputeInterconnectAttachmentSpecPrivateInterconnectInfo struct {
	Tag8021q int `json:"tag8021q" tf:"tag8021q"`
}

type ComputeInterconnectAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	CloudRouterIPAddress    string `json:"cloudRouterIPAddress" tf:"cloud_router_ip_address"`
	CreationTimestamp       string `json:"creationTimestamp" tf:"creation_timestamp"`
	CustomerRouterIPAddress string `json:"customerRouterIPAddress" tf:"customer_router_ip_address"`
	// +optional
	Description       string `json:"description,omitempty" tf:"description,omitempty"`
	GoogleReferenceID string `json:"googleReferenceID" tf:"google_reference_id"`
	Interconnect      string `json:"interconnect" tf:"interconnect"`
	Name              string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	PrivateInterconnectInfo []ComputeInterconnectAttachmentSpecPrivateInterconnectInfo `json:"privateInterconnectInfo" tf:"private_interconnect_info"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region   string `json:"region,omitempty" tf:"region,omitempty"`
	Router   string `json:"router" tf:"router"`
	SelfLink string `json:"selfLink" tf:"self_link"`
}

type ComputeInterconnectAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeInterconnectAttachmentSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
