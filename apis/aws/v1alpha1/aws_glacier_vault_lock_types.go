package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlacierVaultLock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlacierVaultLockSpec   `json:"spec,omitempty"`
	Status            AwsGlacierVaultLockStatus `json:"status,omitempty"`
}

type AwsGlacierVaultLockSpec struct {
	CompleteLock        bool   `json:"complete_lock"`
	IgnoreDeletionError bool   `json:"ignore_deletion_error"`
	Policy              string `json:"policy"`
	VaultName           string `json:"vault_name"`
}



type AwsGlacierVaultLockStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlacierVaultLockList is a list of AwsGlacierVaultLocks
type AwsGlacierVaultLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlacierVaultLock CRD objects
	Items []AwsGlacierVaultLock `json:"items,omitempty"`
}