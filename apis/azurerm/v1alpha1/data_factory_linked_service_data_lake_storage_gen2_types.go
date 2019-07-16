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

type DataFactoryLinkedServiceDataLakeStorageGen2 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServiceDataLakeStorageGen2Spec   `json:"spec,omitempty"`
	Status            DataFactoryLinkedServiceDataLakeStorageGen2Status `json:"status,omitempty"`
}

type DataFactoryLinkedServiceDataLakeStorageGen2Spec struct {
	// +optional
	AdditionalProperties map[string]string `json:"additional_properties,omitempty"`
	// +optional
	Annotations     []string `json:"annotations,omitempty"`
	DataFactoryName string   `json:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	IntegrationRuntimeName string `json:"integration_runtime_name,omitempty"`
	Name                   string `json:"name"`
	// +optional
	Parameters          map[string]string `json:"parameters,omitempty"`
	ResourceGroupName   string            `json:"resource_group_name"`
	ServicePrincipalId  string            `json:"service_principal_id"`
	ServicePrincipalKey string            `json:"service_principal_key"`
	Tenant              string            `json:"tenant"`
	Url                 string            `json:"url"`
}

type DataFactoryLinkedServiceDataLakeStorageGen2Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryLinkedServiceDataLakeStorageGen2List is a list of DataFactoryLinkedServiceDataLakeStorageGen2s
type DataFactoryLinkedServiceDataLakeStorageGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryLinkedServiceDataLakeStorageGen2 CRD objects
	Items []DataFactoryLinkedServiceDataLakeStorageGen2 `json:"items,omitempty"`
}
