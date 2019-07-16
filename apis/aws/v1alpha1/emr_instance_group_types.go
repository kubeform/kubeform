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

type EmrInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrInstanceGroupSpec   `json:"spec,omitempty"`
	Status            EmrInstanceGroupStatus `json:"status,omitempty"`
}

type EmrInstanceGroupSpec struct {
	// +optional
	AutoscalingPolicy string `json:"autoscaling_policy,omitempty"`
	// +optional
	BidPrice  string `json:"bid_price,omitempty"`
	ClusterId string `json:"cluster_id"`
	// +optional
	EbsOptimized bool `json:"ebs_optimized,omitempty"`
	// +optional
	InstanceCount int    `json:"instance_count,omitempty"`
	InstanceType  string `json:"instance_type"`
	// +optional
	Name string `json:"name,omitempty"`
}

type EmrInstanceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
