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

type AwsBudgetsBudgetSpecNotification struct {
	NotificationType         string   `json:"notification_type"`
	SubscriberEmailAddresses []string `json:"subscriber_email_addresses"`
	SubscriberSnsTopicArns   []string `json:"subscriber_sns_topic_arns"`
	ComparisonOperator       string   `json:"comparison_operator"`
	Threshold                float64  `json:"threshold"`
	ThresholdType            string   `json:"threshold_type"`
}

type AwsBudgetsBudgetSpecCostTypes struct {
	IncludeRecurring         bool `json:"include_recurring"`
	IncludeUpfront           bool `json:"include_upfront"`
	UseBlended               bool `json:"use_blended"`
	IncludeRefund            bool `json:"include_refund"`
	IncludeSubscription      bool `json:"include_subscription"`
	IncludeSupport           bool `json:"include_support"`
	IncludeTax               bool `json:"include_tax"`
	UseAmortized             bool `json:"use_amortized"`
	IncludeCredit            bool `json:"include_credit"`
	IncludeDiscount          bool `json:"include_discount"`
	IncludeOtherSubscription bool `json:"include_other_subscription"`
}

type AwsBudgetsBudgetSpec struct {
	TimeUnit        string                 `json:"time_unit"`
	CostFilters     map[string]string      `json:"cost_filters"`
	Notification    []AwsBudgetsBudgetSpec `json:"notification"`
	Name            string                 `json:"name"`
	NamePrefix      string                 `json:"name_prefix"`
	LimitUnit       string                 `json:"limit_unit"`
	CostTypes       []AwsBudgetsBudgetSpec `json:"cost_types"`
	TimePeriodStart string                 `json:"time_period_start"`
	AccountId       string                 `json:"account_id"`
	BudgetType      string                 `json:"budget_type"`
	LimitAmount     string                 `json:"limit_amount"`
	TimePeriodEnd   string                 `json:"time_period_end"`
}



type AwsBudgetsBudgetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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