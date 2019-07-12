package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEmrInstanceGroupSpec   `json:"spec,omitempty"`
	Status            AwsEmrInstanceGroupStatus `json:"status,omitempty"`
}

type AwsEmrInstanceGroupSpecEbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

type AwsEmrInstanceGroupSpec struct {
	RunningInstanceCount int                       `json:"running_instance_count"`
	Status               string                    `json:"status"`
	Name                 string                    `json:"name"`
	EbsOptimized         bool                      `json:"ebs_optimized"`
	EbsConfig            []AwsEmrInstanceGroupSpec `json:"ebs_config"`
	ClusterId            string                    `json:"cluster_id"`
	InstanceType         string                    `json:"instance_type"`
	InstanceCount        int                       `json:"instance_count"`
}



type AwsEmrInstanceGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrInstanceGroupList is a list of AwsEmrInstanceGroups
type AwsEmrInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEmrInstanceGroup CRD objects
	Items []AwsEmrInstanceGroup `json:"items,omitempty"`
}