package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftSubnetGroupStatus `json:"status,omitempty"`
}

type AwsRedshiftSubnetGroupSpec struct {
	Description string            `json:"description"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
	Name        string            `json:"name"`
}

type AwsRedshiftSubnetGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSubnetGroupList is a list of AwsRedshiftSubnetGroups
type AwsRedshiftSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftSubnetGroup CRD objects
	Items []AwsRedshiftSubnetGroup `json:"items,omitempty"`
}
