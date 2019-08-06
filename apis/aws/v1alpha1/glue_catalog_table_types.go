package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorColumns struct {
	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorSerDeInfo struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	SerializationLibrary string `json:"serializationLibrary,omitempty" tf:"serialization_library,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorSkewedInfo struct {
	// +optional
	SkewedColumnNames []string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names,omitempty"`
	// +optional
	SkewedColumnValueLocationMaps map[string]string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps,omitempty"`
	// +optional
	SkewedColumnValues []string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values,omitempty"`
}

type GlueCatalogTableSpecStorageDescriptorSortColumns struct {
	Column    string `json:"column" tf:"column"`
	SortOrder int    `json:"sortOrder" tf:"sort_order"`
}

type GlueCatalogTableSpecStorageDescriptor struct {
	// +optional
	BucketColumns []string `json:"bucketColumns,omitempty" tf:"bucket_columns,omitempty"`
	// +optional
	Columns []GlueCatalogTableSpecStorageDescriptorColumns `json:"columns,omitempty" tf:"columns,omitempty"`
	// +optional
	Compressed bool `json:"compressed,omitempty" tf:"compressed,omitempty"`
	// +optional
	InputFormat string `json:"inputFormat,omitempty" tf:"input_format,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	NumberOfBuckets int `json:"numberOfBuckets,omitempty" tf:"number_of_buckets,omitempty"`
	// +optional
	OutputFormat string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SerDeInfo []GlueCatalogTableSpecStorageDescriptorSerDeInfo `json:"serDeInfo,omitempty" tf:"ser_de_info,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SkewedInfo []GlueCatalogTableSpecStorageDescriptorSkewedInfo `json:"skewedInfo,omitempty" tf:"skewed_info,omitempty"`
	// +optional
	SortColumns []GlueCatalogTableSpecStorageDescriptorSortColumns `json:"sortColumns,omitempty" tf:"sort_columns,omitempty"`
	// +optional
	StoredAsSubDirectories bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories,omitempty"`
}

type GlueCatalogTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CatalogID    string `json:"catalogID,omitempty" tf:"catalog_id,omitempty"`
	DatabaseName string `json:"databaseName" tf:"database_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Owner string `json:"owner,omitempty" tf:"owner,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	PartitionKeys []GlueCatalogTableSpecPartitionKeys `json:"partitionKeys,omitempty" tf:"partition_keys,omitempty"`
	// +optional
	Retention int `json:"retention,omitempty" tf:"retention,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageDescriptor []GlueCatalogTableSpecStorageDescriptor `json:"storageDescriptor,omitempty" tf:"storage_descriptor,omitempty"`
	// +optional
	TableType string `json:"tableType,omitempty" tf:"table_type,omitempty"`
	// +optional
	ViewExpandedText string `json:"viewExpandedText,omitempty" tf:"view_expanded_text,omitempty"`
	// +optional
	ViewOriginalText string `json:"viewOriginalText,omitempty" tf:"view_original_text,omitempty"`
}

type GlueCatalogTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlueCatalogTableSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
