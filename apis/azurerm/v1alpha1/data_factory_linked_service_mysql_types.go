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

type DataFactoryLinkedServiceMysql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServiceMysqlSpec   `json:"spec,omitempty"`
	Status            DataFactoryLinkedServiceMysqlStatus `json:"status,omitempty"`
}

type DataFactoryLinkedServiceMysqlSpec struct {
	// +optional
	AdditionalProperties map[string]string `json:"additional_properties,omitempty"`
	// +optional
	Annotations      []string `json:"annotations,omitempty"`
	ConnectionString string   `json:"connection_string"`
	DataFactoryName  string   `json:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	IntegrationRuntimeName string `json:"integration_runtime_name,omitempty"`
	Name                   string `json:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty"`
	ResourceGroupName string            `json:"resource_group_name"`
}

type DataFactoryLinkedServiceMysqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
