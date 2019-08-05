package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	MaxCapacity json.Number `json:"maxCapacity" tf:"max_capacity"`
	MinCapacity json.Number `json:"minCapacity" tf:"min_capacity"`
}

type MssqlElasticpoolSpecSku struct {
	Capacity int `json:"capacity" tf:"capacity"`
	// +optional
	Family string `json:"family,omitempty" tf:"family,omitempty"`
	Name   string `json:"name" tf:"name"`
	Tier   string `json:"tier" tf:"tier"`
}

type MssqlElasticpoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Location string `json:"location" tf:"location"`
	// +optional
	MaxSizeBytes int `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`
	// +optional
	MaxSizeGb json.Number `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`
	Name      string      `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	PerDatabaseSettings []MssqlElasticpoolSpecPerDatabaseSettings `json:"perDatabaseSettings" tf:"per_database_settings"`
	ResourceGroupName   string                                    `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName          string                                    `json:"serverName" tf:"server_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []MssqlElasticpoolSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ZoneRedundant bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MssqlElasticpoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MssqlElasticpoolSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
