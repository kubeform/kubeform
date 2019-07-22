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

type DnsAaaaRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsAaaaRecordSpec   `json:"spec,omitempty"`
	Status            DnsAaaaRecordStatus `json:"status,omitempty"`
}

type DnsAaaaRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Name        string                    `json:"name" tf:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Records           []string `json:"records" tf:"records"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Ttl      int               `json:"ttl" tf:"ttl"`
	ZoneName string            `json:"zoneName" tf:"zone_name"`
}

type DnsAaaaRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsAaaaRecordList is a list of DnsAaaaRecords
type DnsAaaaRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsAaaaRecord CRD objects
	Items []DnsAaaaRecord `json:"items,omitempty"`
}
