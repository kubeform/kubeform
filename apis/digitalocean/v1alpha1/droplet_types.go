package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Droplet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DropletSpec   `json:"spec,omitempty"`
	Status            DropletStatus `json:"status,omitempty"`
}

type DropletSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Backups            bool   `json:"backups,omitempty" tf:"backups,omitempty"`
	Disk               int    `json:"disk" tf:"disk"`
	Image              string `json:"image" tf:"image"`
	Ipv4Address        string `json:"ipv4Address" tf:"ipv4_address"`
	Ipv4AddressPrivate string `json:"ipv4AddressPrivate" tf:"ipv4_address_private"`
	// +optional
	Ipv6        bool   `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
	Ipv6Address string `json:"ipv6Address" tf:"ipv6_address"`
	Locked      bool   `json:"locked" tf:"locked"`
	Memory      int    `json:"memory" tf:"memory"`
	// +optional
	Monitoring   bool        `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	Name         string      `json:"name" tf:"name"`
	PriceHourly  json.Number `json:"priceHourly" tf:"price_hourly"`
	PriceMonthly json.Number `json:"priceMonthly" tf:"price_monthly"`
	// +optional
	PrivateNetworking bool   `json:"privateNetworking,omitempty" tf:"private_networking,omitempty"`
	Region            string `json:"region" tf:"region"`
	// +optional
	ResizeDisk bool   `json:"resizeDisk,omitempty" tf:"resize_disk,omitempty"`
	Size       string `json:"size" tf:"size"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	Status  string   `json:"status" tf:"status"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	Urn  string   `json:"urn" tf:"urn"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	Vcpus    int    `json:"vcpus" tf:"vcpus"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VolumeIDS []string `json:"volumeIDS,omitempty" tf:"volume_ids,omitempty"`
}

type DropletStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DropletSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DropletList is a list of Droplets
type DropletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Droplet CRD objects
	Items []Droplet `json:"items,omitempty"`
}
