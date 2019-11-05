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

type StorageObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageObjectAccessControlSpec   `json:"spec,omitempty"`
	Status            StorageObjectAccessControlStatus `json:"status,omitempty"`
}

type StorageObjectAccessControlSpecProjectTeam struct {
	// +optional
	ProjectNumber string `json:"projectNumber,omitempty" tf:"project_number,omitempty"`
	// +optional
	Team string `json:"team,omitempty" tf:"team,omitempty"`
}

type StorageObjectAccessControlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	Email  string `json:"email,omitempty" tf:"email,omitempty"`
	Entity string `json:"entity" tf:"entity"`
	// +optional
	EntityID string `json:"entityID,omitempty" tf:"entity_id,omitempty"`
	// +optional
	Generation int64  `json:"generation,omitempty" tf:"generation,omitempty"`
	Object     string `json:"object" tf:"object"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ProjectTeam []StorageObjectAccessControlSpecProjectTeam `json:"projectTeam,omitempty" tf:"project_team,omitempty"`
	Role        string                                      `json:"role" tf:"role"`
}

type StorageObjectAccessControlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageObjectAccessControlSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageObjectAccessControlList is a list of StorageObjectAccessControls
type StorageObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageObjectAccessControl CRD objects
	Items []StorageObjectAccessControl `json:"items,omitempty"`
}
