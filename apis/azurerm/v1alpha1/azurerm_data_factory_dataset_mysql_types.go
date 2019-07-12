package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataFactoryDatasetMysql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryDatasetMysqlSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryDatasetMysqlStatus `json:"status,omitempty"`
}

type AzurermDataFactoryDatasetMysqlSpecSchemaColumn struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type AzurermDataFactoryDatasetMysqlSpec struct {
	Parameters           map[string]string                    `json:"parameters"`
	Annotations          []string                             `json:"annotations"`
	DataFactoryName      string                               `json:"data_factory_name"`
	ResourceGroupName    string                               `json:"resource_group_name"`
	LinkedServiceName    string                               `json:"linked_service_name"`
	Folder               string                               `json:"folder"`
	AdditionalProperties map[string]string                    `json:"additional_properties"`
	SchemaColumn         []AzurermDataFactoryDatasetMysqlSpec `json:"schema_column"`
	Name                 string                               `json:"name"`
	TableName            string                               `json:"table_name"`
	Description          string                               `json:"description"`
}



type AzurermDataFactoryDatasetMysqlStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataFactoryDatasetMysqlList is a list of AzurermDataFactoryDatasetMysqls
type AzurermDataFactoryDatasetMysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryDatasetMysql CRD objects
	Items []AzurermDataFactoryDatasetMysql `json:"items,omitempty"`
}