package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RdsGlobalCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsGlobalClusterSpec   `json:"spec,omitempty"`
	Status            RdsGlobalClusterStatus `json:"status,omitempty"`
}

type RdsGlobalClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	DatabaseName string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion           string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	GlobalClusterIdentifier string `json:"globalClusterIdentifier" tf:"global_cluster_identifier"`
	// +optional
	GlobalClusterResourceID string `json:"globalClusterResourceID,omitempty" tf:"global_cluster_resource_id,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
}

type RdsGlobalClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RdsGlobalClusterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsGlobalClusterList is a list of RdsGlobalClusters
type RdsGlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsGlobalCluster CRD objects
	Items []RdsGlobalCluster `json:"items,omitempty"`
}
