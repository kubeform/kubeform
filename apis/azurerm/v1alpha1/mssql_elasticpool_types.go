package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MssqlElasticpool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlElasticpoolSpec   `json:"spec,omitempty"`
	Status            MssqlElasticpoolStatus `json:"status,omitempty"`
}

type MssqlElasticpoolSpecPerDatabaseSettings struct {
	MaxCapacity json.Number `json:"max_capacity"`
	MinCapacity json.Number `json:"min_capacity"`
}

type MssqlElasticpoolSpecSku struct {
	Capacity int `json:"capacity"`
	// +optional
	Family string `json:"family,omitempty"`
	Name   string `json:"name"`
	Tier   string `json:"tier"`
}

type MssqlElasticpoolSpec struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	PerDatabaseSettings []MssqlElasticpoolSpec `json:"per_database_settings"`
	ResourceGroupName   string                 `json:"resource_group_name"`
	ServerName          string                 `json:"server_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []MssqlElasticpoolSpec `json:"sku"`
	// +optional
	ZoneRedundant bool `json:"zone_redundant,omitempty"`
}

type MssqlElasticpoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MssqlElasticpoolList is a list of MssqlElasticpools
type MssqlElasticpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MssqlElasticpool CRD objects
	Items []MssqlElasticpool `json:"items,omitempty"`
}
