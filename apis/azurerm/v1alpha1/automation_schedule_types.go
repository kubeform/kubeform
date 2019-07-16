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

type AutomationSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationScheduleSpec   `json:"spec,omitempty"`
	Status            AutomationScheduleStatus `json:"status,omitempty"`
}

type AutomationScheduleSpecMonthlyOccurrence struct {
	Day        string `json:"day"`
	Occurrence int    `json:"occurrence"`
}

type AutomationScheduleSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Frequency   string `json:"frequency"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	MonthDays []int64 `json:"month_days,omitempty"`
	// +optional
	MonthlyOccurrence *[]AutomationScheduleSpec `json:"monthly_occurrence,omitempty"`
	Name              string                    `json:"name"`
	ResourceGroupName string                    `json:"resource_group_name"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WeekDays []string `json:"week_days,omitempty"`
}

type AutomationScheduleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
