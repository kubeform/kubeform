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

type AzurermDnsARecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsARecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsARecordStatus `json:"status,omitempty"`
}

type AzurermDnsARecordSpec struct {
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	ZoneName          string            `json:"zone_name"`
	Records           []string          `json:"records"`
	Ttl               int               `json:"ttl"`
	Tags              map[string]string `json:"tags"`
}



type AzurermDnsARecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsARecordList is a list of AzurermDnsARecords
type AzurermDnsARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsARecord CRD objects
	Items []AzurermDnsARecord `json:"items,omitempty"`
}