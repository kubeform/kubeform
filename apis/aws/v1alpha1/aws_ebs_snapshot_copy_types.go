package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEbsSnapshotCopySpec   `json:"spec,omitempty"`
	Status            AwsEbsSnapshotCopyStatus `json:"status,omitempty"`
}

type AwsEbsSnapshotCopySpec struct {
	VolumeId            string            `json:"volume_id"`
	OwnerAlias          string            `json:"owner_alias"`
	VolumeSize          int               `json:"volume_size"`
	KmsKeyId            string            `json:"kms_key_id"`
	DataEncryptionKeyId string            `json:"data_encryption_key_id"`
	SourceRegion        string            `json:"source_region"`
	Description         string            `json:"description"`
	OwnerId             string            `json:"owner_id"`
	Encrypted           bool              `json:"encrypted"`
	SourceSnapshotId    string            `json:"source_snapshot_id"`
	Tags                map[string]string `json:"tags"`
}



type AwsEbsSnapshotCopyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEbsSnapshotCopyList is a list of AwsEbsSnapshotCopys
type AwsEbsSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEbsSnapshotCopy CRD objects
	Items []AwsEbsSnapshotCopy `json:"items,omitempty"`
}