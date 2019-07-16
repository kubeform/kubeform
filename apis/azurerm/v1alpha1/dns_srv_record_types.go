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

type DnsSrvRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsSrvRecordSpec   `json:"spec,omitempty"`
	Status            DnsSrvRecordStatus `json:"status,omitempty"`
}

type DnsSrvRecordSpecRecord struct {
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Target   string `json:"target"`
	Weight   int    `json:"weight"`
}

type DnsSrvRecordSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Record            []DnsSrvRecordSpec `json:"record"`
	ResourceGroupName string             `json:"resource_group_name"`
	Ttl               int                `json:"ttl"`
	ZoneName          string             `json:"zone_name"`
}

type DnsSrvRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsSrvRecordList is a list of DnsSrvRecords
type DnsSrvRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsSrvRecord CRD objects
	Items []DnsSrvRecord `json:"items,omitempty"`
}
