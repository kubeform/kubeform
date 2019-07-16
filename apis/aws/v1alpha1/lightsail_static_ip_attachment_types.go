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

type LightsailStaticIpAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailStaticIpAttachmentSpec   `json:"spec,omitempty"`
	Status            LightsailStaticIpAttachmentStatus `json:"status,omitempty"`
}

type LightsailStaticIpAttachmentSpec struct {
	InstanceName string `json:"instance_name"`
	StaticIpName string `json:"static_ip_name"`
}

type LightsailStaticIpAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LightsailStaticIpAttachmentList is a list of LightsailStaticIpAttachments
type LightsailStaticIpAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LightsailStaticIpAttachment CRD objects
	Items []LightsailStaticIpAttachment `json:"items,omitempty"`
}
