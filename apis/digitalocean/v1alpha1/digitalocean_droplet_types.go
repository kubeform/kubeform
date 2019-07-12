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
	Disk               int      `json:"disk"`
	Vcpus              int      `json:"vcpus"`
	ResizeDisk         bool     `json:"resize_disk"`
	Status             string   `json:"status"`
	Ipv6               bool     `json:"ipv6"`
	SshKeys            []string `json:"ssh_keys"`
	Image              string   `json:"image"`
	Region             string   `json:"region"`
	Urn                string   `json:"urn"`
	PriceHourly        float64  `json:"price_hourly"`
	PriceMonthly       float64  `json:"price_monthly"`
	Backups            bool     `json:"backups"`
	PrivateNetworking  bool     `json:"private_networking"`
	UserData           string   `json:"user_data"`
	Locked             bool     `json:"locked"`
	Ipv6AddressPrivate string   `json:"ipv6_address_private"`
	Ipv4AddressPrivate string   `json:"ipv4_address_private"`
	Tags               []string `json:"tags"`
	Name               string   `json:"name"`
	Size               string   `json:"size"`
	Memory             int      `json:"memory"`
	Ipv6Address        string   `json:"ipv6_address"`
	Ipv4Address        string   `json:"ipv4_address"`
	VolumeIds          []string `json:"volume_ids"`
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