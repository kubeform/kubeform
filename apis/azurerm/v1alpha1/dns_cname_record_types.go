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

type DnsCnameRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsCnameRecordSpec   `json:"spec,omitempty"`
	Status            DnsCnameRecordStatus `json:"status,omitempty"`
}

type DnsCnameRecordSpec struct {
	Name              string `json:"name"`
	Record            string `json:"record"`
	ResourceGroupName string `json:"resource_group_name"`
	Ttl               int    `json:"ttl"`
	ZoneName          string `json:"zone_name"`
}

type DnsCnameRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsCnameRecordList is a list of DnsCnameRecords
type DnsCnameRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsCnameRecord CRD objects
	Items []DnsCnameRecord `json:"items,omitempty"`
}
