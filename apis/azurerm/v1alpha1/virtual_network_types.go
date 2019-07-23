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

type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkStatus `json:"status,omitempty"`
}

type VirtualNetworkSpecDdosProtectionPlan struct {
	Enable bool   `json:"enable" tf:"enable"`
	ID     string `json:"ID" tf:"id"`
}

type VirtualNetworkSpecSubnet struct {
	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`
	Name          string `json:"name" tf:"name"`
	// +optional
	SecurityGroup string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`
}

type VirtualNetworkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:MinItems=1
	AddressSpace []string `json:"addressSpace" tf:"address_space"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DdosProtectionPlan []VirtualNetworkSpecDdosProtectionPlan `json:"ddosProtectionPlan,omitempty" tf:"ddos_protection_plan,omitempty"`
	// +optional
	DnsServers        []string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`
	Location          string   `json:"location" tf:"location"`
	Name              string   `json:"name" tf:"name"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Subnet []VirtualNetworkSpecSubnet `json:"subnet,omitempty" tf:"subnet,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
