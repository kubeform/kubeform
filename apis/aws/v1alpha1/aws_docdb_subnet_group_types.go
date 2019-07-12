package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDocdbSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsDocdbSubnetGroupStatus `json:"status,omitempty"`
}

type AwsDocdbSubnetGroupSpec struct {
	Arn         string            `json:"arn"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Description string            `json:"description"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
}

type AwsDocdbSubnetGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDocdbSubnetGroupList is a list of AwsDocdbSubnetGroups
type AwsDocdbSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbSubnetGroup CRD objects
	Items []AwsDocdbSubnetGroup `json:"items,omitempty"`
}
