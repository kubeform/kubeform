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

type EbsSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsSnapshotSpec   `json:"spec,omitempty"`
	Status            EbsSnapshotStatus `json:"status,omitempty"`
}

type EbsSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DataEncryptionKeyID string `json:"dataEncryptionKeyID" tf:"data_encryption_key_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Encrypted   bool   `json:"encrypted" tf:"encrypted"`
	KmsKeyID    string `json:"kmsKeyID" tf:"kms_key_id"`
	OwnerAlias  string `json:"ownerAlias" tf:"owner_alias"`
	OwnerID     string `json:"ownerID" tf:"owner_id"`
	// +optional
	Tags       map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VolumeID   string            `json:"volumeID" tf:"volume_id"`
	VolumeSize int               `json:"volumeSize" tf:"volume_size"`
}

type EbsSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EbsSnapshotSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EbsSnapshotList is a list of EbsSnapshots
type EbsSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EbsSnapshot CRD objects
	Items []EbsSnapshot `json:"items,omitempty"`
}
