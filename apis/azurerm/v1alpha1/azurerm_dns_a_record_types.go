package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDnsARecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsARecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsARecordStatus `json:"status,omitempty"`
}

type AzurermDnsARecordSpec struct {
	ZoneName          string            `json:"zone_name"`
	Records           []string          `json:"records"`
	Ttl               int               `json:"ttl"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
}



type AzurermDnsARecordStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDnsARecordList is a list of AzurermDnsARecords
type AzurermDnsARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsARecord CRD objects
	Items []AzurermDnsARecord `json:"items,omitempty"`
}