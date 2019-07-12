package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermManagementLock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermManagementLockSpec   `json:"spec,omitempty"`
	Status            AzurermManagementLockStatus `json:"status,omitempty"`
}

type AzurermManagementLockSpec struct {
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	LockLevel string `json:"lock_level"`
	Notes     string `json:"notes"`
}

type AzurermManagementLockStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermManagementLockList is a list of AzurermManagementLocks
type AzurermManagementLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermManagementLock CRD objects
	Items []AzurermManagementLock `json:"items,omitempty"`
}
