/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
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
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
