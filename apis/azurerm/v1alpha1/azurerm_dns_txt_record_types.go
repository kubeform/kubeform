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

type AzurermDnsTxtRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDnsTxtRecordSpec   `json:"spec,omitempty"`
	Status            AzurermDnsTxtRecordStatus `json:"status,omitempty"`
}

type AzurermDnsTxtRecordSpecRecord struct {
	Value string `json:"value"`
}

type AzurermDnsTxtRecordSpec struct {
	Ttl               int                       `json:"ttl"`
	Tags              map[string]string         `json:"tags"`
	Name              string                    `json:"name"`
	ResourceGroupName string                    `json:"resource_group_name"`
	ZoneName          string                    `json:"zone_name"`
	Record            []AzurermDnsTxtRecordSpec `json:"record"`
}



type AzurermDnsTxtRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDnsTxtRecordList is a list of AzurermDnsTxtRecords
type AzurermDnsTxtRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDnsTxtRecord CRD objects
	Items []AzurermDnsTxtRecord `json:"items,omitempty"`
}