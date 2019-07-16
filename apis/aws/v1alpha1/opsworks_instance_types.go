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

type OpsworksInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksInstanceSpec   `json:"spec,omitempty"`
	Status            OpsworksInstanceStatus `json:"status,omitempty"`
}

type OpsworksInstanceSpec struct {
	// +optional
	AgentVersion string `json:"agent_version,omitempty"`
	// +optional
	Architecture string `json:"architecture,omitempty"`
	// +optional
	AutoScalingType string `json:"auto_scaling_type,omitempty"`
	// +optional
	DeleteEbs bool `json:"delete_ebs,omitempty"`
	// +optional
	DeleteEip bool `json:"delete_eip,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebs_optimized,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceType string   `json:"instance_type,omitempty"`
	LayerIds     []string `json:"layer_ids"`
	StackId      string   `json:"stack_id"`
	// +optional
	State string `json:"state,omitempty"`
}

type OpsworksInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksInstanceList is a list of OpsworksInstances
type OpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksInstance CRD objects
	Items []OpsworksInstance `json:"items,omitempty"`
}
