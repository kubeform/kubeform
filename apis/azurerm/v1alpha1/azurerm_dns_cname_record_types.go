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

type AzurermDnsCnameRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsCnameRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsCnameRecordStatus `json:"status,omitempty"`
}

type AzurermDnsCnameRecordSpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	ZoneName          string            `json:"zone_name"`
	Records           string            `json:"records"`
	Record            string            `json:"record"`
	Ttl               int               `json:"ttl"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
}



type AzurermDnsCnameRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsCnameRecordList is a list of AzurermDnsCnameRecords
type AzurermDnsCnameRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsCnameRecord CRD objects
	Items []AzurermDnsCnameRecord `json:"items,omitempty"`
}