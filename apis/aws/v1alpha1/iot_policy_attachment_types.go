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

type IotPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            IotPolicyAttachmentStatus `json:"status,omitempty"`
}

type IotPolicyAttachmentSpec struct {
	Policy      string                    `json:"policy" tf:"policy"`
	Target      string                    `json:"target" tf:"target"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IotPolicyAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotPolicyAttachmentList is a list of IotPolicyAttachments
type IotPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotPolicyAttachment CRD objects
	Items []IotPolicyAttachment `json:"items,omitempty"`
}
