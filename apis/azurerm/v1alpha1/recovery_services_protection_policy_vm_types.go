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

type RecoveryServicesProtectionPolicyVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryServicesProtectionPolicyVmSpec   `json:"spec,omitempty"`
	Status            RecoveryServicesProtectionPolicyVmStatus `json:"status,omitempty"`
}

type RecoveryServicesProtectionPolicyVmSpecBackup struct {
	Frequency string `json:"frequency" tf:"frequency"`
	Time      string `json:"time" tf:"time"`
	// +optional
	Weekdays []string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionDaily struct {
	Count int `json:"count" tf:"count"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionMonthly struct {
	Count    int      `json:"count" tf:"count"`
	Weekdays []string `json:"weekdays" tf:"weekdays"`
	Weeks    []string `json:"weeks" tf:"weeks"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionWeekly struct {
	Count    int      `json:"count" tf:"count"`
	Weekdays []string `json:"weekdays" tf:"weekdays"`
}

type RecoveryServicesProtectionPolicyVmSpecRetentionYearly struct {
	Count    int      `json:"count" tf:"count"`
	Months   []string `json:"months" tf:"months"`
	Weekdays []string `json:"weekdays" tf:"weekdays"`
	Weeks    []string `json:"weeks" tf:"weeks"`
}

type RecoveryServicesProtectionPolicyVmSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	Backup            []RecoveryServicesProtectionPolicyVmSpecBackup `json:"backup" tf:"backup"`
	Name              string                                         `json:"name" tf:"name"`
	RecoveryVaultName string                                         `json:"recoveryVaultName" tf:"recovery_vault_name"`
	ResourceGroupName string                                         `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionDaily []RecoveryServicesProtectionPolicyVmSpecRetentionDaily `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionMonthly []RecoveryServicesProtectionPolicyVmSpecRetentionMonthly `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionWeekly []RecoveryServicesProtectionPolicyVmSpecRetentionWeekly `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionYearly []RecoveryServicesProtectionPolicyVmSpecRetentionYearly `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type RecoveryServicesProtectionPolicyVmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RecoveryServicesProtectionPolicyVmSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
