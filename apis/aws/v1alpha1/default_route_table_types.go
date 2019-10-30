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

type DefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultRouteTableSpec   `json:"spec,omitempty"`
	Status            DefaultRouteTableStatus `json:"status,omitempty"`
}

type DefaultRouteTableSpecRoute struct {
	// +optional
	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
	// +optional
	EgressOnlyGatewayID string `json:"egressOnlyGatewayID,omitempty" tf:"egress_only_gateway_id,omitempty"`
	// +optional
	GatewayID string `json:"gatewayID,omitempty" tf:"gateway_id,omitempty"`
	// +optional
	InstanceID string `json:"instanceID,omitempty" tf:"instance_id,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	// +optional
	NatGatewayID string `json:"natGatewayID,omitempty" tf:"nat_gateway_id,omitempty"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	// +optional
	TransitGatewayID string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id,omitempty"`
	// +optional
	VpcPeeringConnectionID string `json:"vpcPeeringConnectionID,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type DefaultRouteTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DefaultRouteTableID string `json:"defaultRouteTableID" tf:"default_route_table_id"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	PropagatingVgws []string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`
	// +optional
	Route []DefaultRouteTableSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultRouteTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DefaultRouteTableSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultRouteTableList is a list of DefaultRouteTables
type DefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultRouteTable CRD objects
	Items []DefaultRouteTable `json:"items,omitempty"`
}
