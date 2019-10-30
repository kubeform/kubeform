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

type LambdaPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaPermissionSpec   `json:"spec,omitempty"`
	Status            LambdaPermissionStatus `json:"status,omitempty"`
}

type LambdaPermissionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action string `json:"action" tf:"action"`
	// +optional
	EventSourceToken string `json:"eventSourceToken,omitempty" tf:"event_source_token,omitempty"`
	FunctionName     string `json:"functionName" tf:"function_name"`
	Principal        string `json:"principal" tf:"principal"`
	// +optional
	Qualifier string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
	// +optional
	SourceAccount string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`
	// +optional
	SourceArn string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
	// +optional
	StatementID string `json:"statementID,omitempty" tf:"statement_id,omitempty"`
	// +optional
	StatementIDPrefix string `json:"statementIDPrefix,omitempty" tf:"statement_id_prefix,omitempty"`
}

type LambdaPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LambdaPermissionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaPermissionList is a list of LambdaPermissions
type LambdaPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaPermission CRD objects
	Items []LambdaPermission `json:"items,omitempty"`
}
