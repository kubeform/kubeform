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

type RecoveryServicesProtectedVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryServicesProtectedVmSpec   `json:"spec,omitempty"`
	Status            RecoveryServicesProtectedVmStatus `json:"status,omitempty"`
}

type RecoveryServicesProtectedVmSpec struct {
	BackupPolicyId    string `json:"backup_policy_id"`
	RecoveryVaultName string `json:"recovery_vault_name"`
	ResourceGroupName string `json:"resource_group_name"`
	SourceVmId        string `json:"source_vm_id"`
}

type RecoveryServicesProtectedVmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
