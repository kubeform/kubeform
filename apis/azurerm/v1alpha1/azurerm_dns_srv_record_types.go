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

type AzurermDnsSrvRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsSrvRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsSrvRecordStatus `json:"status,omitempty"`
}

type AzurermDnsSrvRecordSpecRecord struct {
	Priority int    `json:"priority"`
	Weight   int    `json:"weight"`
	Port     int    `json:"port"`
	Target   string `json:"target"`
}

type AzurermDnsSrvRecordSpec struct {
	Ttl               int                       `json:"ttl"`
	Tags              map[string]string         `json:"tags"`
	Name              string                    `json:"name"`
	ResourceGroupName string                    `json:"resource_group_name"`
	ZoneName          string                    `json:"zone_name"`
	Record            []AzurermDnsSrvRecordSpec `json:"record"`
}

type AzurermDnsSrvRecordStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsSrvRecordList is a list of AzurermDnsSrvRecords
type AzurermDnsSrvRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsSrvRecord CRD objects
	Items []AzurermDnsSrvRecord `json:"items,omitempty"`
}
