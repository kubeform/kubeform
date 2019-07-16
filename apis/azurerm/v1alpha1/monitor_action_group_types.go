package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitorActionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionGroupSpec   `json:"spec,omitempty"`
	Status            MonitorActionGroupStatus `json:"status,omitempty"`
}

type MonitorActionGroupSpecEmailReceiver struct {
	EmailAddress string `json:"email_address"`
	Name         string `json:"name"`
}

type MonitorActionGroupSpecSmsReceiver struct {
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type MonitorActionGroupSpecWebhookReceiver struct {
	Name       string `json:"name"`
	ServiceUri string `json:"service_uri"`
}

type MonitorActionGroupSpec struct {
	// +optional
	EmailReceiver *[]MonitorActionGroupSpec `json:"email_receiver,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ShortName         string `json:"short_name"`
	// +optional
	SmsReceiver *[]MonitorActionGroupSpec `json:"sms_receiver,omitempty"`
	// +optional
	WebhookReceiver *[]MonitorActionGroupSpec `json:"webhook_receiver,omitempty"`
}

type MonitorActionGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
