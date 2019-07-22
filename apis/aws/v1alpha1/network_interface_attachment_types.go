package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type NetworkInterfaceAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceAttachmentSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceAttachmentStatus `json:"status,omitempty"`
}

type NetworkInterfaceAttachmentSpec struct {
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
	DeviceIndex        int                       `json:"deviceIndex" tf:"device_index"`
	InstanceID         string                    `json:"instanceID" tf:"instance_id"`
	NetworkInterfaceID string                    `json:"networkInterfaceID" tf:"network_interface_id"`
}

type NetworkInterfaceAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceAttachmentList is a list of NetworkInterfaceAttachments
type NetworkInterfaceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceAttachment CRD objects
	Items []NetworkInterfaceAttachment `json:"items,omitempty"`
}
