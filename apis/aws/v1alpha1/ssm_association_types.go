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

type SsmAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmAssociationSpec   `json:"spec,omitempty"`
	Status            SsmAssociationStatus `json:"status,omitempty"`
}

type SsmAssociationSpecOutputLocation struct {
	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`
	// +optional
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type SsmAssociationSpecTargets struct {
	Key    string   `json:"key" tf:"key"`
	Values []string `json:"values" tf:"values"`
}

type SsmAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssociationID string `json:"associationID,omitempty" tf:"association_id,omitempty"`
	// +optional
	AssociationName string `json:"associationName,omitempty" tf:"association_name,omitempty"`
	// +optional
	ComplianceSeverity string `json:"complianceSeverity,omitempty" tf:"compliance_severity,omitempty"`
	// +optional
	DocumentVersion string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`
	// +optional
	InstanceID string `json:"instanceID,omitempty" tf:"instance_id,omitempty"`
	// +optional
	MaxConcurrency string `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`
	// +optional
	MaxErrors string `json:"maxErrors,omitempty" tf:"max_errors,omitempty"`
	Name      string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OutputLocation []SsmAssociationSpecOutputLocation `json:"outputLocation,omitempty" tf:"output_location,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	ScheduleExpression string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Targets []SsmAssociationSpecTargets `json:"targets,omitempty" tf:"targets,omitempty"`
}

type SsmAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmAssociationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmAssociationList is a list of SsmAssociations
type SsmAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmAssociation CRD objects
	Items []SsmAssociation `json:"items,omitempty"`
}
