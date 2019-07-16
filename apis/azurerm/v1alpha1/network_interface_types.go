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

type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

type NetworkInterfaceSpecIpConfiguration struct {
	Name string `json:"name"`
	// +optional
	PrivateIpAddress           string `json:"private_ip_address,omitempty"`
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	// +optional
	PrivateIpAddressVersion string `json:"private_ip_address_version,omitempty"`
	// +optional
	PublicIpAddressId string `json:"public_ip_address_id,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
}

type NetworkInterfaceSpec struct {
	// +optional
	EnableAcceleratedNetworking bool `json:"enable_accelerated_networking,omitempty"`
	// +optional
	EnableIpForwarding bool                   `json:"enable_ip_forwarding,omitempty"`
	IpConfiguration    []NetworkInterfaceSpec `json:"ip_configuration"`
	Location           string                 `json:"location"`
	Name               string                 `json:"name"`
	// +optional
	NetworkSecurityGroupId string `json:"network_security_group_id,omitempty"`
	ResourceGroupName      string `json:"resource_group_name"`
}

type NetworkInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
