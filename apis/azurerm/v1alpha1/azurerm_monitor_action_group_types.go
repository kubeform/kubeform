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

type AzurermMonitorActionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorActionGroupSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorActionGroupStatus `json:"status,omitempty"`
}

type AzurermMonitorActionGroupSpecEmailReceiver struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

type AzurermMonitorActionGroupSpecSmsReceiver struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}

type AzurermMonitorActionGroupSpecWebhookReceiver struct {
	Name       string `json:"name"`
	ServiceUri string `json:"service_uri"`
}

type AzurermMonitorActionGroupSpec struct {
	Tags              map[string]string               `json:"tags"`
	Name              string                          `json:"name"`
	ResourceGroupName string                          `json:"resource_group_name"`
	ShortName         string                          `json:"short_name"`
	Enabled           bool                            `json:"enabled"`
	EmailReceiver     []AzurermMonitorActionGroupSpec `json:"email_receiver"`
	SmsReceiver       []AzurermMonitorActionGroupSpec `json:"sms_receiver"`
	WebhookReceiver   []AzurermMonitorActionGroupSpec `json:"webhook_receiver"`
}

type AzurermMonitorActionGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMonitorActionGroupList is a list of AzurermMonitorActionGroups
type AzurermMonitorActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorActionGroup CRD objects
	Items []AzurermMonitorActionGroup `json:"items,omitempty"`
}
