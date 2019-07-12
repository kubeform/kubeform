package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNeptuneSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneSubnetGroupStatus `json:"status,omitempty"`
}

type AwsNeptuneSubnetGroupSpec struct {
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Description string            `json:"description"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
	Arn         string            `json:"arn"`
}



type AwsNeptuneSubnetGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNeptuneSubnetGroupList is a list of AwsNeptuneSubnetGroups
type AwsNeptuneSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneSubnetGroup CRD objects
	Items []AwsNeptuneSubnetGroup `json:"items,omitempty"`
}