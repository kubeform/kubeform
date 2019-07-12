package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDnsNsRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsNsRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsNsRecordStatus `json:"status,omitempty"`
}

type AzurermDnsNsRecordSpecRecord struct {
	Nsdname string `json:"nsdname"`
}

type AzurermDnsNsRecordSpec struct {
	ResourceGroupName string                   `json:"resource_group_name"`
	ZoneName          string                   `json:"zone_name"`
	Records           []string                 `json:"records"`
	Record            []AzurermDnsNsRecordSpec `json:"record"`
	Ttl               int                      `json:"ttl"`
	Tags              map[string]string        `json:"tags"`
	Name              string                   `json:"name"`
}

type AzurermDnsNsRecordStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDnsNsRecordList is a list of AzurermDnsNsRecords
type AzurermDnsNsRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsNsRecord CRD objects
	Items []AzurermDnsNsRecord `json:"items,omitempty"`
}
