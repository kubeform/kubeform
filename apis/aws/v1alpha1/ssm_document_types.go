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

type SsmDocument struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmDocumentSpec   `json:"spec,omitempty"`
	Status            SsmDocumentStatus `json:"status,omitempty"`
}

type SsmDocumentSpecParameter struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type SsmDocumentSpecPermissions struct {
	AccountIDS string `json:"accountIDS" tf:"account_ids"`
	Type       string `json:"type" tf:"type"`
}

type SsmDocumentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn     string `json:"arn,omitempty" tf:"arn,omitempty"`
	Content string `json:"content" tf:"content"`
	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	DefaultVersion string `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DocumentFormat string `json:"documentFormat,omitempty" tf:"document_format,omitempty"`
	DocumentType   string `json:"documentType" tf:"document_type"`
	// +optional
	Hash string `json:"hash,omitempty" tf:"hash,omitempty"`
	// +optional
	HashType string `json:"hashType,omitempty" tf:"hash_type,omitempty"`
	// +optional
	LatestVersion string `json:"latestVersion,omitempty" tf:"latest_version,omitempty"`
	Name          string `json:"name" tf:"name"`
	// +optional
	Owner string `json:"owner,omitempty" tf:"owner,omitempty"`
	// +optional
	Parameter []SsmDocumentSpecParameter `json:"parameter,omitempty" tf:"parameter,omitempty"`
	// +optional
	Permissions map[string]SsmDocumentSpecPermissions `json:"permissions,omitempty" tf:"permissions,omitempty"`
	// +optional
	PlatformTypes []string `json:"platformTypes,omitempty" tf:"platform_types,omitempty"`
	// +optional
	SchemaVersion string `json:"schemaVersion,omitempty" tf:"schema_version,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SsmDocumentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmDocumentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmDocumentList is a list of SsmDocuments
type SsmDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmDocument CRD objects
	Items []SsmDocument `json:"items,omitempty"`
}
