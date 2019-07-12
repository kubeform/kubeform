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

type AwsPlacementGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPlacementGroupSpec   `json:"spec,omitempty"`
	Status            AwsPlacementGroupStatus `json:"status,omitempty"`
}

type AwsPlacementGroupSpec struct {
	Name     string `json:"name"`
	Strategy string `json:"strategy"`
}

type AwsPlacementGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsPlacementGroupList is a list of AwsPlacementGroups
type AwsPlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPlacementGroup CRD objects
	Items []AwsPlacementGroup `json:"items,omitempty"`
}
