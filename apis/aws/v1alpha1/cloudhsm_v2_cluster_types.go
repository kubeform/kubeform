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

type CloudhsmV2Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudhsmV2ClusterSpec   `json:"spec,omitempty"`
	Status            CloudhsmV2ClusterStatus `json:"status,omitempty"`
}

type CloudhsmV2ClusterSpec struct {
	HsmType string `json:"hsm_type"`
	// +optional
	SourceBackupIdentifier string `json:"source_backup_identifier,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type CloudhsmV2ClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudhsmV2ClusterList is a list of CloudhsmV2Clusters
type CloudhsmV2ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudhsmV2Cluster CRD objects
	Items []CloudhsmV2Cluster `json:"items,omitempty"`
}
