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

type DnsPtrRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsPtrRecordSpec   `json:"spec,omitempty"`
	Status            DnsPtrRecordStatus `json:"status,omitempty"`
}

type DnsPtrRecordSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Records           []string `json:"records"`
	ResourceGroupName string   `json:"resource_group_name"`
	Ttl               int      `json:"ttl"`
	ZoneName          string   `json:"zone_name"`
}

type DnsPtrRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsPtrRecordList is a list of DnsPtrRecords
type DnsPtrRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsPtrRecord CRD objects
	Items []DnsPtrRecord `json:"items,omitempty"`
}
