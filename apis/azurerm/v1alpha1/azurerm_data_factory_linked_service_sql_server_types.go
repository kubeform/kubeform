package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataFactoryLinkedServiceSqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryLinkedServiceSqlServerSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryLinkedServiceSqlServerStatus `json:"status,omitempty"`
}

type AzurermDataFactoryLinkedServiceSqlServerSpec struct {
	Name                   string            `json:"name"`
	ResourceGroupName      string            `json:"resource_group_name"`
	IntegrationRuntimeName string            `json:"integration_runtime_name"`
	Annotations            []string          `json:"annotations"`
	DataFactoryName        string            `json:"data_factory_name"`
	ConnectionString       string            `json:"connection_string"`
	Description            string            `json:"description"`
	Parameters             map[string]string `json:"parameters"`
	AdditionalProperties   map[string]string `json:"additional_properties"`
}

type AzurermDataFactoryLinkedServiceSqlServerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataFactoryLinkedServiceSqlServerList is a list of AzurermDataFactoryLinkedServiceSqlServers
type AzurermDataFactoryLinkedServiceSqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryLinkedServiceSqlServer CRD objects
	Items []AzurermDataFactoryLinkedServiceSqlServer `json:"items,omitempty"`
}
