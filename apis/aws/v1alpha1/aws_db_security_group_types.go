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

type AwsDbSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsDbSecurityGroupStatus `json:"status,omitempty"`
}

type AwsDbSecurityGroupSpecIngress struct {
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
	Cidr                 string `json:"cidr"`
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupId      string `json:"security_group_id"`
}

type AwsDbSecurityGroupSpec struct {
	Ingress     []AwsDbSecurityGroupSpec `json:"ingress"`
	Tags        map[string]string        `json:"tags"`
	Arn         string                   `json:"arn"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
}

type AwsDbSecurityGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbSecurityGroupList is a list of AwsDbSecurityGroups
type AwsDbSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbSecurityGroup CRD objects
	Items []AwsDbSecurityGroup `json:"items,omitempty"`
}
