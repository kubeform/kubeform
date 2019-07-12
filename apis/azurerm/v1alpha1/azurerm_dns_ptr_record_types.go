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

type AzurermDnsPtrRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsPtrRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsPtrRecordStatus `json:"status,omitempty"`
}

type AzurermDnsPtrRecordSpec struct {
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	ZoneName          string            `json:"zone_name"`
	Records           []string          `json:"records"`
	Ttl               int               `json:"ttl"`
	Tags              map[string]string `json:"tags"`
}

type AzurermDnsPtrRecordStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsPtrRecordList is a list of AzurermDnsPtrRecords
type AzurermDnsPtrRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsPtrRecord CRD objects
	Items []AzurermDnsPtrRecord `json:"items,omitempty"`
}
