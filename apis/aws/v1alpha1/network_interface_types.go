package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

type NetworkInterfaceSpecAttachment struct {
	DeviceIndex int    `json:"deviceIndex" tf:"device_index"`
	Instance    string `json:"instance" tf:"instance"`
}

type NetworkInterfaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Attachment []NetworkInterfaceSpecAttachment `json:"attachment,omitempty" tf:"attachment,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	PrivateIP string `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	PrivateIPS []string `json:"privateIPS,omitempty" tf:"private_ips,omitempty"`
	// +optional
	PrivateIPSCount int `json:"privateIPSCount,omitempty" tf:"private_ips_count,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SourceDestCheck bool   `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`
	SubnetID        string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *NetworkInterfaceSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceList is a list of NetworkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterface CRD objects
	Items []NetworkInterface `json:"items,omitempty"`
}
