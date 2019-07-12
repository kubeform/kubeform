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

type AzurermDataFactoryDatasetPostgresql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryDatasetPostgresqlSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryDatasetPostgresqlStatus `json:"status,omitempty"`
}

type AzurermDataFactoryDatasetPostgresqlSpecSchemaColumn struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type AzurermDataFactoryDatasetPostgresqlSpec struct {
	Folder               string                                    `json:"folder"`
	SchemaColumn         []AzurermDataFactoryDatasetPostgresqlSpec `json:"schema_column"`
	Name                 string                                    `json:"name"`
	ResourceGroupName    string                                    `json:"resource_group_name"`
	LinkedServiceName    string                                    `json:"linked_service_name"`
	TableName            string                                    `json:"table_name"`
	Parameters           map[string]string                         `json:"parameters"`
	Annotations          []string                                  `json:"annotations"`
	DataFactoryName      string                                    `json:"data_factory_name"`
	Description          string                                    `json:"description"`
	AdditionalProperties map[string]string                         `json:"additional_properties"`
}

type AzurermDataFactoryDatasetPostgresqlStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataFactoryDatasetPostgresqlList is a list of AzurermDataFactoryDatasetPostgresqls
type AzurermDataFactoryDatasetPostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryDatasetPostgresql CRD objects
	Items []AzurermDataFactoryDatasetPostgresql `json:"items,omitempty"`
}
