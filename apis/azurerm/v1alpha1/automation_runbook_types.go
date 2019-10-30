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

type AutomationRunbook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationRunbookSpec   `json:"spec,omitempty"`
	Status            AutomationRunbookStatus `json:"status,omitempty"`
}

type AutomationRunbookSpecPublishContentLinkHash struct {
	Algorithm string `json:"algorithm" tf:"algorithm"`
	Value     string `json:"value" tf:"value"`
}

type AutomationRunbookSpecPublishContentLink struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Hash []AutomationRunbookSpecPublishContentLinkHash `json:"hash,omitempty" tf:"hash,omitempty"`
	Uri  string                                        `json:"uri" tf:"uri"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type AutomationRunbookSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName string `json:"accountName" tf:"account_name"`
	// +optional
	Content string `json:"content,omitempty" tf:"content,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Location    string `json:"location" tf:"location"`
	LogProgress bool   `json:"logProgress" tf:"log_progress"`
	LogVerbose  bool   `json:"logVerbose" tf:"log_verbose"`
	Name        string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	PublishContentLink []AutomationRunbookSpecPublishContentLink `json:"publishContentLink" tf:"publish_content_link"`
	ResourceGroupName  string                                    `json:"resourceGroupName" tf:"resource_group_name"`
	RunbookType        string                                    `json:"runbookType" tf:"runbook_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AutomationRunbookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutomationRunbookSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationRunbookList is a list of AutomationRunbooks
type AutomationRunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationRunbook CRD objects
	Items []AutomationRunbook `json:"items,omitempty"`
}
