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

type CloudbuildTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudbuildTriggerSpec   `json:"spec,omitempty"`
	Status            CloudbuildTriggerStatus `json:"status,omitempty"`
}

type CloudbuildTriggerSpecBuildStep struct {
	// +optional
	Args string `json:"args,omitempty" tf:"args,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type CloudbuildTriggerSpecBuild struct {
	// +optional
	Images []string `json:"images,omitempty" tf:"images,omitempty"`
	// +optional
	Step []CloudbuildTriggerSpecBuildStep `json:"step,omitempty" tf:"step,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudbuildTriggerSpecTriggerTemplate struct {
	// +optional
	BranchName string `json:"branchName,omitempty" tf:"branch_name,omitempty"`
	// +optional
	CommitSha string `json:"commitSha,omitempty" tf:"commit_sha,omitempty"`
	// +optional
	Dir string `json:"dir,omitempty" tf:"dir,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	RepoName string `json:"repoName,omitempty" tf:"repo_name,omitempty"`
	// +optional
	TagName string `json:"tagName,omitempty" tf:"tag_name,omitempty"`
}

type CloudbuildTriggerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Contents of the build template.
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Build []CloudbuildTriggerSpecBuild `json:"build,omitempty" tf:"build,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Filename string `json:"filename,omitempty" tf:"filename,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Substitutions map[string]string `json:"substitutions,omitempty" tf:"substitutions,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TriggerTemplate []CloudbuildTriggerSpecTriggerTemplate `json:"triggerTemplate,omitempty" tf:"trigger_template,omitempty"`
}

type CloudbuildTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudbuildTriggerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudbuildTriggerList is a list of CloudbuildTriggers
type CloudbuildTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudbuildTrigger CRD objects
	Items []CloudbuildTrigger `json:"items,omitempty"`
}
