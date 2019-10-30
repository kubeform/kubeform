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

type ServicebusSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusSubscriptionSpec   `json:"spec,omitempty"`
	Status            ServicebusSubscriptionStatus `json:"status,omitempty"`
}

type ServicebusSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoDeleteOnIdle string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`
	// +optional
	// Deprecated
	DeadLetteringOnFilterEvaluationExceptions bool `json:"deadLetteringOnFilterEvaluationExceptions,omitempty" tf:"dead_lettering_on_filter_evaluation_exceptions,omitempty"`
	// +optional
	DeadLetteringOnMessageExpiration bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`
	// +optional
	DefaultMessageTtl string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`
	// +optional
	EnableBatchedOperations bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`
	// +optional
	ForwardTo string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	LockDuration     string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`
	MaxDeliveryCount int    `json:"maxDeliveryCount" tf:"max_delivery_count"`
	Name             string `json:"name" tf:"name"`
	NamespaceName    string `json:"namespaceName" tf:"namespace_name"`
	// +optional
	RequiresSession   bool   `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	TopicName         string `json:"topicName" tf:"topic_name"`
}

type ServicebusSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServicebusSubscriptionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusSubscriptionList is a list of ServicebusSubscriptions
type ServicebusSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusSubscription CRD objects
	Items []ServicebusSubscription `json:"items,omitempty"`
}
