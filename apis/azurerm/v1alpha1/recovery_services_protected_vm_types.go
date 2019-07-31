package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RecoveryServicesProtectedVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryServicesProtectedVmSpec   `json:"spec,omitempty"`
	Status            RecoveryServicesProtectedVmStatus `json:"status,omitempty"`
}

type RecoveryServicesProtectedVmSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackupPolicyID    string `json:"backupPolicyID" tf:"backup_policy_id"`
	RecoveryVaultName string `json:"recoveryVaultName" tf:"recovery_vault_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	SourceVmID        string `json:"sourceVmID" tf:"source_vm_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RecoveryServicesProtectedVmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RecoveryServicesProtectedVmList is a list of RecoveryServicesProtectedVms
type RecoveryServicesProtectedVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RecoveryServicesProtectedVm CRD objects
	Items []RecoveryServicesProtectedVm `json:"items,omitempty"`
}
