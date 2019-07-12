package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticacheSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsElasticacheSecurityGroupStatus `json:"status,omitempty"`
}

type AwsElasticacheSecurityGroupSpec struct {
	Description        string   `json:"description"`
	Name               string   `json:"name"`
	SecurityGroupNames []string `json:"security_group_names"`
}



type AwsElasticacheSecurityGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheSecurityGroupList is a list of AwsElasticacheSecurityGroups
type AwsElasticacheSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheSecurityGroup CRD objects
	Items []AwsElasticacheSecurityGroup `json:"items,omitempty"`
}