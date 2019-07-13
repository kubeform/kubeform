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

type AwsRedshiftSnapshotCopyGrant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftSnapshotCopyGrantSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftSnapshotCopyGrantStatus `json:"status,omitempty"`
}

type AwsRedshiftSnapshotCopyGrantSpec struct {
	Arn                   string            `json:"arn"`
	SnapshotCopyGrantName string            `json:"snapshot_copy_grant_name"`
	KmsKeyId              string            `json:"kms_key_id"`
	Tags                  map[string]string `json:"tags"`
}



type AwsRedshiftSnapshotCopyGrantStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRedshiftSnapshotCopyGrantList is a list of AwsRedshiftSnapshotCopyGrants
type AwsRedshiftSnapshotCopyGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftSnapshotCopyGrant CRD objects
	Items []AwsRedshiftSnapshotCopyGrant `json:"items,omitempty"`
}