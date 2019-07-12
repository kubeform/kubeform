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

type AwsOpsworksRdsDbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksRdsDbInstanceSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksRdsDbInstanceStatus `json:"status,omitempty"`
}

type AwsOpsworksRdsDbInstanceSpec struct {
	RdsDbInstanceArn string `json:"rds_db_instance_arn"`
	DbPassword       string `json:"db_password"`
	DbUser           string `json:"db_user"`
	StackId          string `json:"stack_id"`
}

type AwsOpsworksRdsDbInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksRdsDbInstanceList is a list of AwsOpsworksRdsDbInstances
type AwsOpsworksRdsDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksRdsDbInstance CRD objects
	Items []AwsOpsworksRdsDbInstance `json:"items,omitempty"`
}
