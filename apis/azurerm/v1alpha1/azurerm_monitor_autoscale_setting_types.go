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

type AzurermMonitorAutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorAutoscaleSettingSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorAutoscaleSettingStatus `json:"status,omitempty"`
}

type AzurermMonitorAutoscaleSettingSpecProfileCapacity struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
	Default int `json:"default"`
}

type AzurermMonitorAutoscaleSettingSpecProfileRuleMetricTrigger struct {
	Operator         string  `json:"operator"`
	Threshold        float64 `json:"threshold"`
	MetricName       string  `json:"metric_name"`
	MetricResourceId string  `json:"metric_resource_id"`
	TimeGrain        string  `json:"time_grain"`
	Statistic        string  `json:"statistic"`
	TimeWindow       string  `json:"time_window"`
	TimeAggregation  string  `json:"time_aggregation"`
}

type AzurermMonitorAutoscaleSettingSpecProfileRuleScaleAction struct {
	Direction string `json:"direction"`
	Type      string `json:"type"`
	Value     int    `json:"value"`
	Cooldown  string `json:"cooldown"`
}

type AzurermMonitorAutoscaleSettingSpecProfileRule struct {
	MetricTrigger []AzurermMonitorAutoscaleSettingSpecProfileRule `json:"metric_trigger"`
	ScaleAction   []AzurermMonitorAutoscaleSettingSpecProfileRule `json:"scale_action"`
}

type AzurermMonitorAutoscaleSettingSpecProfileFixedDate struct {
	Timezone string `json:"timezone"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

type AzurermMonitorAutoscaleSettingSpecProfileRecurrence struct {
	Timezone string   `json:"timezone"`
	Days     []string `json:"days"`
	Hours    []int64  `json:"hours"`
	Minutes  []int64  `json:"minutes"`
}

type AzurermMonitorAutoscaleSettingSpecProfile struct {
	Name       string                                      `json:"name"`
	Capacity   []AzurermMonitorAutoscaleSettingSpecProfile `json:"capacity"`
	Rule       []AzurermMonitorAutoscaleSettingSpecProfile `json:"rule"`
	FixedDate  []AzurermMonitorAutoscaleSettingSpecProfile `json:"fixed_date"`
	Recurrence []AzurermMonitorAutoscaleSettingSpecProfile `json:"recurrence"`
}

type AzurermMonitorAutoscaleSettingSpecNotificationEmail struct {
	SendToSubscriptionAdministrator   bool     `json:"send_to_subscription_administrator"`
	SendToSubscriptionCoAdministrator bool     `json:"send_to_subscription_co_administrator"`
	CustomEmails                      []string `json:"custom_emails"`
}

type AzurermMonitorAutoscaleSettingSpecNotificationWebhook struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermMonitorAutoscaleSettingSpecNotification struct {
	Email   []AzurermMonitorAutoscaleSettingSpecNotification `json:"email"`
	Webhook []AzurermMonitorAutoscaleSettingSpecNotification `json:"webhook"`
}

type AzurermMonitorAutoscaleSettingSpec struct {
	ResourceGroupName string                               `json:"resource_group_name"`
	Location          string                               `json:"location"`
	TargetResourceId  string                               `json:"target_resource_id"`
	Enabled           bool                                 `json:"enabled"`
	Profile           []AzurermMonitorAutoscaleSettingSpec `json:"profile"`
	Notification      []AzurermMonitorAutoscaleSettingSpec `json:"notification"`
	Tags              map[string]string                    `json:"tags"`
	Name              string                               `json:"name"`
}



type AzurermMonitorAutoscaleSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMonitorAutoscaleSettingList is a list of AzurermMonitorAutoscaleSettings
type AzurermMonitorAutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorAutoscaleSetting CRD objects
	Items []AzurermMonitorAutoscaleSetting `json:"items,omitempty"`
}