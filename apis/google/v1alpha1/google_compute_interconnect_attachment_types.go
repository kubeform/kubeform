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

type GoogleComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status            GoogleComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

type GoogleComputeInterconnectAttachmentSpecPrivateInterconnectInfo struct {
	Tag8021q int `json:"tag8021q"`
}

type GoogleComputeInterconnectAttachmentSpec struct {
	Region                  string                                    `json:"region"`
	CloudRouterIpAddress    string                                    `json:"cloud_router_ip_address"`
	Name                    string                                    `json:"name"`
	Router                  string                                    `json:"router"`
	Description             string                                    `json:"description"`
	CreationTimestamp       string                                    `json:"creation_timestamp"`
	CustomerRouterIpAddress string                                    `json:"customer_router_ip_address"`
	GoogleReferenceId       string                                    `json:"google_reference_id"`
	PrivateInterconnectInfo []GoogleComputeInterconnectAttachmentSpec `json:"private_interconnect_info"`
	Project                 string                                    `json:"project"`
	Interconnect            string                                    `json:"interconnect"`
	SelfLink                string                                    `json:"self_link"`
}

type GoogleComputeInterconnectAttachmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeInterconnectAttachmentList is a list of GoogleComputeInterconnectAttachments
type GoogleComputeInterconnectAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeInterconnectAttachment CRD objects
	Items []GoogleComputeInterconnectAttachment `json:"items,omitempty"`
}
