package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDnsCaaRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsCaaRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsCaaRecordStatus `json:"status,omitempty"`
}

type AzurermDnsCaaRecordSpecRecord struct {
	Flags int    `json:"flags"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type AzurermDnsCaaRecordSpec struct {
	ResourceGroupName string                    `json:"resource_group_name"`
	ZoneName          string                    `json:"zone_name"`
	Record            []AzurermDnsCaaRecordSpec `json:"record"`
	Ttl               int                       `json:"ttl"`
	Tags              map[string]string         `json:"tags"`
	Name              string                    `json:"name"`
}

type AzurermDnsCaaRecordStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDnsCaaRecordList is a list of AzurermDnsCaaRecords
type AzurermDnsCaaRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsCaaRecord CRD objects
	Items []AzurermDnsCaaRecord `json:"items,omitempty"`
}
