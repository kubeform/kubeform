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

type AwsDbSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsDbSubnetGroupStatus `json:"status,omitempty"`
}

type AwsDbSubnetGroupSpec struct {
	Arn         string            `json:"arn"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Description string            `json:"description"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
}

type AwsDbSubnetGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbSubnetGroupList is a list of AwsDbSubnetGroups
type AwsDbSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbSubnetGroup CRD objects
	Items []AwsDbSubnetGroup `json:"items,omitempty"`
}
