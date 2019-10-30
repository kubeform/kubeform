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

type MonitorActionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionGroupSpec   `json:"spec,omitempty"`
	Status            MonitorActionGroupStatus `json:"status,omitempty"`
}

type MonitorActionGroupSpecEmailReceiver struct {
	EmailAddress string `json:"emailAddress" tf:"email_address"`
	Name         string `json:"name" tf:"name"`
}

type MonitorActionGroupSpecSmsReceiver struct {
	CountryCode string `json:"countryCode" tf:"country_code"`
	Name        string `json:"name" tf:"name"`
	PhoneNumber string `json:"phoneNumber" tf:"phone_number"`
}

type MonitorActionGroupSpecWebhookReceiver struct {
	Name       string `json:"name" tf:"name"`
	ServiceURI string `json:"serviceURI" tf:"service_uri"`
}

type MonitorActionGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EmailReceiver []MonitorActionGroupSpecEmailReceiver `json:"emailReceiver,omitempty" tf:"email_receiver,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ShortName         string `json:"shortName" tf:"short_name"`
	// +optional
	SmsReceiver []MonitorActionGroupSpecSmsReceiver `json:"smsReceiver,omitempty" tf:"sms_receiver,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	WebhookReceiver []MonitorActionGroupSpecWebhookReceiver `json:"webhookReceiver,omitempty" tf:"webhook_receiver,omitempty"`
}

type MonitorActionGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorActionGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorActionGroupList is a list of MonitorActionGroups
type MonitorActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorActionGroup CRD objects
	Items []MonitorActionGroup `json:"items,omitempty"`
}
