package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDnsAaaaRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsAaaaRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsAaaaRecordStatus `json:"status,omitempty"`
}

type AzurermDnsAaaaRecordSpec struct {
	Records           []string          `json:"records"`
	Ttl               int               `json:"ttl"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	ZoneName          string            `json:"zone_name"`
}



type AzurermDnsAaaaRecordStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDnsAaaaRecordList is a list of AzurermDnsAaaaRecords
type AzurermDnsAaaaRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsAaaaRecord CRD objects
	Items []AzurermDnsAaaaRecord `json:"items,omitempty"`
}