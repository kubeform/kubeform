package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DataFactoryLinkedServiceMysql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServiceMysqlSpec   `json:"spec,omitempty"`
	Status            DataFactoryLinkedServiceMysqlStatus `json:"status,omitempty"`
}

type DataFactoryLinkedServiceMysqlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`
	// +optional
	Annotations      []string `json:"annotations,omitempty" tf:"annotations,omitempty"`
	ConnectionString string   `json:"connectionString" tf:"connection_string"`
	DataFactoryName  string   `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	IntegrationRuntimeName string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name"`
}

type DataFactoryLinkedServiceMysqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactoryLinkedServiceMysqlSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryLinkedServiceMysqlList is a list of DataFactoryLinkedServiceMysqls
type DataFactoryLinkedServiceMysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryLinkedServiceMysql CRD objects
	Items []DataFactoryLinkedServiceMysql `json:"items,omitempty"`
}
