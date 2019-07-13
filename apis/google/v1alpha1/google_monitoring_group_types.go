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

type GoogleMonitoringGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleMonitoringGroupSpec   `json:"spec,omitempty"`
	Status            GoogleMonitoringGroupStatus `json:"status,omitempty"`
}

type GoogleMonitoringGroupSpec struct {
	Name        string `json:"name"`
	Project     string `json:"project"`
	DisplayName string `json:"display_name"`
	Filter      string `json:"filter"`
	IsCluster   bool   `json:"is_cluster"`
	ParentName  string `json:"parent_name"`
}



type GoogleMonitoringGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleMonitoringGroupList is a list of GoogleMonitoringGroups
type GoogleMonitoringGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleMonitoringGroup CRD objects
	Items []GoogleMonitoringGroup `json:"items,omitempty"`
}