package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
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
