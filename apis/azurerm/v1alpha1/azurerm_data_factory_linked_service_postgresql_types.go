package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataFactoryLinkedServicePostgresql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryLinkedServicePostgresqlSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryLinkedServicePostgresqlStatus `json:"status,omitempty"`
}

type AzurermDataFactoryLinkedServicePostgresqlSpec struct {
	ResourceGroupName      string            `json:"resource_group_name"`
	ConnectionString       string            `json:"connection_string"`
	AdditionalProperties   map[string]string `json:"additional_properties"`
	Name                   string            `json:"name"`
	DataFactoryName        string            `json:"data_factory_name"`
	Description            string            `json:"description"`
	IntegrationRuntimeName string            `json:"integration_runtime_name"`
	Parameters             map[string]string `json:"parameters"`
	Annotations            []string          `json:"annotations"`
}

type AzurermDataFactoryLinkedServicePostgresqlStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataFactoryLinkedServicePostgresqlList is a list of AzurermDataFactoryLinkedServicePostgresqls
type AzurermDataFactoryLinkedServicePostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryLinkedServicePostgresql CRD objects
	Items []AzurermDataFactoryLinkedServicePostgresql `json:"items,omitempty"`
}
