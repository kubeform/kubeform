package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Subnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec,omitempty"`
	Status            SubnetStatus `json:"status,omitempty"`
}

type SubnetSpecDelegationServiceDelegation struct {
	// +optional
	Actions []string `json:"actions,omitempty" tf:"actions,omitempty"`
	Name    string   `json:"name" tf:"name"`
}

type SubnetSpecDelegation struct {
	Name string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	ServiceDelegation []SubnetSpecDelegationServiceDelegation `json:"serviceDelegation" tf:"service_delegation"`
}

type SubnetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`
	// +optional
	Delegation []SubnetSpecDelegation `json:"delegation,omitempty" tf:"delegation,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpConfigurations []string `json:"ipConfigurations,omitempty" tf:"ip_configurations,omitempty"`
	Name             string   `json:"name" tf:"name"`
	// +optional
	// Deprecated
	NetworkSecurityGroupID string `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id,omitempty"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// Deprecated
	RouteTableID string `json:"routeTableID,omitempty" tf:"route_table_id,omitempty"`
	// +optional
	ServiceEndpoints   []string `json:"serviceEndpoints,omitempty" tf:"service_endpoints,omitempty"`
	VirtualNetworkName string   `json:"virtualNetworkName" tf:"virtual_network_name"`
}

type SubnetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SubnetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubnetList is a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Subnet CRD objects
	Items []Subnet `json:"items,omitempty"`
}
