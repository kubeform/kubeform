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

type DnsZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZoneSpec   `json:"spec,omitempty"`
	Status            DnsZoneStatus `json:"status,omitempty"`
}

type DnsZoneSpec struct {
	Name string `json:"name"`
	// +optional
	RegistrationVirtualNetworkIds []string `json:"registration_virtual_network_ids,omitempty"`
	// +optional
	ResolutionVirtualNetworkIds []string `json:"resolution_virtual_network_ids,omitempty"`
	ResourceGroupName           string   `json:"resource_group_name"`
	// +optional
	ZoneType string `json:"zone_type,omitempty"`
}

type DnsZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
