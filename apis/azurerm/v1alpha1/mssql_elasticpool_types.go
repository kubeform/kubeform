/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

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

type MssqlElasticpool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlElasticpoolSpec   `json:"spec,omitempty"`
	Status            MssqlElasticpoolStatus `json:"status,omitempty"`
}

type MssqlElasticpoolSpecElasticPoolProperties struct {
	// +optional
	// Deprecated
	CreationDate string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`
	// +optional
	// Deprecated
	LicenseType string `json:"licenseType,omitempty" tf:"license_type,omitempty"`
	// +optional
	// Deprecated
	MaxSizeBytes int64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`
	// +optional
	// Deprecated
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	// Deprecated
	ZoneRedundant bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MssqlElasticpoolSpecPerDatabaseSettings struct {
	MaxCapacity float64 `json:"maxCapacity" tf:"max_capacity"`
	MinCapacity float64 `json:"minCapacity" tf:"min_capacity"`
}

type MssqlElasticpoolSpecSku struct {
	Capacity int64 `json:"capacity" tf:"capacity"`
	// +optional
	Family string `json:"family,omitempty" tf:"family,omitempty"`
	Name   string `json:"name" tf:"name"`
	Tier   string `json:"tier" tf:"tier"`
}

type MssqlElasticpoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	ElasticPoolProperties []MssqlElasticpoolSpecElasticPoolProperties `json:"elasticPoolProperties,omitempty" tf:"elastic_pool_properties,omitempty"`
	Location              string                                      `json:"location" tf:"location"`
	// +optional
	MaxSizeBytes int64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`
	// +optional
	MaxSizeGb float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`
	Name      string  `json:"name" tf:"name"`
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
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
