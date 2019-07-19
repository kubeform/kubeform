package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DataFactoryDatasetMysql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryDatasetMysqlSpec   `json:"spec,omitempty"`
	Status            DataFactoryDatasetMysqlStatus `json:"status,omitempty"`
}

type DataFactoryDatasetMysqlSpecSchemaColumn struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataFactoryDatasetMysqlSpec struct {
	// +optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`
	// +optional
	Annotations     []string `json:"annotations,omitempty" tf:"annotations,omitempty"`
	DataFactoryName string   `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Folder            string `json:"folder,omitempty" tf:"folder,omitempty"`
	LinkedServiceName string `json:"linkedServiceName" tf:"linked_service_name"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SchemaColumn []DataFactoryDatasetMysqlSpecSchemaColumn `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
	// +optional
	TableName   string                    `json:"tableName,omitempty" tf:"table_name,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DataFactoryDatasetMysqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryDatasetMysqlList is a list of DataFactoryDatasetMysqls
type DataFactoryDatasetMysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryDatasetMysql CRD objects
	Items []DataFactoryDatasetMysql `json:"items,omitempty"`
}
