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

type AzurermPrivateDnsZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPrivateDnsZoneSpec   `json:"spec,omitempty"`
	Status            AzurermPrivateDnsZoneStatus `json:"status,omitempty"`
}

type AzurermPrivateDnsZoneSpec struct {
	NumberOfRecordSets                             int               `json:"number_of_record_sets"`
	MaxNumberOfRecordSets                          int               `json:"max_number_of_record_sets"`
	MaxNumberOfVirtualNetworkLinks                 int               `json:"max_number_of_virtual_network_links"`
	MaxNumberOfVirtualNetworkLinksWithRegistration int               `json:"max_number_of_virtual_network_links_with_registration"`
	ResourceGroupName                              string            `json:"resource_group_name"`
	Tags                                           map[string]string `json:"tags"`
	Name                                           string            `json:"name"`
}



type AzurermPrivateDnsZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPrivateDnsZoneList is a list of AzurermPrivateDnsZones
type AzurermPrivateDnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPrivateDnsZone CRD objects
	Items []AzurermPrivateDnsZone `json:"items,omitempty"`
}