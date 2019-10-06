package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RedshiftSnapshotCopyGrant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftSnapshotCopyGrantSpec   `json:"spec,omitempty"`
	Status            RedshiftSnapshotCopyGrantStatus `json:"status,omitempty"`
}

type RedshiftSnapshotCopyGrantSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	KmsKeyID              string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	SnapshotCopyGrantName string `json:"snapshotCopyGrantName" tf:"snapshot_copy_grant_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RedshiftSnapshotCopyGrantStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RedshiftSnapshotCopyGrantSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftSnapshotCopyGrantList is a list of RedshiftSnapshotCopyGrants
type RedshiftSnapshotCopyGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftSnapshotCopyGrant CRD objects
	Items []RedshiftSnapshotCopyGrant `json:"items,omitempty"`
}
