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

type AwsRdsGlobalCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsGlobalClusterSpec   `json:"spec,omitempty"`
	Status            AwsRdsGlobalClusterStatus `json:"status,omitempty"`
}

type AwsRdsGlobalClusterSpec struct {
	DeletionProtection      bool   `json:"deletion_protection"`
	Engine                  string `json:"engine"`
	EngineVersion           string `json:"engine_version"`
	GlobalClusterIdentifier string `json:"global_cluster_identifier"`
	GlobalClusterResourceId string `json:"global_cluster_resource_id"`
	StorageEncrypted        bool   `json:"storage_encrypted"`
	Arn                     string `json:"arn"`
	DatabaseName            string `json:"database_name"`
}

type AwsRdsGlobalClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRdsGlobalClusterList is a list of AwsRdsGlobalClusters
type AwsRdsGlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsGlobalCluster CRD objects
	Items []AwsRdsGlobalCluster `json:"items,omitempty"`
}
