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

type SqlDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status            SqlDatabaseInstanceStatus `json:"status,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Value string `json:"value,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	// +optional
	Day int `json:"day,omitempty"`
	// +optional
	Hour int `json:"hour,omitempty"`
	// +optional
	UpdateTrack string `json:"update_track,omitempty"`
}

type SqlDatabaseInstanceSpecSettings struct {
	// +optional
	AuthorizedGaeApplications []string `json:"authorized_gae_applications,omitempty"`
	// +optional
	DatabaseFlags *[]SqlDatabaseInstanceSpecSettings `json:"database_flags,omitempty"`
	// +optional
	DiskAutoresize bool `json:"disk_autoresize,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenanceWindow *[]SqlDatabaseInstanceSpecSettings `json:"maintenance_window,omitempty"`
	// +optional
	PricingPlan string `json:"pricing_plan,omitempty"`
	// +optional
	ReplicationType string `json:"replication_type,omitempty"`
	Tier            string `json:"tier"`
	// +optional
	UserLabels map[string]string `json:"user_labels,omitempty"`
}

type SqlDatabaseInstanceSpec struct {
	// +optional
	DatabaseVersion string `json:"database_version,omitempty"`
	// +optional
	Region string `json:"region,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Settings []SqlDatabaseInstanceSpec `json:"settings"`
}

type SqlDatabaseInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseInstanceList is a list of SqlDatabaseInstances
type SqlDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlDatabaseInstance CRD objects
	Items []SqlDatabaseInstance `json:"items,omitempty"`
}
