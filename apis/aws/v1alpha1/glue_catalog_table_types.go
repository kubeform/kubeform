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

type GlueCatalogTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueCatalogTableSpec   `json:"spec,omitempty"`
	Status            GlueCatalogTableStatus `json:"status,omitempty"`
}

type GlueCatalogTableSpecPartitionKeys struct {
	// +optional
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	// +optional
	Type string `json:"type,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorColumns struct {
	// +optional
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	// +optional
	Type string `json:"type,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorSerDeInfo struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	SerializationLibrary string `json:"serialization_library,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorSkewedInfo struct {
	// +optional
	SkewedColumnNames []string `json:"skewed_column_names,omitempty"`
	// +optional
	SkewedColumnValueLocationMaps map[string]string `json:"skewed_column_value_location_maps,omitempty"`
	// +optional
	SkewedColumnValues []string `json:"skewed_column_values,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorSortColumns struct {
	Column    string `json:"column"`
	SortOrder int    `json:"sort_order"`
}

type GlueCatalogTableSpecStorageDescriptor struct {
	// +optional
	BucketColumns []string `json:"bucket_columns,omitempty"`
	// +optional
	Columns *[]GlueCatalogTableSpecStorageDescriptor `json:"columns,omitempty"`
	// +optional
	Compressed bool `json:"compressed,omitempty"`
	// +optional
	InputFormat string `json:"input_format,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	NumberOfBuckets int `json:"number_of_buckets,omitempty"`
	// +optional
	OutputFormat string `json:"output_format,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SerDeInfo *[]GlueCatalogTableSpecStorageDescriptor `json:"ser_de_info,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SkewedInfo *[]GlueCatalogTableSpecStorageDescriptor `json:"skewed_info,omitempty"`
	// +optional
	SortColumns *[]GlueCatalogTableSpecStorageDescriptor `json:"sort_columns,omitempty"`
	// +optional
	StoredAsSubDirectories bool `json:"stored_as_sub_directories,omitempty"`
}

type GlueCatalogTableSpec struct {
	DatabaseName string `json:"database_name"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	Owner string `json:"owner,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	PartitionKeys *[]GlueCatalogTableSpec `json:"partition_keys,omitempty"`
	// +optional
	Retention int `json:"retention,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageDescriptor *[]GlueCatalogTableSpec `json:"storage_descriptor,omitempty"`
	// +optional
	TableType string `json:"table_type,omitempty"`
	// +optional
	ViewExpandedText string `json:"view_expanded_text,omitempty"`
	// +optional
	ViewOriginalText string `json:"view_original_text,omitempty"`
}

type GlueCatalogTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueCatalogTableList is a list of GlueCatalogTables
type GlueCatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueCatalogTable CRD objects
	Items []GlueCatalogTable `json:"items,omitempty"`
}
