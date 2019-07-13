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

type AwsNetworkInterfaceAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNetworkInterfaceAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsNetworkInterfaceAttachmentStatus `json:"status,omitempty"`
}

type AwsNetworkInterfaceAttachmentSpec struct {
	DeviceIndex        int    `json:"device_index"`
	InstanceId         string `json:"instance_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
	AttachmentId       string `json:"attachment_id"`
	Status             string `json:"status"`
}



type AwsNetworkInterfaceAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNetworkInterfaceAttachmentList is a list of AwsNetworkInterfaceAttachments
type AwsNetworkInterfaceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNetworkInterfaceAttachment CRD objects
	Items []AwsNetworkInterfaceAttachment `json:"items,omitempty"`
}