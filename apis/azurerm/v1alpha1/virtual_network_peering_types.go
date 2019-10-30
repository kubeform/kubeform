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

type VirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkPeeringStatus `json:"status,omitempty"`
}

type VirtualNetworkPeeringSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowForwardedTraffic bool `json:"allowForwardedTraffic,omitempty" tf:"allow_forwarded_traffic,omitempty"`
	// +optional
	AllowGatewayTransit bool `json:"allowGatewayTransit,omitempty" tf:"allow_gateway_transit,omitempty"`
	// +optional
	AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess,omitempty" tf:"allow_virtual_network_access,omitempty"`
	Name                      string `json:"name" tf:"name"`
	RemoteVirtualNetworkID    string `json:"remoteVirtualNetworkID" tf:"remote_virtual_network_id"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	UseRemoteGateways  bool   `json:"useRemoteGateways,omitempty" tf:"use_remote_gateways,omitempty"`
	VirtualNetworkName string `json:"virtualNetworkName" tf:"virtual_network_name"`
}

type VirtualNetworkPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualNetworkPeeringSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkPeeringList is a list of VirtualNetworkPeerings
type VirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetworkPeering CRD objects
	Items []VirtualNetworkPeering `json:"items,omitempty"`
}
