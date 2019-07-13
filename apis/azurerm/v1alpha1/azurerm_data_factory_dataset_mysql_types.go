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
	Annotations          []string                             `json:"annotations"`
	SchemaColumn         []AzurermDataFactoryDatasetMysqlSpec `json:"schema_column"`
	Name                 string                               `json:"name"`
	ResourceGroupName    string                               `json:"resource_group_name"`
	Parameters           map[string]string                    `json:"parameters"`
	Description          string                               `json:"description"`
	AdditionalProperties map[string]string                    `json:"additional_properties"`
	DataFactoryName      string                               `json:"data_factory_name"`
	LinkedServiceName    string                               `json:"linked_service_name"`
	TableName            string                               `json:"table_name"`
	Folder               string                               `json:"folder"`
}



type AzurermDataFactoryDatasetMysqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataFactoryDatasetMysqlList is a list of AzurermDataFactoryDatasetMysqls
type AzurermDataFactoryDatasetMysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryDatasetMysql CRD objects
	Items []AzurermDataFactoryDatasetMysql `json:"items,omitempty"`
}