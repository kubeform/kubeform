package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitorAutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorAutoscaleSettingSpec   `json:"spec,omitempty"`
	Status            MonitorAutoscaleSettingStatus `json:"status,omitempty"`
}

type MonitorAutoscaleSettingSpecNotificationEmail struct {
	// +optional
	CustomEmails []string `json:"custom_emails,omitempty"`
	// +optional
	SendToSubscriptionAdministrator bool `json:"send_to_subscription_administrator,omitempty"`
	// +optional
	SendToSubscriptionCoAdministrator bool `json:"send_to_subscription_co_administrator,omitempty"`
}

type MonitorAutoscaleSettingSpecNotificationWebhook struct {
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	ServiceUri string            `json:"service_uri"`
}

type MonitorAutoscaleSettingSpecNotification struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Email *[]MonitorAutoscaleSettingSpecNotification `json:"email,omitempty"`
	// +optional
	Webhook *[]MonitorAutoscaleSettingSpecNotification `json:"webhook,omitempty"`
}

type MonitorAutoscaleSettingSpecProfileCapacity struct {
	Default int `json:"default"`
	Maximum int `json:"maximum"`
	Minimum int `json:"minimum"`
}

type MonitorAutoscaleSettingSpecProfileFixedDate struct {
	End   string `json:"end"`
	Start string `json:"start"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
}

type MonitorAutoscaleSettingSpecProfileRecurrence struct {
	Days []string `json:"days"`
	// +kubebuilder:validation:MaxItems=1
	Hours []int64 `json:"hours"`
	// +kubebuilder:validation:MaxItems=1
	Minutes []int64 `json:"minutes"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
}

type MonitorAutoscaleSettingSpecProfileRuleMetricTrigger struct {
	MetricName       string      `json:"metric_name"`
	MetricResourceId string      `json:"metric_resource_id"`
	Operator         string      `json:"operator"`
	Statistic        string      `json:"statistic"`
	Threshold        json.Number `json:"threshold"`
	TimeAggregation  string      `json:"time_aggregation"`
	TimeGrain        string      `json:"time_grain"`
	TimeWindow       string      `json:"time_window"`
}

type MonitorAutoscaleSettingSpecProfileRuleScaleAction struct {
	Cooldown  string `json:"cooldown"`
	Direction string `json:"direction"`
	Type      string `json:"type"`
	Value     int    `json:"value"`
}

type MonitorAutoscaleSettingSpecProfileRule struct {
	// +kubebuilder:validation:MaxItems=1
	MetricTrigger []MonitorAutoscaleSettingSpecProfileRule `json:"metric_trigger"`
	// +kubebuilder:validation:MaxItems=1
	ScaleAction []MonitorAutoscaleSettingSpecProfileRule `json:"scale_action"`
}

type MonitorAutoscaleSettingSpecProfile struct {
	// +kubebuilder:validation:MaxItems=1
	Capacity []MonitorAutoscaleSettingSpecProfile `json:"capacity"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedDate *[]MonitorAutoscaleSettingSpecProfile `json:"fixed_date,omitempty"`
	Name      string                                `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Recurrence *[]MonitorAutoscaleSettingSpecProfile `json:"recurrence,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Rule *[]MonitorAutoscaleSettingSpecProfile `json:"rule,omitempty"`
}

type MonitorAutoscaleSettingSpec struct {
	// +optional
	Enabled  bool   `json:"enabled,omitempty"`
	Location string `json:"location"`
	Name     string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Notification *[]MonitorAutoscaleSettingSpec `json:"notification,omitempty"`
	// +kubebuilder:validation:MaxItems=20
	Profile           []MonitorAutoscaleSettingSpec `json:"profile"`
	ResourceGroupName string                        `json:"resource_group_name"`
	TargetResourceId  string                        `json:"target_resource_id"`
}

type MonitorAutoscaleSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorAutoscaleSettingList is a list of MonitorAutoscaleSettings
type MonitorAutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorAutoscaleSetting CRD objects
	Items []MonitorAutoscaleSetting `json:"items,omitempty"`
}
