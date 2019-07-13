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

type AzurermDnsMxRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsMxRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsMxRecordStatus `json:"status,omitempty"`
}

type AzurermDnsMxRecordSpecRecord struct {
	Preference string `json:"preference"`
	Exchange   string `json:"exchange"`
}

type AzurermDnsMxRecordSpec struct {
	Name              string                   `json:"name"`
	ResourceGroupName string                   `json:"resource_group_name"`
	ZoneName          string                   `json:"zone_name"`
	Record            []AzurermDnsMxRecordSpec `json:"record"`
	Ttl               int                      `json:"ttl"`
	Tags              map[string]string        `json:"tags"`
}



type AzurermDnsMxRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsMxRecordList is a list of AzurermDnsMxRecords
type AzurermDnsMxRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsMxRecord CRD objects
	Items []AzurermDnsMxRecord `json:"items,omitempty"`
}