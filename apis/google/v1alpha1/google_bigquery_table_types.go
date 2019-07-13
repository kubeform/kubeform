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

type GoogleBigqueryTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBigqueryTableSpec   `json:"spec,omitempty"`
	Status            GoogleBigqueryTableStatus `json:"status,omitempty"`
}

type GoogleBigqueryTableSpecView struct {
	Query        string `json:"query"`
	UseLegacySql bool   `json:"use_legacy_sql"`
}

type GoogleBigqueryTableSpecTimePartitioning struct {
	ExpirationMs int    `json:"expiration_ms"`
	Type         string `json:"type"`
	Field        string `json:"field"`
}

type GoogleBigqueryTableSpec struct {
	CreationTime     int                       `json:"creation_time"`
	TableId          string                    `json:"table_id"`
	DatasetId        string                    `json:"dataset_id"`
	View             []GoogleBigqueryTableSpec `json:"view"`
	TimePartitioning []GoogleBigqueryTableSpec `json:"time_partitioning"`
	NumBytes         int                       `json:"num_bytes"`
	Type             string                    `json:"type"`
	Project          string                    `json:"project"`
	FriendlyName     string                    `json:"friendly_name"`
	LastModifiedTime int                       `json:"last_modified_time"`
	Location         string                    `json:"location"`
	ExpirationTime   int                       `json:"expiration_time"`
	NumLongTermBytes int                       `json:"num_long_term_bytes"`
	NumRows          int                       `json:"num_rows"`
	SelfLink         string                    `json:"self_link"`
	Description      string                    `json:"description"`
	Labels           map[string]string         `json:"labels"`
	Schema           string                    `json:"schema"`
	Etag             string                    `json:"etag"`
}



type GoogleBigqueryTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBigqueryTableList is a list of GoogleBigqueryTables
type GoogleBigqueryTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigqueryTable CRD objects
	Items []GoogleBigqueryTable `json:"items,omitempty"`
}