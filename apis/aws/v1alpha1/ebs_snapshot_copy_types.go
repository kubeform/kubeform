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

type EbsSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsSnapshotCopySpec   `json:"spec,omitempty"`
	Status            EbsSnapshotCopyStatus `json:"status,omitempty"`
}

type EbsSnapshotCopySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	KmsKeyID         string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	SourceRegion     string `json:"sourceRegion" tf:"source_region"`
	SourceSnapshotID string `json:"sourceSnapshotID" tf:"source_snapshot_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EbsSnapshotCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EbsSnapshotCopySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EbsSnapshotCopyList is a list of EbsSnapshotCopys
type EbsSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EbsSnapshotCopy CRD objects
	Items []EbsSnapshotCopy `json:"items,omitempty"`
}
