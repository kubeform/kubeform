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

type AwsEbsVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEbsVolumeSpec   `json:"spec,omitempty"`
	Status            AwsEbsVolumeStatus `json:"status,omitempty"`
}

type AwsEbsVolumeSpec struct {
	Iops             int               `json:"iops"`
	KmsKeyId         string            `json:"kms_key_id"`
	SnapshotId       string            `json:"snapshot_id"`
	Type             string            `json:"type"`
	Tags             map[string]string `json:"tags"`
	Arn              string            `json:"arn"`
	AvailabilityZone string            `json:"availability_zone"`
	Encrypted        bool              `json:"encrypted"`
	Size             int               `json:"size"`
}



type AwsEbsVolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEbsVolumeList is a list of AwsEbsVolumes
type AwsEbsVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEbsVolume CRD objects
	Items []AwsEbsVolume `json:"items,omitempty"`
}