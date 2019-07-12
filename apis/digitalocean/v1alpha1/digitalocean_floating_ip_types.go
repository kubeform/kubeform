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

type DigitaloceanFloatingIp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanFloatingIpSpec   `json:"spec,omitempty"`
	Status            DigitaloceanFloatingIpStatus `json:"status,omitempty"`
}

type DigitaloceanFloatingIpSpec struct {
	Region    string `json:"region"`
	Urn       string `json:"urn"`
	IpAddress string `json:"ip_address"`
	DropletId int    `json:"droplet_id"`
}

type DigitaloceanFloatingIpStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanFloatingIpList is a list of DigitaloceanFloatingIps
type DigitaloceanFloatingIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanFloatingIp CRD objects
	Items []DigitaloceanFloatingIp `json:"items,omitempty"`
}
