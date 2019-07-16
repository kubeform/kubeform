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

type DataFactoryDatasetSqlServerTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryDatasetSqlServerTableSpec   `json:"spec,omitempty"`
	Status            DataFactoryDatasetSqlServerTableStatus `json:"status,omitempty"`
}

type DataFactoryDatasetSqlServerTableSpecSchemaColumn struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	Type string `json:"type,omitempty"`
}

type DataFactoryDatasetSqlServerTableSpec struct {
	// +optional
	AdditionalProperties map[string]string `json:"additional_properties,omitempty"`
	// +optional
	Annotations     []string `json:"annotations,omitempty"`
	DataFactoryName string   `json:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Folder            string `json:"folder,omitempty"`
	LinkedServiceName string `json:"linked_service_name"`
	Name              string `json:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty"`
	ResourceGroupName string            `json:"resource_group_name"`
	// +optional
	SchemaColumn *[]DataFactoryDatasetSqlServerTableSpec `json:"schema_column,omitempty"`
	// +optional
	TableName string `json:"table_name,omitempty"`
}

type DataFactoryDatasetSqlServerTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryDatasetSqlServerTableList is a list of DataFactoryDatasetSqlServerTables
type DataFactoryDatasetSqlServerTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryDatasetSqlServerTable CRD objects
	Items []DataFactoryDatasetSqlServerTable `json:"items,omitempty"`
}
