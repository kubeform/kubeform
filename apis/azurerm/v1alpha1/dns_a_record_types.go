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

type DnsARecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsARecordSpec   `json:"spec,omitempty"`
	Status            DnsARecordStatus `json:"status,omitempty"`
}

type DnsARecordSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Records           []string `json:"records"`
	ResourceGroupName string   `json:"resource_group_name"`
	Ttl               int      `json:"ttl"`
	ZoneName          string   `json:"zone_name"`
}

type DnsARecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsARecordList is a list of DnsARecords
type DnsARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsARecord CRD objects
	Items []DnsARecord `json:"items,omitempty"`
}
