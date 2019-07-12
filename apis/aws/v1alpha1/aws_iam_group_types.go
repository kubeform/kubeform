package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamGroupSpec   `json:"spec,omitempty"`
	Status            AwsIamGroupStatus `json:"status,omitempty"`
}

type AwsIamGroupSpec struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Arn      string `json:"arn"`
	UniqueId string `json:"unique_id"`
}



type AwsIamGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupList is a list of AwsIamGroups
type AwsIamGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamGroup CRD objects
	Items []AwsIamGroup `json:"items,omitempty"`
}