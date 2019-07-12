package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorResourceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsInspectorResourceGroupSpec   `json:"spec,omitempty"`
	Status            AwsInspectorResourceGroupStatus `json:"status,omitempty"`
}

type AwsInspectorResourceGroupSpec struct {
	Tags map[string]string `json:"tags"`
	Arn  string            `json:"arn"`
}

type AwsInspectorResourceGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorResourceGroupList is a list of AwsInspectorResourceGroups
type AwsInspectorResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInspectorResourceGroup CRD objects
	Items []AwsInspectorResourceGroup `json:"items,omitempty"`
}
