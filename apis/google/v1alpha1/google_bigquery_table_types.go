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
	Labels           map[string]string         `json:"labels"`
	ExpirationTime   int                       `json:"expiration_time"`
	Description      string                    `json:"description"`
	View             []GoogleBigqueryTableSpec `json:"view"`
	CreationTime     int                       `json:"creation_time"`
	LastModifiedTime int                       `json:"last_modified_time"`
	TableId          string                    `json:"table_id"`
	Project          string                    `json:"project"`
	FriendlyName     string                    `json:"friendly_name"`
	Type             string                    `json:"type"`
	DatasetId        string                    `json:"dataset_id"`
	TimePartitioning []GoogleBigqueryTableSpec `json:"time_partitioning"`
	Etag             string                    `json:"etag"`
	Location         string                    `json:"location"`
	NumBytes         int                       `json:"num_bytes"`
	NumLongTermBytes int                       `json:"num_long_term_bytes"`
	NumRows          int                       `json:"num_rows"`
	SelfLink         string                    `json:"self_link"`
	Schema           string                    `json:"schema"`
}

type GoogleBigqueryTableStatus struct {
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
