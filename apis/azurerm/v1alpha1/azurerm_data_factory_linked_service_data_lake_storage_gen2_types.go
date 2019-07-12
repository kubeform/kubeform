package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataFactoryLinkedServiceDataLakeStorageGen2 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryLinkedServiceDataLakeStorageGen2Spec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryLinkedServiceDataLakeStorageGen2Status `json:"status,omitempty"`
}

type AzurermDataFactoryLinkedServiceDataLakeStorageGen2Spec struct {
	Tenant                 string            `json:"tenant"`
	IntegrationRuntimeName string            `json:"integration_runtime_name"`
	Parameters             map[string]string `json:"parameters"`
	DataFactoryName        string            `json:"data_factory_name"`
	ResourceGroupName      string            `json:"resource_group_name"`
	Url                    string            `json:"url"`
	ServicePrincipalId     string            `json:"service_principal_id"`
	AdditionalProperties   map[string]string `json:"additional_properties"`
	Name                   string            `json:"name"`
	ServicePrincipalKey    string            `json:"service_principal_key"`
	Description            string            `json:"description"`
	Annotations            []string          `json:"annotations"`
}

type AzurermDataFactoryLinkedServiceDataLakeStorageGen2Status struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataFactoryLinkedServiceDataLakeStorageGen2List is a list of AzurermDataFactoryLinkedServiceDataLakeStorageGen2s
type AzurermDataFactoryLinkedServiceDataLakeStorageGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryLinkedServiceDataLakeStorageGen2 CRD objects
	Items []AzurermDataFactoryLinkedServiceDataLakeStorageGen2 `json:"items,omitempty"`
}
