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

type AwsSnapshotCreateVolumePermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnapshotCreateVolumePermissionSpec   `json:"spec,omitempty"`
	Status            AwsSnapshotCreateVolumePermissionStatus `json:"status,omitempty"`
}

type AwsSnapshotCreateVolumePermissionSpec struct {
	SnapshotId string `json:"snapshot_id"`
	AccountId  string `json:"account_id"`
}



type AwsSnapshotCreateVolumePermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSnapshotCreateVolumePermissionList is a list of AwsSnapshotCreateVolumePermissions
type AwsSnapshotCreateVolumePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnapshotCreateVolumePermission CRD objects
	Items []AwsSnapshotCreateVolumePermission `json:"items,omitempty"`
}