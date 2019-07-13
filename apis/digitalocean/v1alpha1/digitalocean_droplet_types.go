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
	PriceMonthly       float64  `json:"price_monthly"`
	Ipv6AddressPrivate string   `json:"ipv6_address_private"`
	Monitoring         bool     `json:"monitoring"`
	Ipv4Address        string   `json:"ipv4_address"`
	Name               string   `json:"name"`
	Region             string   `json:"region"`
	Size               string   `json:"size"`
	Disk               int      `json:"disk"`
	PriceHourly        float64  `json:"price_hourly"`
	Backups            bool     `json:"backups"`
	PrivateNetworking  bool     `json:"private_networking"`
	Image              string   `json:"image"`
	Status             string   `json:"status"`
	Locked             bool     `json:"locked"`
	Ipv6               bool     `json:"ipv6"`
	SshKeys            []string `json:"ssh_keys"`
	VolumeIds          []string `json:"volume_ids"`
	Urn                string   `json:"urn"`
	Vcpus              int      `json:"vcpus"`
	Memory             int      `json:"memory"`
	ResizeDisk         bool     `json:"resize_disk"`
	Ipv6Address        string   `json:"ipv6_address"`
	Ipv4AddressPrivate string   `json:"ipv4_address_private"`
	UserData           string   `json:"user_data"`
	Tags               []string `json:"tags"`
}



type DigitaloceanDropletStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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