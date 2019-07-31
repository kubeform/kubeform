package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ExpirationMs int `json:"expirationMs,omitempty" tf:"expiration_ms,omitempty"`
	// +optional
	Field string `json:"field,omitempty" tf:"field,omitempty"`
	Type  string `json:"type" tf:"type"`
}

type BigqueryTableSpecView struct {
	Query string `json:"query" tf:"query"`
	// +optional
	UseLegacySQL bool `json:"useLegacySQL,omitempty" tf:"use_legacy_sql,omitempty"`
}

type BigqueryTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DatasetID string `json:"datasetID" tf:"dataset_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ExpirationTime int `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`
	// +optional
	FriendlyName string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Schema  string `json:"schema,omitempty" tf:"schema,omitempty"`
	TableID string `json:"tableID" tf:"table_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimePartitioning []BigqueryTableSpecTimePartitioning `json:"timePartitioning,omitempty" tf:"time_partitioning,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	View []BigqueryTableSpecView `json:"view,omitempty" tf:"view,omitempty"`
}

type BigqueryTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
