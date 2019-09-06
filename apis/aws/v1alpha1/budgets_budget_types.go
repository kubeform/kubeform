package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type BudgetsBudget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetsBudgetSpec   `json:"spec,omitempty"`
	Status            BudgetsBudgetStatus `json:"status,omitempty"`
}

type BudgetsBudgetSpecCostTypes struct {
	// +optional
	IncludeCredit bool `json:"includeCredit,omitempty" tf:"include_credit,omitempty"`
	// +optional
	IncludeDiscount bool `json:"includeDiscount,omitempty" tf:"include_discount,omitempty"`
	// +optional
	IncludeOtherSubscription bool `json:"includeOtherSubscription,omitempty" tf:"include_other_subscription,omitempty"`
	// +optional
	IncludeRecurring bool `json:"includeRecurring,omitempty" tf:"include_recurring,omitempty"`
	// +optional
	IncludeRefund bool `json:"includeRefund,omitempty" tf:"include_refund,omitempty"`
	// +optional
	IncludeSubscription bool `json:"includeSubscription,omitempty" tf:"include_subscription,omitempty"`
	// +optional
	IncludeSupport bool `json:"includeSupport,omitempty" tf:"include_support,omitempty"`
	// +optional
	IncludeTax bool `json:"includeTax,omitempty" tf:"include_tax,omitempty"`
	// +optional
	IncludeUpfront bool `json:"includeUpfront,omitempty" tf:"include_upfront,omitempty"`
	// +optional
	UseAmortized bool `json:"useAmortized,omitempty" tf:"use_amortized,omitempty"`
	// +optional
	UseBlended bool `json:"useBlended,omitempty" tf:"use_blended,omitempty"`
}

type BudgetsBudgetSpecNotification struct {
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`
	NotificationType   string `json:"notificationType" tf:"notification_type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriberEmailAddresses []string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriberSnsTopicArns []string    `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns,omitempty"`
	Threshold              json.Number `json:"threshold" tf:"threshold"`
	ThresholdType          string      `json:"thresholdType" tf:"threshold_type"`
}

type BudgetsBudgetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID  string `json:"accountID,omitempty" tf:"account_id,omitempty"`
	BudgetType string `json:"budgetType" tf:"budget_type"`
	// +optional
	CostFilters map[string]string `json:"costFilters,omitempty" tf:"cost_filters,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CostTypes   []BudgetsBudgetSpecCostTypes `json:"costTypes,omitempty" tf:"cost_types,omitempty"`
	LimitAmount string                       `json:"limitAmount" tf:"limit_amount"`
	LimitUnit   string                       `json:"limitUnit" tf:"limit_unit"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Notification []BudgetsBudgetSpecNotification `json:"notification,omitempty" tf:"notification,omitempty"`
	// +optional
	TimePeriodEnd   string `json:"timePeriodEnd,omitempty" tf:"time_period_end,omitempty"`
	TimePeriodStart string `json:"timePeriodStart" tf:"time_period_start"`
	TimeUnit        string `json:"timeUnit" tf:"time_unit"`
}

type BudgetsBudgetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BudgetsBudgetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BudgetsBudgetList is a list of BudgetsBudgets
type BudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BudgetsBudget CRD objects
	Items []BudgetsBudget `json:"items,omitempty"`
}
