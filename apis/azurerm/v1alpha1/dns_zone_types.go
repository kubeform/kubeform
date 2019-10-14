package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
