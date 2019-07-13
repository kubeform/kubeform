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

type AwsRedshiftSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftSecurityGroupStatus `json:"status,omitempty"`
}

type AwsRedshiftSecurityGroupSpecIngress struct {
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
	Cidr                 string `json:"cidr"`
}

type AwsRedshiftSecurityGroupSpec struct {
	Ingress     []AwsRedshiftSecurityGroupSpec `json:"ingress"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
}



type AwsRedshiftSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRedshiftSecurityGroupList is a list of AwsRedshiftSecurityGroups
type AwsRedshiftSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftSecurityGroup CRD objects
	Items []AwsRedshiftSecurityGroup `json:"items,omitempty"`
}