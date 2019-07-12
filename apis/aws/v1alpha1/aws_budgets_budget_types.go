package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBudgetsBudget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBudgetsBudgetSpec   `json:"spec,omitempty"`
	Status            AwsBudgetsBudgetStatus `json:"status,omitempty"`
}

type AwsBudgetsBudgetSpecCostTypes struct {
	IncludeSubscription      bool `json:"include_subscription"`
	IncludeUpfront           bool `json:"include_upfront"`
	UseAmortized             bool `json:"use_amortized"`
	UseBlended               bool `json:"use_blended"`
	IncludeCredit            bool `json:"include_credit"`
	IncludeRecurring         bool `json:"include_recurring"`
	IncludeRefund            bool `json:"include_refund"`
	IncludeSupport           bool `json:"include_support"`
	IncludeTax               bool `json:"include_tax"`
	IncludeDiscount          bool `json:"include_discount"`
	IncludeOtherSubscription bool `json:"include_other_subscription"`
}

type AwsBudgetsBudgetSpecNotification struct {
	ComparisonOperator       string   `json:"comparison_operator"`
	Threshold                float64  `json:"threshold"`
	ThresholdType            string   `json:"threshold_type"`
	NotificationType         string   `json:"notification_type"`
	SubscriberEmailAddresses []string `json:"subscriber_email_addresses"`
	SubscriberSnsTopicArns   []string `json:"subscriber_sns_topic_arns"`
}

type AwsBudgetsBudgetSpec struct {
	CostFilters     map[string]string      `json:"cost_filters"`
	AccountId       string                 `json:"account_id"`
	NamePrefix      string                 `json:"name_prefix"`
	BudgetType      string                 `json:"budget_type"`
	LimitUnit       string                 `json:"limit_unit"`
	CostTypes       []AwsBudgetsBudgetSpec `json:"cost_types"`
	TimePeriodStart string                 `json:"time_period_start"`
	Name            string                 `json:"name"`
	LimitAmount     string                 `json:"limit_amount"`
	TimePeriodEnd   string                 `json:"time_period_end"`
	TimeUnit        string                 `json:"time_unit"`
	Notification    []AwsBudgetsBudgetSpec `json:"notification"`
}

type AwsBudgetsBudgetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBudgetsBudgetList is a list of AwsBudgetsBudgets
type AwsBudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBudgetsBudget CRD objects
	Items []AwsBudgetsBudget `json:"items,omitempty"`
}
