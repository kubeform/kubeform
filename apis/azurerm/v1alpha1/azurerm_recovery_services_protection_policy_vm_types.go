package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRecoveryServicesProtectionPolicyVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRecoveryServicesProtectionPolicyVmSpec   `json:"spec,omitempty"`
	Status            AzurermRecoveryServicesProtectionPolicyVmStatus `json:"status,omitempty"`
}

type AzurermRecoveryServicesProtectionPolicyVmSpecRetentionDaily struct {
	Count int `json:"count"`
}

type AzurermRecoveryServicesProtectionPolicyVmSpecRetentionWeekly struct {
	Count    int      `json:"count"`
	Weekdays []string `json:"weekdays"`
}

type AzurermRecoveryServicesProtectionPolicyVmSpecRetentionMonthly struct {
	Weeks    []string `json:"weeks"`
	Weekdays []string `json:"weekdays"`
	Count    int      `json:"count"`
}

type AzurermRecoveryServicesProtectionPolicyVmSpecRetentionYearly struct {
	Weeks    []string `json:"weeks"`
	Weekdays []string `json:"weekdays"`
	Count    int      `json:"count"`
	Months   []string `json:"months"`
}

type AzurermRecoveryServicesProtectionPolicyVmSpecBackup struct {
	Frequency string   `json:"frequency"`
	Time      string   `json:"time"`
	Weekdays  []string `json:"weekdays"`
}

type AzurermRecoveryServicesProtectionPolicyVmSpec struct {
	ResourceGroupName string                                          `json:"resource_group_name"`
	RetentionDaily    []AzurermRecoveryServicesProtectionPolicyVmSpec `json:"retention_daily"`
	RetentionWeekly   []AzurermRecoveryServicesProtectionPolicyVmSpec `json:"retention_weekly"`
	RetentionMonthly  []AzurermRecoveryServicesProtectionPolicyVmSpec `json:"retention_monthly"`
	RetentionYearly   []AzurermRecoveryServicesProtectionPolicyVmSpec `json:"retention_yearly"`
	Name              string                                          `json:"name"`
	RecoveryVaultName string                                          `json:"recovery_vault_name"`
	Timezone          string                                          `json:"timezone"`
	Backup            []AzurermRecoveryServicesProtectionPolicyVmSpec `json:"backup"`
	Tags              map[string]string                               `json:"tags"`
}



type AzurermRecoveryServicesProtectionPolicyVmStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRecoveryServicesProtectionPolicyVmList is a list of AzurermRecoveryServicesProtectionPolicyVms
type AzurermRecoveryServicesProtectionPolicyVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRecoveryServicesProtectionPolicyVm CRD objects
	Items []AzurermRecoveryServicesProtectionPolicyVm `json:"items,omitempty"`
}