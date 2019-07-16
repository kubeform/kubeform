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

type ResourcegroupsGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcegroupsGroupSpec   `json:"spec,omitempty"`
	Status            ResourcegroupsGroupStatus `json:"status,omitempty"`
}

type ResourcegroupsGroupSpecResourceQuery struct {
	Query string `json:"query"`
	// +optional
	Type string `json:"type,omitempty"`
}

type ResourcegroupsGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	ResourceQuery []ResourcegroupsGroupSpec `json:"resource_query"`
}

type ResourcegroupsGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
