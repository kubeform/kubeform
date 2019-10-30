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

type DnsZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZoneSpec   `json:"spec,omitempty"`
	Status            DnsZoneStatus `json:"status,omitempty"`
}

type DnsZoneSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	MaxNumberOfRecordSets int    `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets,omitempty"`
	Name                  string `json:"name" tf:"name"`
	// +optional
	NameServers []string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`
	// +optional
	NumberOfRecordSets int `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets,omitempty"`
	// +optional
	RegistrationVirtualNetworkIDS []string `json:"registrationVirtualNetworkIDS,omitempty" tf:"registration_virtual_network_ids,omitempty"`
	// +optional
	ResolutionVirtualNetworkIDS []string `json:"resolutionVirtualNetworkIDS,omitempty" tf:"resolution_virtual_network_ids,omitempty"`
	ResourceGroupName           string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ZoneType string `json:"zoneType,omitempty" tf:"zone_type,omitempty"`
}

type DnsZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DnsZoneSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsZoneList is a list of DnsZones
type DnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsZone CRD objects
	Items []DnsZone `json:"items,omitempty"`
}
