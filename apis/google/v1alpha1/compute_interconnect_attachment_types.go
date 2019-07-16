package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status            ComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

type ComputeInterconnectAttachmentSpec struct {
	// +optional
	Description  string `json:"description,omitempty"`
	Interconnect string `json:"interconnect"`
	Name         string `json:"name"`
	Router       string `json:"router"`
}

type ComputeInterconnectAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
