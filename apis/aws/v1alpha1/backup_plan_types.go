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

type BackupPlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPlanSpec   `json:"spec,omitempty"`
	Status            BackupPlanStatus `json:"status,omitempty"`
}

type BackupPlanSpecRuleLifecycle struct {
	// +optional
	ColdStorageAfter int `json:"cold_storage_after,omitempty"`
	// +optional
	DeleteAfter int `json:"delete_after,omitempty"`
}

type BackupPlanSpecRule struct {
	// +optional
	CompletionWindow int `json:"completion_window,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Lifecycle *[]BackupPlanSpecRule `json:"lifecycle,omitempty"`
	// +optional
	RecoveryPointTags map[string]string `json:"recovery_point_tags,omitempty"`
	RuleName          string            `json:"rule_name"`
	// +optional
	Schedule string `json:"schedule,omitempty"`
	// +optional
	StartWindow     int    `json:"start_window,omitempty"`
	TargetVaultName string `json:"target_vault_name"`
}

type BackupPlanSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Rule []BackupPlanSpec `json:"rule"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type BackupPlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BackupPlanList is a list of BackupPlans
type BackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BackupPlan CRD objects
	Items []BackupPlan `json:"items,omitempty"`
}
