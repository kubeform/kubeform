package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type BigtableInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigtableInstanceSpec   `json:"spec,omitempty"`
	Status            BigtableInstanceStatus `json:"status,omitempty"`
}

type BigtableInstanceSpecCluster struct {
	// +optional
	ClusterID string `json:"clusterID,omitempty" tf:"cluster_id,omitempty"`
	// +optional
	NumNodes int `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`
	// +optional
	StorageType string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type BigtableInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Cluster []BigtableInstanceSpecCluster `json:"cluster,omitempty" tf:"cluster,omitempty"`
	// +optional
	// Deprecated
	ClusterID string `json:"clusterID,omitempty" tf:"cluster_id,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	InstanceType string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	// Deprecated
	NumNodes int `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// Deprecated
	StorageType string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
	// +optional
	// Deprecated
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type BigtableInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigtableInstanceList is a list of BigtableInstances
type BigtableInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigtableInstance CRD objects
	Items []BigtableInstance `json:"items,omitempty"`
}
