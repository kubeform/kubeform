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

type MonitoringGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringGroupSpec   `json:"spec,omitempty"`
	Status            MonitoringGroupStatus `json:"status,omitempty"`
}

type MonitoringGroupSpec struct {
	DisplayName string `json:"display_name"`
	Filter      string `json:"filter"`
	// +optional
	IsCluster bool `json:"is_cluster,omitempty"`
	// +optional
	ParentName string `json:"parent_name,omitempty"`
}

type MonitoringGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitoringGroupList is a list of MonitoringGroups
type MonitoringGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitoringGroup CRD objects
	Items []MonitoringGroup `json:"items,omitempty"`
}
