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

type BudgetsBudget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetsBudgetSpec   `json:"spec,omitempty"`
	Status            BudgetsBudgetStatus `json:"status,omitempty"`
}

type BudgetsBudgetSpecNotification struct {
	ComparisonOperator string `json:"comparison_operator"`
	NotificationType   string `json:"notification_type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriberEmailAddresses []string `json:"subscriber_email_addresses,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriberSnsTopicArns []string    `json:"subscriber_sns_topic_arns,omitempty"`
	Threshold              json.Number `json:"threshold"`
	ThresholdType          string      `json:"threshold_type"`
}

type BudgetsBudgetSpec struct {
	BudgetType  string `json:"budget_type"`
	LimitAmount string `json:"limit_amount"`
	LimitUnit   string `json:"limit_unit"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Notification *[]BudgetsBudgetSpec `json:"notification,omitempty"`
	// +optional
	TimePeriodEnd   string `json:"time_period_end,omitempty"`
	TimePeriodStart string `json:"time_period_start"`
	TimeUnit        string `json:"time_unit"`
}

type BudgetsBudgetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
