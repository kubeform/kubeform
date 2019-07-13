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

type AwsGlueCatalogTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueCatalogTableSpec   `json:"spec,omitempty"`
	Status            AwsGlueCatalogTableStatus `json:"status,omitempty"`
}

type AwsGlueCatalogTableSpecStorageDescriptorSerDeInfo struct {
	Name                 string            `json:"name"`
	Parameters           map[string]string `json:"parameters"`
	SerializationLibrary string            `json:"serialization_library"`
}

type AwsGlueCatalogTableSpecStorageDescriptorSkewedInfo struct {
	SkewedColumnNames             []string          `json:"skewed_column_names"`
	SkewedColumnValues            []string          `json:"skewed_column_values"`
	SkewedColumnValueLocationMaps map[string]string `json:"skewed_column_value_location_maps"`
}

type AwsGlueCatalogTableSpecStorageDescriptorSortColumns struct {
	Column    string `json:"column"`
	SortOrder int    `json:"sort_order"`
}

type AwsGlueCatalogTableSpecStorageDescriptorColumns struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type AwsGlueCatalogTableSpecStorageDescriptor struct {
	OutputFormat           string                                     `json:"output_format"`
	Parameters             map[string]string                          `json:"parameters"`
	SerDeInfo              []AwsGlueCatalogTableSpecStorageDescriptor `json:"ser_de_info"`
	SkewedInfo             []AwsGlueCatalogTableSpecStorageDescriptor `json:"skewed_info"`
	SortColumns            []AwsGlueCatalogTableSpecStorageDescriptor `json:"sort_columns"`
	Location               string                                     `json:"location"`
	NumberOfBuckets        int                                        `json:"number_of_buckets"`
	Compressed             bool                                       `json:"compressed"`
	InputFormat            string                                     `json:"input_format"`
	StoredAsSubDirectories bool                                       `json:"stored_as_sub_directories"`
	BucketColumns          []string                                   `json:"bucket_columns"`
	Columns                []AwsGlueCatalogTableSpecStorageDescriptor `json:"columns"`
}

type AwsGlueCatalogTableSpecPartitionKeys struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type AwsGlueCatalogTableSpec struct {
	StorageDescriptor []AwsGlueCatalogTableSpec `json:"storage_descriptor"`
	TableType         string                    `json:"table_type"`
	ViewOriginalText  string                    `json:"view_original_text"`
	DatabaseName      string                    `json:"database_name"`
	Name              string                    `json:"name"`
	Parameters        map[string]string         `json:"parameters"`
	Retention         int                       `json:"retention"`
	ViewExpandedText  string                    `json:"view_expanded_text"`
	CatalogId         string                    `json:"catalog_id"`
	Description       string                    `json:"description"`
	Owner             string                    `json:"owner"`
	PartitionKeys     []AwsGlueCatalogTableSpec `json:"partition_keys"`
}



type AwsGlueCatalogTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueCatalogTableList is a list of AwsGlueCatalogTables
type AwsGlueCatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueCatalogTable CRD objects
	Items []AwsGlueCatalogTable `json:"items,omitempty"`
}