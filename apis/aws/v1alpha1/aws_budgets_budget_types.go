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

type AwsBudgetsBudget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBudgetsBudgetSpec   `json:"spec,omitempty"`
	Status            AwsBudgetsBudgetStatus `json:"status,omitempty"`
}

type AwsBudgetsBudgetSpecCostTypes struct {
	IncludeRecurring         bool `json:"include_recurring"`
	IncludeSubscription      bool `json:"include_subscription"`
	IncludeSupport           bool `json:"include_support"`
	IncludeUpfront           bool `json:"include_upfront"`
	IncludeCredit            bool `json:"include_credit"`
	IncludeDiscount          bool `json:"include_discount"`
	IncludeOtherSubscription bool `json:"include_other_subscription"`
	UseBlended               bool `json:"use_blended"`
	IncludeRefund            bool `json:"include_refund"`
	IncludeTax               bool `json:"include_tax"`
	UseAmortized             bool `json:"use_amortized"`
}

type AwsBudgetsBudgetSpecNotification struct {
	ThresholdType            string   `json:"threshold_type"`
	NotificationType         string   `json:"notification_type"`
	SubscriberEmailAddresses []string `json:"subscriber_email_addresses"`
	SubscriberSnsTopicArns   []string `json:"subscriber_sns_topic_arns"`
	ComparisonOperator       string   `json:"comparison_operator"`
	Threshold                float64  `json:"threshold"`
}

type AwsBudgetsBudgetSpec struct {
	TimePeriodEnd   string                 `json:"time_period_end"`
	CostTypes       []AwsBudgetsBudgetSpec `json:"cost_types"`
	TimePeriodStart string                 `json:"time_period_start"`
	NamePrefix      string                 `json:"name_prefix"`
	BudgetType      string                 `json:"budget_type"`
	LimitAmount     string                 `json:"limit_amount"`
	LimitUnit       string                 `json:"limit_unit"`
	TimeUnit        string                 `json:"time_unit"`
	CostFilters     map[string]string      `json:"cost_filters"`
	AccountId       string                 `json:"account_id"`
	Name            string                 `json:"name"`
	Notification    []AwsBudgetsBudgetSpec `json:"notification"`
}

type AwsBudgetsBudgetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsBudgetsBudgetList is a list of AwsBudgetsBudgets
type AwsBudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBudgetsBudget CRD objects
	Items []AwsBudgetsBudget `json:"items,omitempty"`
}
