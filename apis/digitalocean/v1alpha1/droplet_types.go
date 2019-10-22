package v1alpha1

import (
	"encoding/json"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Droplet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DropletSpec   `json:"spec,omitempty"`
	Status            DropletStatus `json:"status,omitempty"`
}

type DropletSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Backups bool `json:"backups,omitempty" tf:"backups,omitempty"`
	// +optional
	Disk  int    `json:"disk,omitempty" tf:"disk,omitempty"`
	Image string `json:"image" tf:"image"`
	// +optional
	Ipv4Address string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`
	// +optional
	Ipv4AddressPrivate string `json:"ipv4AddressPrivate,omitempty" tf:"ipv4_address_private,omitempty"`
	// +optional
	Ipv6 bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
	// +optional
	Ipv6Address string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`
	// +optional
	Locked bool `json:"locked,omitempty" tf:"locked,omitempty"`
	// +optional
	Memory int `json:"memory,omitempty" tf:"memory,omitempty"`
	// +optional
	Monitoring bool   `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	PriceHourly json.Number `json:"priceHourly,omitempty" tf:"price_hourly,omitempty"`
	// +optional
	PriceMonthly json.Number `json:"priceMonthly,omitempty" tf:"price_monthly,omitempty"`
	// +optional
	PrivateNetworking bool   `json:"privateNetworking,omitempty" tf:"private_networking,omitempty"`
	Region            string `json:"region" tf:"region"`
	// +optional
	ResizeDisk bool   `json:"resizeDisk,omitempty" tf:"resize_disk,omitempty"`
	Size       string `json:"size" tf:"size"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Urn string `json:"urn,omitempty" tf:"urn,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	Vcpus int `json:"vcpus,omitempty" tf:"vcpus,omitempty"`
	// +optional
	VolumeIDS []string `json:"volumeIDS,omitempty" tf:"volume_ids,omitempty"`
}

type DropletStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DropletSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
