package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmPatchGroupSpec   `json:"spec,omitempty"`
	Status            AwsSsmPatchGroupStatus `json:"status,omitempty"`
}

type AwsSsmPatchGroupSpec struct {
	BaselineId string `json:"baseline_id"`
	PatchGroup string `json:"patch_group"`
}

type AwsSsmPatchGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchGroupList is a list of AwsSsmPatchGroups
type AwsSsmPatchGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmPatchGroup CRD objects
	Items []AwsSsmPatchGroup `json:"items,omitempty"`
}
