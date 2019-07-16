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

type DataFactoryLinkedServicePostgresql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServicePostgresqlSpec   `json:"spec,omitempty"`
	Status            DataFactoryLinkedServicePostgresqlStatus `json:"status,omitempty"`
}

type DataFactoryLinkedServicePostgresqlSpec struct {
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

type DataFactoryLinkedServicePostgresqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryLinkedServicePostgresqlList is a list of DataFactoryLinkedServicePostgresqls
type DataFactoryLinkedServicePostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryLinkedServicePostgresql CRD objects
	Items []DataFactoryLinkedServicePostgresql `json:"items,omitempty"`
}