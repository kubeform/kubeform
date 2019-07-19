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

type ResourcegroupsGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcegroupsGroupSpec   `json:"spec,omitempty"`
	Status            ResourcegroupsGroupStatus `json:"status,omitempty"`
}

type ResourcegroupsGroupSpecResourceQuery struct {
	Query string `json:"query" tf:"query"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourcegroupsGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	ResourceQuery []ResourcegroupsGroupSpecResourceQuery `json:"resourceQuery" tf:"resource_query"`
	ProviderRef   core.LocalObjectReference              `json:"providerRef" tf:"-"`
}

type ResourcegroupsGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ResourcegroupsGroupList is a list of ResourcegroupsGroups
type ResourcegroupsGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ResourcegroupsGroup CRD objects
	Items []ResourcegroupsGroup `json:"items,omitempty"`
}
