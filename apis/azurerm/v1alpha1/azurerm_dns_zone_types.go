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

type AzurermDnsZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsZoneSpec   `json:"spec,omitempty"`
	Status            AzurermDnsZoneStatus `json:"status,omitempty"`
}

type AzurermDnsZoneSpec struct {
	ResourceGroupName             string            `json:"resource_group_name"`
	NumberOfRecordSets            int               `json:"number_of_record_sets"`
	MaxNumberOfRecordSets         int               `json:"max_number_of_record_sets"`
	Tags                          map[string]string `json:"tags"`
	Name                          string            `json:"name"`
	NameServers                   []string          `json:"name_servers"`
	ZoneType                      string            `json:"zone_type"`
	RegistrationVirtualNetworkIds []string          `json:"registration_virtual_network_ids"`
	ResolutionVirtualNetworkIds   []string          `json:"resolution_virtual_network_ids"`
}



type AzurermDnsZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsZoneList is a list of AzurermDnsZones
type AzurermDnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsZone CRD objects
	Items []AzurermDnsZone `json:"items,omitempty"`
}