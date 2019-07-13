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

type AzurermAutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutoscaleSettingSpec   `json:"spec,omitempty"`
	Status            AzurermAutoscaleSettingStatus `json:"status,omitempty"`
}

type AzurermAutoscaleSettingSpecProfileFixedDate struct {
	Timezone string `json:"timezone"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

type AzurermAutoscaleSettingSpecProfileRecurrence struct {
	Hours    []int64  `json:"hours"`
	Minutes  []int64  `json:"minutes"`
	Timezone string   `json:"timezone"`
	Days     []string `json:"days"`
}

type AzurermAutoscaleSettingSpecProfileCapacity struct {
	Minimum int `json:"minimum"`
	Maximum int `json:"maximum"`
	Default int `json:"default"`
}

type AzurermAutoscaleSettingSpecProfileRuleMetricTrigger struct {
	MetricResourceId string  `json:"metric_resource_id"`
	TimeGrain        string  `json:"time_grain"`
	Statistic        string  `json:"statistic"`
	TimeWindow       string  `json:"time_window"`
	TimeAggregation  string  `json:"time_aggregation"`
	Operator         string  `json:"operator"`
	Threshold        float64 `json:"threshold"`
	MetricName       string  `json:"metric_name"`
}

type AzurermAutoscaleSettingSpecProfileRuleScaleAction struct {
	Direction string `json:"direction"`
	Type      string `json:"type"`
	Value     int    `json:"value"`
	Cooldown  string `json:"cooldown"`
}

type AzurermAutoscaleSettingSpecProfileRule struct {
	MetricTrigger []AzurermAutoscaleSettingSpecProfileRule `json:"metric_trigger"`
	ScaleAction   []AzurermAutoscaleSettingSpecProfileRule `json:"scale_action"`
}

type AzurermAutoscaleSettingSpecProfile struct {
	FixedDate  []AzurermAutoscaleSettingSpecProfile `json:"fixed_date"`
	Recurrence []AzurermAutoscaleSettingSpecProfile `json:"recurrence"`
	Name       string                               `json:"name"`
	Capacity   []AzurermAutoscaleSettingSpecProfile `json:"capacity"`
	Rule       []AzurermAutoscaleSettingSpecProfile `json:"rule"`
}

type AzurermAutoscaleSettingSpecNotificationWebhook struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermAutoscaleSettingSpecNotificationEmail struct {
	SendToSubscriptionCoAdministrator bool     `json:"send_to_subscription_co_administrator"`
	CustomEmails                      []string `json:"custom_emails"`
	SendToSubscriptionAdministrator   bool     `json:"send_to_subscription_administrator"`
}

type AzurermAutoscaleSettingSpecNotification struct {
	Webhook []AzurermAutoscaleSettingSpecNotification `json:"webhook"`
	Email   []AzurermAutoscaleSettingSpecNotification `json:"email"`
}

type AzurermAutoscaleSettingSpec struct {
	Tags              map[string]string             `json:"tags"`
	Name              string                        `json:"name"`
	ResourceGroupName string                        `json:"resource_group_name"`
	Location          string                        `json:"location"`
	TargetResourceId  string                        `json:"target_resource_id"`
	Enabled           bool                          `json:"enabled"`
	Profile           []AzurermAutoscaleSettingSpec `json:"profile"`
	Notification      []AzurermAutoscaleSettingSpec `json:"notification"`
}



type AzurermAutoscaleSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutoscaleSettingList is a list of AzurermAutoscaleSettings
type AzurermAutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutoscaleSetting CRD objects
	Items []AzurermAutoscaleSetting `json:"items,omitempty"`
}