package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AutomationSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationScheduleSpec   `json:"spec,omitempty"`
	Status            AutomationScheduleStatus `json:"status,omitempty"`
}

type AutomationScheduleSpecMonthlyOccurrence struct {
	Day        string `json:"day" tf:"day"`
	Occurrence int    `json:"occurrence" tf:"occurrence"`
}

type AutomationScheduleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	AccountName string `json:"accountName,omitempty" tf:"account_name,omitempty"`
	// +optional
	AutomationAccountName string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ExpiryTime string `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`
	Frequency  string `json:"frequency" tf:"frequency"`
	// +optional
	Interval int `json:"interval,omitempty" tf:"interval,omitempty"`
	// +optional
	MonthDays []int64 `json:"monthDays,omitempty" tf:"month_days,omitempty"`
	// +optional
	MonthlyOccurrence []AutomationScheduleSpecMonthlyOccurrence `json:"monthlyOccurrence,omitempty" tf:"monthly_occurrence,omitempty"`
	Name              string                                    `json:"name" tf:"name"`
	ResourceGroupName string                                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
	// +optional
	WeekDays []string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

type AutomationScheduleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutomationScheduleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationScheduleList is a list of AutomationSchedules
type AutomationScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationSchedule CRD objects
	Items []AutomationSchedule `json:"items,omitempty"`
}
