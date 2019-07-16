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

type DnsMxRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsMxRecordSpec   `json:"spec,omitempty"`
	Status            DnsMxRecordStatus `json:"status,omitempty"`
}

type DnsMxRecordSpecRecord struct {
	Exchange   string `json:"exchange"`
	Preference string `json:"preference"`
}

type DnsMxRecordSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Record            []DnsMxRecordSpec `json:"record"`
	ResourceGroupName string            `json:"resource_group_name"`
	Ttl               int               `json:"ttl"`
	ZoneName          string            `json:"zone_name"`
}

type DnsMxRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsMxRecordList is a list of DnsMxRecords
type DnsMxRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsMxRecord CRD objects
	Items []DnsMxRecord `json:"items,omitempty"`
}
