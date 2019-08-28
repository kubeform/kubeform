package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SnapshotCreateVolumePermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotCreateVolumePermissionSpec   `json:"spec,omitempty"`
	Status            SnapshotCreateVolumePermissionStatus `json:"status,omitempty"`
}

type SnapshotCreateVolumePermissionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountID  string `json:"accountID" tf:"account_id"`
	SnapshotID string `json:"snapshotID" tf:"snapshot_id"`
}



type SnapshotCreateVolumePermissionStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *SnapshotCreateVolumePermissionSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnapshotCreateVolumePermissionList is a list of SnapshotCreateVolumePermissions
type SnapshotCreateVolumePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnapshotCreateVolumePermission CRD objects
	Items []SnapshotCreateVolumePermission `json:"items,omitempty"`
}