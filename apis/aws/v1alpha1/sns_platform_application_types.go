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

type SnsPlatformApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsPlatformApplicationSpec   `json:"spec,omitempty"`
	Status            SnsPlatformApplicationStatus `json:"status,omitempty"`
}

type SnsPlatformApplicationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	EventDeliveryFailureTopicArn string `json:"eventDeliveryFailureTopicArn,omitempty" tf:"event_delivery_failure_topic_arn,omitempty"`
	// +optional
	EventEndpointCreatedTopicArn string `json:"eventEndpointCreatedTopicArn,omitempty" tf:"event_endpoint_created_topic_arn,omitempty"`
	// +optional
	EventEndpointDeletedTopicArn string `json:"eventEndpointDeletedTopicArn,omitempty" tf:"event_endpoint_deleted_topic_arn,omitempty"`
	// +optional
	EventEndpointUpdatedTopicArn string `json:"eventEndpointUpdatedTopicArn,omitempty" tf:"event_endpoint_updated_topic_arn,omitempty"`
	// +optional
	FailureFeedbackRoleArn string `json:"failureFeedbackRoleArn,omitempty" tf:"failure_feedback_role_arn,omitempty"`
	Name                   string `json:"name" tf:"name"`
	Platform               string `json:"platform" tf:"platform"`
	PlatformCredential     string `json:"platformCredential" tf:"platform_credential"`
	// +optional
	PlatformPrincipal string `json:"platformPrincipal,omitempty" tf:"platform_principal,omitempty"`
	// +optional
	SuccessFeedbackRoleArn string `json:"successFeedbackRoleArn,omitempty" tf:"success_feedback_role_arn,omitempty"`
	// +optional
	SuccessFeedbackSampleRate string `json:"successFeedbackSampleRate,omitempty" tf:"success_feedback_sample_rate,omitempty"`
}

type SnsPlatformApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SnsPlatformApplicationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsPlatformApplicationList is a list of SnsPlatformApplications
type SnsPlatformApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsPlatformApplication CRD objects
	Items []SnsPlatformApplication `json:"items,omitempty"`
}
