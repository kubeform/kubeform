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

type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkStatus `json:"status,omitempty"`
}

type VirtualNetworkSpecDdosProtectionPlan struct {
	Enable bool   `json:"enable"`
	Id     string `json:"id"`
}

type VirtualNetworkSpec struct {
	// +kubebuilder:validation:MinItems=1
	AddressSpace []string `json:"address_space"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DdosProtectionPlan *[]VirtualNetworkSpec `json:"ddos_protection_plan,omitempty"`
	// +optional
	DnsServers        []string `json:"dns_servers,omitempty"`
	Location          string   `json:"location"`
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
}

type VirtualNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkList is a list of VirtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetwork CRD objects
	Items []VirtualNetwork `json:"items,omitempty"`
}
