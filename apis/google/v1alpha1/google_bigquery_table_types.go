package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleBigqueryTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBigqueryTableSpec   `json:"spec,omitempty"`
	Status            GoogleBigqueryTableStatus `json:"status,omitempty"`
}

type GoogleBigqueryTableSpecTimePartitioning struct {
	Type         string `json:"type"`
	Field        string `json:"field"`
	ExpirationMs int    `json:"expiration_ms"`
}

type GoogleBigqueryTableSpecView struct {
	Query        string `json:"query"`
	UseLegacySql bool   `json:"use_legacy_sql"`
}

type GoogleBigqueryTableSpec struct {
	Location         string                    `json:"location"`
	NumRows          int                       `json:"num_rows"`
	DatasetId        string                    `json:"dataset_id"`
	Project          string                    `json:"project"`
	Description      string                    `json:"description"`
	TableId          string                    `json:"table_id"`
	TimePartitioning []GoogleBigqueryTableSpec `json:"time_partitioning"`
	Etag             string                    `json:"etag"`
	Labels           map[string]string         `json:"labels"`
	Schema           string                    `json:"schema"`
	CreationTime     int                       `json:"creation_time"`
	LastModifiedTime int                       `json:"last_modified_time"`
	NumBytes         int                       `json:"num_bytes"`
	NumLongTermBytes int                       `json:"num_long_term_bytes"`
	SelfLink         string                    `json:"self_link"`
	ExpirationTime   int                       `json:"expiration_time"`
	FriendlyName     string                    `json:"friendly_name"`
	View             []GoogleBigqueryTableSpec `json:"view"`
	Type             string                    `json:"type"`
}



type GoogleBigqueryTableStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBigqueryTableList is a list of GoogleBigqueryTables
type GoogleBigqueryTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigqueryTable CRD objects
	Items []GoogleBigqueryTable `json:"items,omitempty"`
}