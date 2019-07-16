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

type BigqueryTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigqueryTableSpec   `json:"spec,omitempty"`
	Status            BigqueryTableStatus `json:"status,omitempty"`
}

type BigqueryTableSpecTimePartitioning struct {
	// +optional
	ExpirationMs int `json:"expiration_ms,omitempty"`
	// +optional
	Field string `json:"field,omitempty"`
	Type  string `json:"type"`
}

type BigqueryTableSpecView struct {
	Query string `json:"query"`
	// +optional
	UseLegacySql bool `json:"use_legacy_sql,omitempty"`
}

type BigqueryTableSpec struct {
	DatasetId string `json:"dataset_id"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	FriendlyName string `json:"friendly_name,omitempty"`
	// +optional
	Labels  map[string]string `json:"labels,omitempty"`
	TableId string            `json:"table_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimePartitioning *[]BigqueryTableSpec `json:"time_partitioning,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	View *[]BigqueryTableSpec `json:"view,omitempty"`
}

type BigqueryTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigqueryTableList is a list of BigqueryTables
type BigqueryTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigqueryTable CRD objects
	Items []BigqueryTable `json:"items,omitempty"`
}
