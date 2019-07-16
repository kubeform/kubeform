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

type Subnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec,omitempty"`
	Status            SubnetStatus `json:"status,omitempty"`
}

type SubnetSpecDelegationServiceDelegation struct {
	// +optional
	Actions []string `json:"actions,omitempty"`
	Name    string   `json:"name"`
}

type SubnetSpecDelegation struct {
	Name string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	ServiceDelegation []SubnetSpecDelegation `json:"service_delegation"`
}

type SubnetSpec struct {
	AddressPrefix string `json:"address_prefix"`
	// +optional
	Delegation *[]SubnetSpec `json:"delegation,omitempty"`
	Name       string        `json:"name"`
	// +optional
	// Deprecated
	NetworkSecurityGroupId string `json:"network_security_group_id,omitempty"`
	ResourceGroupName      string `json:"resource_group_name"`
	// +optional
	// Deprecated
	RouteTableId string `json:"route_table_id,omitempty"`
	// +optional
	ServiceEndpoints   []string `json:"service_endpoints,omitempty"`
	VirtualNetworkName string   `json:"virtual_network_name"`
}

type SubnetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
