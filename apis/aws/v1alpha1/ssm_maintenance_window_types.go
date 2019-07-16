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

type SsmMaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowSpec   `json:"spec,omitempty"`
	Status            SsmMaintenanceWindowStatus `json:"status,omitempty"`
}

type SsmMaintenanceWindowSpec struct {
	// +optional
	AllowUnassociatedTargets bool `json:"allow_unassociated_targets,omitempty"`
	Cutoff                   int  `json:"cutoff"`
	Duration                 int  `json:"duration"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	EndDate  string `json:"end_date,omitempty"`
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	// +optional
	ScheduleTimezone string `json:"schedule_timezone,omitempty"`
	// +optional
	StartDate string `json:"start_date,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SsmMaintenanceWindowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmMaintenanceWindowList is a list of SsmMaintenanceWindows
type SsmMaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmMaintenanceWindow CRD objects
	Items []SsmMaintenanceWindow `json:"items,omitempty"`
}
