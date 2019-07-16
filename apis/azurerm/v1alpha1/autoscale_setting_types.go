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

type AutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscaleSettingSpec   `json:"spec,omitempty"`
	Status            AutoscaleSettingStatus `json:"status,omitempty"`
}

type AutoscaleSettingSpecNotificationEmail struct {
	// +optional
	CustomEmails []string `json:"custom_emails,omitempty"`
	// +optional
	SendToSubscriptionAdministrator bool `json:"send_to_subscription_administrator,omitempty"`
	// +optional
	SendToSubscriptionCoAdministrator bool `json:"send_to_subscription_co_administrator,omitempty"`
}

type AutoscaleSettingSpecNotificationWebhook struct {
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	ServiceUri string            `json:"service_uri"`
}

type AutoscaleSettingSpecNotification struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Email *[]AutoscaleSettingSpecNotification `json:"email,omitempty"`
	// +optional
	Webhook *[]AutoscaleSettingSpecNotification `json:"webhook,omitempty"`
}

type AutoscaleSettingSpecProfileCapacity struct {
	Default int `json:"default"`
	Maximum int `json:"maximum"`
	Minimum int `json:"minimum"`
}

type AutoscaleSettingSpecProfileFixedDate struct {
	End   string `json:"end"`
	Start string `json:"start"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
}

type AutoscaleSettingSpecProfileRecurrence struct {
	Days []string `json:"days"`
	// +kubebuilder:validation:MaxItems=1
	Hours []int64 `json:"hours"`
	// +kubebuilder:validation:MaxItems=1
	Minutes []int64 `json:"minutes"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
}

type AutoscaleSettingSpecProfileRuleMetricTrigger struct {
	MetricName       string      `json:"metric_name"`
	MetricResourceId string      `json:"metric_resource_id"`
	Operator         string      `json:"operator"`
	Statistic        string      `json:"statistic"`
	Threshold        json.Number `json:"threshold"`
	TimeAggregation  string      `json:"time_aggregation"`
	TimeGrain        string      `json:"time_grain"`
	TimeWindow       string      `json:"time_window"`
}

type AutoscaleSettingSpecProfileRuleScaleAction struct {
	Cooldown  string `json:"cooldown"`
	Direction string `json:"direction"`
	Type      string `json:"type"`
	Value     int    `json:"value"`
}

type AutoscaleSettingSpecProfileRule struct {
	// +kubebuilder:validation:MaxItems=1
	MetricTrigger []AutoscaleSettingSpecProfileRule `json:"metric_trigger"`
	// +kubebuilder:validation:MaxItems=1
	ScaleAction []AutoscaleSettingSpecProfileRule `json:"scale_action"`
}

type AutoscaleSettingSpecProfile struct {
	// +kubebuilder:validation:MaxItems=1
	Capacity []AutoscaleSettingSpecProfile `json:"capacity"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedDate *[]AutoscaleSettingSpecProfile `json:"fixed_date,omitempty"`
	Name      string                         `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Recurrence *[]AutoscaleSettingSpecProfile `json:"recurrence,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Rule *[]AutoscaleSettingSpecProfile `json:"rule,omitempty"`
}

type AutoscaleSettingSpec struct {
	// +optional
	Enabled  bool   `json:"enabled,omitempty"`
	Location string `json:"location"`
	Name     string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Notification *[]AutoscaleSettingSpec `json:"notification,omitempty"`
	// +kubebuilder:validation:MaxItems=20
	Profile           []AutoscaleSettingSpec `json:"profile"`
	ResourceGroupName string                 `json:"resource_group_name"`
	TargetResourceId  string                 `json:"target_resource_id"`
}

type AutoscaleSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscaleSettingList is a list of AutoscaleSettings
type AutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscaleSetting CRD objects
	Items []AutoscaleSetting `json:"items,omitempty"`
}
