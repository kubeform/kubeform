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

type KmsGrant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsGrantSpec   `json:"spec,omitempty"`
	Status            KmsGrantStatus `json:"status,omitempty"`
}

type KmsGrantSpecConstraints struct {
	// +optional
	EncryptionContextEquals map[string]string `json:"encryptionContextEquals,omitempty" tf:"encryption_context_equals,omitempty"`
	// +optional
	EncryptionContextSubset map[string]string `json:"encryptionContextSubset,omitempty" tf:"encryption_context_subset,omitempty"`
}

type KmsGrantSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Constraints []KmsGrantSpecConstraints `json:"constraints,omitempty" tf:"constraints,omitempty"`
	// +optional
	GrantCreationTokens []string `json:"grantCreationTokens,omitempty" tf:"grant_creation_tokens,omitempty"`
	// +optional
	GrantID string `json:"grantID,omitempty" tf:"grant_id,omitempty"`
	// +optional
	GrantToken       string `json:"grantToken,omitempty" tf:"grant_token,omitempty"`
	GranteePrincipal string `json:"granteePrincipal" tf:"grantee_principal"`
	KeyID            string `json:"keyID" tf:"key_id"`
	// +optional
	Name       string   `json:"name,omitempty" tf:"name,omitempty"`
	Operations []string `json:"operations" tf:"operations"`
	// +optional
	RetireOnDelete bool `json:"retireOnDelete,omitempty" tf:"retire_on_delete,omitempty"`
	// +optional
	RetiringPrincipal string `json:"retiringPrincipal,omitempty" tf:"retiring_principal,omitempty"`
}

type KmsGrantStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KmsGrantSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsGrantList is a list of KmsGrants
type KmsGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsGrant CRD objects
	Items []KmsGrant `json:"items,omitempty"`
}
