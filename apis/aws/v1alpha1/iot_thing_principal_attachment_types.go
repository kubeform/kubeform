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

type IotThingPrincipalAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotThingPrincipalAttachmentSpec   `json:"spec,omitempty"`
	Status            IotThingPrincipalAttachmentStatus `json:"status,omitempty"`
}

type IotThingPrincipalAttachmentSpec struct {
	Principal   string                    `json:"principal" tf:"principal"`
	Thing       string                    `json:"thing" tf:"thing"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IotThingPrincipalAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotThingPrincipalAttachmentList is a list of IotThingPrincipalAttachments
type IotThingPrincipalAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotThingPrincipalAttachment CRD objects
	Items []IotThingPrincipalAttachment `json:"items,omitempty"`
}
