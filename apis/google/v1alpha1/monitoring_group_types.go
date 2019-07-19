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

type MonitoringGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringGroupSpec   `json:"spec,omitempty"`
	Status            MonitoringGroupStatus `json:"status,omitempty"`
}

type MonitoringGroupSpec struct {
	DisplayName string `json:"displayName" tf:"display_name"`
	Filter      string `json:"filter" tf:"filter"`
	// +optional
	IsCluster bool `json:"isCluster,omitempty" tf:"is_cluster,omitempty"`
	// +optional
	ParentName string `json:"parentName,omitempty" tf:"parent_name,omitempty"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MonitoringGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
