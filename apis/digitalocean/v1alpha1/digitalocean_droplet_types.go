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

type DigitaloceanDroplet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanDropletSpec   `json:"spec,omitempty"`
	Status            DigitaloceanDropletStatus `json:"status,omitempty"`
}

type DigitaloceanDropletSpec struct {
	Status             string   `json:"status"`
	PrivateNetworking  bool     `json:"private_networking"`
	Ipv4AddressPrivate string   `json:"ipv4_address_private"`
	Tags               []string `json:"tags"`
	Name               string   `json:"name"`
	Size               string   `json:"size"`
	Disk               int      `json:"disk"`
	PriceHourly        float64  `json:"price_hourly"`
	Monitoring         bool     `json:"monitoring"`
	Region             string   `json:"region"`
	Urn                string   `json:"urn"`
	Vcpus              int      `json:"vcpus"`
	SshKeys            []string `json:"ssh_keys"`
	Ipv6Address        string   `json:"ipv6_address"`
	UserData           string   `json:"user_data"`
	VolumeIds          []string `json:"volume_ids"`
	Image              string   `json:"image"`
	PriceMonthly       float64  `json:"price_monthly"`
	ResizeDisk         bool     `json:"resize_disk"`
	Ipv6               bool     `json:"ipv6"`
	Ipv4Address        string   `json:"ipv4_address"`
	Memory             int      `json:"memory"`
	Locked             bool     `json:"locked"`
	Backups            bool     `json:"backups"`
	Ipv6AddressPrivate string   `json:"ipv6_address_private"`
}

type DigitaloceanDropletStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanDropletList is a list of DigitaloceanDroplets
type DigitaloceanDropletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanDroplet CRD objects
	Items []DigitaloceanDroplet `json:"items,omitempty"`
}
