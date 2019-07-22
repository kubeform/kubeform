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

type LightsailStaticIPAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailStaticIPAttachmentSpec   `json:"spec,omitempty"`
	Status            LightsailStaticIPAttachmentStatus `json:"status,omitempty"`
}

type LightsailStaticIPAttachmentSpec struct {
	ProviderRef  core.LocalObjectReference `json:"providerRef" tf:"-"`
	InstanceName string                    `json:"instanceName" tf:"instance_name"`
	StaticIPName string                    `json:"staticIPName" tf:"static_ip_name"`
}

type LightsailStaticIPAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LightsailStaticIPAttachmentList is a list of LightsailStaticIPAttachments
type LightsailStaticIPAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LightsailStaticIPAttachment CRD objects
	Items []LightsailStaticIPAttachment `json:"items,omitempty"`
}
