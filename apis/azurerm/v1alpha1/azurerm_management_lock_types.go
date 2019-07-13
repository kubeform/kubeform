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

type AzurermManagementLock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermManagementLockSpec   `json:"spec,omitempty"`
	Status            AzurermManagementLockStatus `json:"status,omitempty"`
}

type AzurermManagementLockSpec struct {
	LockLevel string `json:"lock_level"`
	Notes     string `json:"notes"`
	Name      string `json:"name"`
	Scope     string `json:"scope"`
}



type AzurermManagementLockStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermManagementLockList is a list of AzurermManagementLocks
type AzurermManagementLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermManagementLock CRD objects
	Items []AzurermManagementLock `json:"items,omitempty"`
}