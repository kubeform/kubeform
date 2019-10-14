package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreationTime int    `json:"creationTime,omitempty" tf:"creation_time,omitempty"`
	DatasetID    string `json:"datasetID" tf:"dataset_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	// +optional
	ExpirationTime int `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`
	// +optional
	FriendlyName string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LastModifiedTime int `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	NumBytes int `json:"numBytes,omitempty" tf:"num_bytes,omitempty"`
	// +optional
	NumLongTermBytes int `json:"numLongTermBytes,omitempty" tf:"num_long_term_bytes,omitempty"`
	// +optional
	NumRows int `json:"numRows,omitempty" tf:"num_rows,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Schema string `json:"schema,omitempty" tf:"schema,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	TableID  string `json:"tableID" tf:"table_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TimePartitioning []BigqueryTableSpecTimePartitioning `json:"timePartitioning,omitempty" tf:"time_partitioning,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	View []BigqueryTableSpecView `json:"view,omitempty" tf:"view,omitempty"`
}

type BigqueryTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BigqueryTableSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
