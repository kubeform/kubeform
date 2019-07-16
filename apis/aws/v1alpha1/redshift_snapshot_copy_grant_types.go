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

type RedshiftSnapshotCopyGrant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftSnapshotCopyGrantSpec   `json:"spec,omitempty"`
	Status            RedshiftSnapshotCopyGrantStatus `json:"status,omitempty"`
}

type RedshiftSnapshotCopyGrantSpec struct {
	SnapshotCopyGrantName string `json:"snapshot_copy_grant_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RedshiftSnapshotCopyGrantStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
