package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueCatalogTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueCatalogTableSpec   `json:"spec,omitempty"`
	Status            AwsGlueCatalogTableStatus `json:"status,omitempty"`
}

type AwsGlueCatalogTableSpecPartitionKeys struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type AwsGlueCatalogTableSpecStorageDescriptorSerDeInfo struct {
	Name                 string            `json:"name"`
	Parameters           map[string]string `json:"parameters"`
	SerializationLibrary string            `json:"serialization_library"`
}

type AwsGlueCatalogTableSpecStorageDescriptorColumns struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type AwsGlueCatalogTableSpecStorageDescriptorSkewedInfo struct {
	SkewedColumnValues            []string          `json:"skewed_column_values"`
	SkewedColumnValueLocationMaps map[string]string `json:"skewed_column_value_location_maps"`
	SkewedColumnNames             []string          `json:"skewed_column_names"`
}

type AwsGlueCatalogTableSpecStorageDescriptorSortColumns struct {
	Column    string `json:"column"`
	SortOrder int    `json:"sort_order"`
}

type AwsGlueCatalogTableSpecStorageDescriptor struct {
	Parameters             map[string]string                          `json:"parameters"`
	SerDeInfo              []AwsGlueCatalogTableSpecStorageDescriptor `json:"ser_de_info"`
	StoredAsSubDirectories bool                                       `json:"stored_as_sub_directories"`
	Location               string                                     `json:"location"`
	Columns                []AwsGlueCatalogTableSpecStorageDescriptor `json:"columns"`
	Compressed             bool                                       `json:"compressed"`
	InputFormat            string                                     `json:"input_format"`
	NumberOfBuckets        int                                        `json:"number_of_buckets"`
	OutputFormat           string                                     `json:"output_format"`
	SkewedInfo             []AwsGlueCatalogTableSpecStorageDescriptor `json:"skewed_info"`
	SortColumns            []AwsGlueCatalogTableSpecStorageDescriptor `json:"sort_columns"`
	BucketColumns          []string                                   `json:"bucket_columns"`
}

type AwsGlueCatalogTableSpec struct {
	Retention         int                       `json:"retention"`
	TableType         string                    `json:"table_type"`
	ViewExpandedText  string                    `json:"view_expanded_text"`
	CatalogId         string                    `json:"catalog_id"`
	Owner             string                    `json:"owner"`
	Name              string                    `json:"name"`
	Parameters        map[string]string         `json:"parameters"`
	PartitionKeys     []AwsGlueCatalogTableSpec `json:"partition_keys"`
	StorageDescriptor []AwsGlueCatalogTableSpec `json:"storage_descriptor"`
	ViewOriginalText  string                    `json:"view_original_text"`
	DatabaseName      string                    `json:"database_name"`
	Description       string                    `json:"description"`
}

type AwsGlueCatalogTableStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueCatalogTableList is a list of AwsGlueCatalogTables
type AwsGlueCatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueCatalogTable CRD objects
	Items []AwsGlueCatalogTable `json:"items,omitempty"`
}
