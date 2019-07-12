package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsResourcegroupsGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsResourcegroupsGroupSpec   `json:"spec,omitempty"`
	Status            AwsResourcegroupsGroupStatus `json:"status,omitempty"`
}

type AwsResourcegroupsGroupSpecResourceQuery struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

type AwsResourcegroupsGroupSpec struct {
	ResourceQuery []AwsResourcegroupsGroupSpec `json:"resource_query"`
	Arn           string                       `json:"arn"`
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
}



type AwsResourcegroupsGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsResourcegroupsGroupList is a list of AwsResourcegroupsGroups
type AwsResourcegroupsGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsResourcegroupsGroup CRD objects
	Items []AwsResourcegroupsGroup `json:"items,omitempty"`
}