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

type AzurermDataFactoryDatasetSqlServerTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryDatasetSqlServerTableSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryDatasetSqlServerTableStatus `json:"status,omitempty"`
}

type AzurermDataFactoryDatasetSqlServerTableSpecSchemaColumn struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type AzurermDataFactoryDatasetSqlServerTableSpec struct {
	Parameters           map[string]string                             `json:"parameters"`
	SchemaColumn         []AzurermDataFactoryDatasetSqlServerTableSpec `json:"schema_column"`
	Name                 string                                        `json:"name"`
	DataFactoryName      string                                        `json:"data_factory_name"`
	LinkedServiceName    string                                        `json:"linked_service_name"`
	TableName            string                                        `json:"table_name"`
	AdditionalProperties map[string]string                             `json:"additional_properties"`
	ResourceGroupName    string                                        `json:"resource_group_name"`
	Description          string                                        `json:"description"`
	Annotations          []string                                      `json:"annotations"`
	Folder               string                                        `json:"folder"`
}



type AzurermDataFactoryDatasetSqlServerTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataFactoryDatasetSqlServerTableList is a list of AzurermDataFactoryDatasetSqlServerTables
type AzurermDataFactoryDatasetSqlServerTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryDatasetSqlServerTable CRD objects
	Items []AzurermDataFactoryDatasetSqlServerTable `json:"items,omitempty"`
}