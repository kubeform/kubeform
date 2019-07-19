package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
	// +optional
	Backups bool   `json:"backups,omitempty" tf:"backups,omitempty"`
	Image   string `json:"image" tf:"image"`
	// +optional
	Ipv6 bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
	// +optional
	Monitoring bool   `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	PrivateNetworking bool   `json:"privateNetworking,omitempty" tf:"private_networking,omitempty"`
	Region            string `json:"region" tf:"region"`
	// +optional
	ResizeDisk bool   `json:"resizeDisk,omitempty" tf:"resize_disk,omitempty"`
	Size       string `json:"size" tf:"size"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VolumeIDS   []string                  `json:"volumeIDS,omitempty" tf:"volume_ids,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DropletStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
