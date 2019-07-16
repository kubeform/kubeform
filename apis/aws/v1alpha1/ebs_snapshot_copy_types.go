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

type EbsSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsSnapshotCopySpec   `json:"spec,omitempty"`
	Status            EbsSnapshotCopyStatus `json:"status,omitempty"`
}

type EbsSnapshotCopySpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty"`
	// +optional
	KmsKeyId         string `json:"kms_key_id,omitempty"`
	SourceRegion     string `json:"source_region"`
	SourceSnapshotId string `json:"source_snapshot_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type EbsSnapshotCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
