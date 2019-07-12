package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAutomationSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationScheduleSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationScheduleStatus `json:"status,omitempty"`
}

type AzurermAutomationScheduleSpecMonthlyOccurrence struct {
	Day        string `json:"day"`
	Occurrence int    `json:"occurrence"`
}

type AzurermAutomationScheduleSpec struct {
	MonthlyOccurrence     []AzurermAutomationScheduleSpec `json:"monthly_occurrence"`
	Name                  string                          `json:"name"`
	AutomationAccountName string                          `json:"automation_account_name"`
	Frequency             string                          `json:"frequency"`
	Description           string                          `json:"description"`
	Timezone              string                          `json:"timezone"`
	WeekDays              []string                        `json:"week_days"`
	ResourceGroupName     string                          `json:"resource_group_name"`
	AccountName           string                          `json:"account_name"`
	Interval              int                             `json:"interval"`
	StartTime             string                          `json:"start_time"`
	ExpiryTime            string                          `json:"expiry_time"`
	MonthDays             []int64                         `json:"month_days"`
}

type AzurermAutomationScheduleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAutomationScheduleList is a list of AzurermAutomationSchedules
type AzurermAutomationScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationSchedule CRD objects
	Items []AzurermAutomationSchedule `json:"items,omitempty"`
}
