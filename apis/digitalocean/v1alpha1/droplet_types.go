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

type Droplet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DropletSpec   `json:"spec,omitempty"`
	Status            DropletStatus `json:"status,omitempty"`
}

type DropletSpec struct {
	// +optional
	Backups bool   `json:"backups,omitempty"`
	Image   string `json:"image"`
	// +optional
	Ipv6 bool `json:"ipv6,omitempty"`
	// +optional
	Monitoring bool   `json:"monitoring,omitempty"`
	Name       string `json:"name"`
	// +optional
	PrivateNetworking bool   `json:"private_networking,omitempty"`
	Region            string `json:"region"`
	// +optional
	ResizeDisk bool   `json:"resize_disk,omitempty"`
	Size       string `json:"size"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"ssh_keys,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
	// +optional
	UserData string `json:"user_data,omitempty"`
}

type DropletStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
