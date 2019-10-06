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

type EmrInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrInstanceGroupSpec   `json:"spec,omitempty"`
	Status            EmrInstanceGroupStatus `json:"status,omitempty"`
}

type EmrInstanceGroupSpecEbsConfig struct {
	// +optional
	Iops int    `json:"iops,omitempty" tf:"iops,omitempty"`
	Size int    `json:"size" tf:"size"`
	Type string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance int `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type EmrInstanceGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterID string `json:"clusterID" tf:"cluster_id"`
	// +optional
	EbsConfig []EmrInstanceGroupSpecEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	InstanceCount int    `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`
	InstanceType  string `json:"instanceType" tf:"instance_type"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	RunningInstanceCount int `json:"runningInstanceCount,omitempty" tf:"running_instance_count,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
}

type EmrInstanceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EmrInstanceGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EmrInstanceGroupList is a list of EmrInstanceGroups
type EmrInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EmrInstanceGroup CRD objects
	Items []EmrInstanceGroup `json:"items,omitempty"`
}
