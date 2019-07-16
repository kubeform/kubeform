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

type Ami struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiSpec   `json:"spec,omitempty"`
	Status            AmiStatus `json:"status,omitempty"`
}

type AmiSpec struct {
	// +optional
	Architecture string `json:"architecture,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EnaSupport bool `json:"ena_support,omitempty"`
	// +optional
	KernelId string `json:"kernel_id,omitempty"`
	Name     string `json:"name"`
	// +optional
	RamdiskId string `json:"ramdisk_id,omitempty"`
	// +optional
	RootDeviceName string `json:"root_device_name,omitempty"`
	// +optional
	SriovNetSupport string `json:"sriov_net_support,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	VirtualizationType string `json:"virtualization_type,omitempty"`
}

type AmiStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiList is a list of Amis
type AmiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ami CRD objects
	Items []Ami `json:"items,omitempty"`
}
