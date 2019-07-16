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

type RecoveryServicesProtectionPolicyVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryServicesProtectionPolicyVmSpec   `json:"spec,omitempty"`
	Status            RecoveryServicesProtectionPolicyVmStatus `json:"status,omitempty"`
}

type RecoveryServicesProtectionPolicyVmSpecBackup struct {
	Frequency string `json:"frequency"`
	Time      string `json:"time"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Weekdays []string `json:"weekdays,omitempty"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionDaily struct {
	Count int `json:"count"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionMonthly struct {
	Count int `json:"count"`
	// +kubebuilder:validation:UniqueItems=true
	Weekdays []string `json:"weekdays"`
	// +kubebuilder:validation:UniqueItems=true
	Weeks []string `json:"weeks"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionWeekly struct {
	Count int `json:"count"`
	// +kubebuilder:validation:UniqueItems=true
	Weekdays []string `json:"weekdays"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionYearly struct {
	Count int `json:"count"`
	// +kubebuilder:validation:UniqueItems=true
	Months []string `json:"months"`
	// +kubebuilder:validation:UniqueItems=true
	Weekdays []string `json:"weekdays"`
	// +kubebuilder:validation:UniqueItems=true
	Weeks []string `json:"weeks"`
}

type RecoveryServicesProtectionPolicyVmSpec struct {
	// +kubebuilder:validation:MaxItems=1
	Backup            []RecoveryServicesProtectionPolicyVmSpec `json:"backup"`
	Name              string                                   `json:"name"`
	RecoveryVaultName string                                   `json:"recovery_vault_name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionDaily *[]RecoveryServicesProtectionPolicyVmSpec `json:"retention_daily,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionMonthly *[]RecoveryServicesProtectionPolicyVmSpec `json:"retention_monthly,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionWeekly *[]RecoveryServicesProtectionPolicyVmSpec `json:"retention_weekly,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionYearly *[]RecoveryServicesProtectionPolicyVmSpec `json:"retention_yearly,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty"`
}

type RecoveryServicesProtectionPolicyVmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RecoveryServicesProtectionPolicyVmList is a list of RecoveryServicesProtectionPolicyVms
type RecoveryServicesProtectionPolicyVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RecoveryServicesProtectionPolicyVm CRD objects
	Items []RecoveryServicesProtectionPolicyVm `json:"items,omitempty"`
}
