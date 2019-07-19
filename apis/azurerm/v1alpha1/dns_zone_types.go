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

type DnsZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZoneSpec   `json:"spec,omitempty"`
	Status            DnsZoneStatus `json:"status,omitempty"`
}

type DnsZoneSpec struct {
	Name string `json:"name" tf:"name"`
	// +optional
	RegistrationVirtualNetworkIDS []string `json:"registrationVirtualNetworkIDS,omitempty" tf:"registration_virtual_network_ids,omitempty"`
	// +optional
	ResolutionVirtualNetworkIDS []string `json:"resolutionVirtualNetworkIDS,omitempty" tf:"resolution_virtual_network_ids,omitempty"`
	ResourceGroupName           string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ZoneType    string                    `json:"zoneType,omitempty" tf:"zone_type,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DnsZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
