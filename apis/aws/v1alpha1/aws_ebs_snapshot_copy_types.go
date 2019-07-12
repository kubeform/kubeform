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

type AwsEbsSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEbsSnapshotCopySpec   `json:"spec,omitempty"`
	Status            AwsEbsSnapshotCopyStatus `json:"status,omitempty"`
}

type AwsEbsSnapshotCopySpec struct {
	SourceRegion        string            `json:"source_region"`
	SourceSnapshotId    string            `json:"source_snapshot_id"`
	Tags                map[string]string `json:"tags"`
	Description         string            `json:"description"`
	Encrypted           bool              `json:"encrypted"`
	KmsKeyId            string            `json:"kms_key_id"`
	VolumeSize          int               `json:"volume_size"`
	DataEncryptionKeyId string            `json:"data_encryption_key_id"`
	VolumeId            string            `json:"volume_id"`
	OwnerId             string            `json:"owner_id"`
	OwnerAlias          string            `json:"owner_alias"`
}

type AwsEbsSnapshotCopyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEbsSnapshotCopyList is a list of AwsEbsSnapshotCopys
type AwsEbsSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEbsSnapshotCopy CRD objects
	Items []AwsEbsSnapshotCopy `json:"items,omitempty"`
}
