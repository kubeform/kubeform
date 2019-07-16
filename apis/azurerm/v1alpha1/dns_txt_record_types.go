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

type DnsTxtRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsTxtRecordSpec   `json:"spec,omitempty"`
	Status            DnsTxtRecordStatus `json:"status,omitempty"`
}

type DnsTxtRecordSpecRecord struct {
	Value string `json:"value"`
}

type DnsTxtRecordSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Record            []DnsTxtRecordSpec `json:"record"`
	ResourceGroupName string             `json:"resource_group_name"`
	Ttl               int                `json:"ttl"`
	ZoneName          string             `json:"zone_name"`
}

type DnsTxtRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsTxtRecordList is a list of DnsTxtRecords
type DnsTxtRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsTxtRecord CRD objects
	Items []DnsTxtRecord `json:"items,omitempty"`
}
