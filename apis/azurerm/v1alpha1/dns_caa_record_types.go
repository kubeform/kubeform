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

type DnsCaaRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsCaaRecordSpec   `json:"spec,omitempty"`
	Status            DnsCaaRecordStatus `json:"status,omitempty"`
}

type DnsCaaRecordSpecRecord struct {
	Flags int    `json:"flags"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type DnsCaaRecordSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Record            []DnsCaaRecordSpec `json:"record"`
	ResourceGroupName string             `json:"resource_group_name"`
	Ttl               int                `json:"ttl"`
	ZoneName          string             `json:"zone_name"`
}

type DnsCaaRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsCaaRecordList is a list of DnsCaaRecords
type DnsCaaRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsCaaRecord CRD objects
	Items []DnsCaaRecord `json:"items,omitempty"`
}
