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

type AwsEbsSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEbsSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsEbsSnapshotStatus `json:"status,omitempty"`
}

type AwsEbsSnapshotSpec struct {
	VolumeId            string            `json:"volume_id"`
	Description         string            `json:"description"`
	KmsKeyId            string            `json:"kms_key_id"`
	DataEncryptionKeyId string            `json:"data_encryption_key_id"`
	Tags                map[string]string `json:"tags"`
	OwnerId             string            `json:"owner_id"`
	OwnerAlias          string            `json:"owner_alias"`
	Encrypted           bool              `json:"encrypted"`
	VolumeSize          int               `json:"volume_size"`
}

type AwsEbsSnapshotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEbsSnapshotList is a list of AwsEbsSnapshots
type AwsEbsSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEbsSnapshot CRD objects
	Items []AwsEbsSnapshot `json:"items,omitempty"`
}
