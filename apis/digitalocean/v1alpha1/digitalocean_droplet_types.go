package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanDroplet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanDropletSpec   `json:"spec,omitempty"`
	Status            DigitaloceanDropletStatus `json:"status,omitempty"`
}

type DigitaloceanDropletSpec struct {
	Ipv6Address        string   `json:"ipv6_address"`
	VolumeIds          []string `json:"volume_ids"`
	Image              string   `json:"image"`
	Size               string   `json:"size"`
	Disk               int      `json:"disk"`
	Vcpus              int      `json:"vcpus"`
	Memory             int      `json:"memory"`
	Region             string   `json:"region"`
	PriceHourly        float64  `json:"price_hourly"`
	Status             string   `json:"status"`
	Tags               []string `json:"tags"`
	Ipv6               bool     `json:"ipv6"`
	Ipv6AddressPrivate string   `json:"ipv6_address_private"`
	Ipv4Address        string   `json:"ipv4_address"`
	Name               string   `json:"name"`
	PriceMonthly       float64  `json:"price_monthly"`
	ResizeDisk         bool     `json:"resize_disk"`
	Locked             bool     `json:"locked"`
	Backups            bool     `json:"backups"`
	Ipv4AddressPrivate string   `json:"ipv4_address_private"`
	SshKeys            []string `json:"ssh_keys"`
	UserData           string   `json:"user_data"`
	Urn                string   `json:"urn"`
	PrivateNetworking  bool     `json:"private_networking"`
	Monitoring         bool     `json:"monitoring"`
}

type DigitaloceanDropletStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanDropletList is a list of DigitaloceanDroplets
type DigitaloceanDropletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanDroplet CRD objects
	Items []DigitaloceanDroplet `json:"items,omitempty"`
}
