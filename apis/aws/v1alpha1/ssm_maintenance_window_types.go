package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AllowUnassociatedTargets bool `json:"allowUnassociatedTargets,omitempty" tf:"allow_unassociated_targets,omitempty"`
	Cutoff                   int  `json:"cutoff" tf:"cutoff"`
	Duration                 int  `json:"duration" tf:"duration"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	EndDate  string `json:"endDate,omitempty" tf:"end_date,omitempty"`
	Name     string `json:"name" tf:"name"`
	Schedule string `json:"schedule" tf:"schedule"`
	// +optional
	ScheduleTimezone string `json:"scheduleTimezone,omitempty" tf:"schedule_timezone,omitempty"`
	// +optional
	StartDate string `json:"startDate,omitempty" tf:"start_date,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SsmMaintenanceWindowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmMaintenanceWindowSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
