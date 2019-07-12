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
	BidPrice             string                    `json:"bid_price"`
	EbsConfig            []AwsEmrInstanceGroupSpec `json:"ebs_config"`
	InstanceCount        int                       `json:"instance_count"`
	InstanceType         string                    `json:"instance_type"`
	RunningInstanceCount int                       `json:"running_instance_count"`
	Status               string                    `json:"status"`
	AutoscalingPolicy    string                    `json:"autoscaling_policy"`
	ClusterId            string                    `json:"cluster_id"`
	EbsOptimized         bool                      `json:"ebs_optimized"`
	Name                 string                    `json:"name"`
}

type AwsEmrInstanceGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEmrInstanceGroupList is a list of AwsEmrInstanceGroups
type AwsEmrInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEmrInstanceGroup CRD objects
	Items []AwsEmrInstanceGroup `json:"items,omitempty"`
}
