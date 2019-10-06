package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`
	// +optional
	Annotations     []string `json:"annotations,omitempty" tf:"annotations,omitempty"`
	DataFactoryName string   `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	IntegrationRuntimeName string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	Parameters          map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	ResourceGroupName   string            `json:"resourceGroupName" tf:"resource_group_name"`
	ServicePrincipalID  string            `json:"servicePrincipalID" tf:"service_principal_id"`
	ServicePrincipalKey string            `json:"servicePrincipalKey" tf:"service_principal_key"`
	Tenant              string            `json:"tenant" tf:"tenant"`
	Url                 string            `json:"url" tf:"url"`
}

type DataFactoryLinkedServiceDataLakeStorageGen2Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactoryLinkedServiceDataLakeStorageGen2Spec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
