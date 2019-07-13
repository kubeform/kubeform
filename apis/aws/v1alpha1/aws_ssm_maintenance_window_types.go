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

type AwsSsmMaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmMaintenanceWindowSpec   `json:"spec,omitempty"`
	Status            AwsSsmMaintenanceWindowStatus `json:"status,omitempty"`
}

type AwsSsmMaintenanceWindowSpec struct {
	StartDate                string            `json:"start_date"`
	Cutoff                   int               `json:"cutoff"`
	AllowUnassociatedTargets bool              `json:"allow_unassociated_targets"`
	Duration                 int               `json:"duration"`
	Enabled                  bool              `json:"enabled"`
	EndDate                  string            `json:"end_date"`
	ScheduleTimezone         string            `json:"schedule_timezone"`
	Tags                     map[string]string `json:"tags"`
	Name                     string            `json:"name"`
	Schedule                 string            `json:"schedule"`
}



type AwsSsmMaintenanceWindowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmMaintenanceWindowList is a list of AwsSsmMaintenanceWindows
type AwsSsmMaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmMaintenanceWindow CRD objects
	Items []AwsSsmMaintenanceWindow `json:"items,omitempty"`
}