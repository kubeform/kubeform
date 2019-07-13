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

type AzurermDataFactoryLinkedServiceDataLakeStorageGen2 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryLinkedServiceDataLakeStorageGen2Spec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryLinkedServiceDataLakeStorageGen2Status `json:"status,omitempty"`
}

type AzurermDataFactoryLinkedServiceDataLakeStorageGen2Spec struct {
	Parameters             map[string]string `json:"parameters"`
	DataFactoryName        string            `json:"data_factory_name"`
	Url                    string            `json:"url"`
	Tenant                 string            `json:"tenant"`
	Description            string            `json:"description"`
	IntegrationRuntimeName string            `json:"integration_runtime_name"`
	Annotations            []string          `json:"annotations"`
	AdditionalProperties   map[string]string `json:"additional_properties"`
	Name                   string            `json:"name"`
	ResourceGroupName      string            `json:"resource_group_name"`
	ServicePrincipalId     string            `json:"service_principal_id"`
	ServicePrincipalKey    string            `json:"service_principal_key"`
}



type AzurermDataFactoryLinkedServiceDataLakeStorageGen2Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataFactoryLinkedServiceDataLakeStorageGen2List is a list of AzurermDataFactoryLinkedServiceDataLakeStorageGen2s
type AzurermDataFactoryLinkedServiceDataLakeStorageGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryLinkedServiceDataLakeStorageGen2 CRD objects
	Items []AzurermDataFactoryLinkedServiceDataLakeStorageGen2 `json:"items,omitempty"`
}