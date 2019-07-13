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

type AwsLightsailStaticIpAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLightsailStaticIpAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsLightsailStaticIpAttachmentStatus `json:"status,omitempty"`
}

type AwsLightsailStaticIpAttachmentSpec struct {
	StaticIpName string `json:"static_ip_name"`
	InstanceName string `json:"instance_name"`
}



type AwsLightsailStaticIpAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLightsailStaticIpAttachmentList is a list of AwsLightsailStaticIpAttachments
type AwsLightsailStaticIpAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLightsailStaticIpAttachment CRD objects
	Items []AwsLightsailStaticIpAttachment `json:"items,omitempty"`
}