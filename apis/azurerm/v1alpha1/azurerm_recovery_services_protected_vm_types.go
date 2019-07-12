package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRecoveryServicesProtectedVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRecoveryServicesProtectedVmSpec   `json:"spec,omitempty"`
	Status            AzurermRecoveryServicesProtectedVmStatus `json:"status,omitempty"`
}

type AzurermRecoveryServicesProtectedVmSpec struct {
	Tags              map[string]string `json:"tags"`
	ResourceGroupName string            `json:"resource_group_name"`
	RecoveryVaultName string            `json:"recovery_vault_name"`
	SourceVmId        string            `json:"source_vm_id"`
	BackupPolicyId    string            `json:"backup_policy_id"`
}



type AzurermRecoveryServicesProtectedVmStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRecoveryServicesProtectedVmList is a list of AzurermRecoveryServicesProtectedVms
type AzurermRecoveryServicesProtectedVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRecoveryServicesProtectedVm CRD objects
	Items []AzurermRecoveryServicesProtectedVm `json:"items,omitempty"`
}