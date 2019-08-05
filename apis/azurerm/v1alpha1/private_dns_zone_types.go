package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PrivateDNSZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSZoneSpec   `json:"spec,omitempty"`
	Status            PrivateDNSZoneStatus `json:"status,omitempty"`
}

type PrivateDNSZoneSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	MaxNumberOfRecordSets int `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets,omitempty"`
	// +optional
	MaxNumberOfVirtualNetworkLinks int `json:"maxNumberOfVirtualNetworkLinks,omitempty" tf:"max_number_of_virtual_network_links,omitempty"`
	// +optional
	MaxNumberOfVirtualNetworkLinksWithRegistration int    `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty" tf:"max_number_of_virtual_network_links_with_registration,omitempty"`
	Name                                           string `json:"name" tf:"name"`
	// +optional
	NumberOfRecordSets int    `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets,omitempty"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PrivateDNSZoneSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PrivateDNSZoneList is a list of PrivateDNSZones
type PrivateDNSZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrivateDNSZone CRD objects
	Items []PrivateDNSZone `json:"items,omitempty"`
}
