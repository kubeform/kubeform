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

type BigqueryDataset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigqueryDatasetSpec   `json:"spec,omitempty"`
	Status            BigqueryDatasetStatus `json:"status,omitempty"`
}

type BigqueryDatasetSpec struct {
	DatasetId string `json:"dataset_id"`
	// +optional
	DefaultTableExpirationMs int `json:"default_table_expiration_ms,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	FriendlyName string `json:"friendly_name,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
}

type BigqueryDatasetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigqueryDatasetList is a list of BigqueryDatasets
type BigqueryDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigqueryDataset CRD objects
	Items []BigqueryDataset `json:"items,omitempty"`
}
