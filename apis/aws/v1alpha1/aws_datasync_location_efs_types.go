package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDatasyncLocationEfs struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDatasyncLocationEfsSpec   `json:"spec,omitempty"`
	Status            AwsDatasyncLocationEfsStatus `json:"status,omitempty"`
}

type AwsDatasyncLocationEfsSpecEc2Config struct {
	SubnetArn         string   `json:"subnet_arn"`
	SecurityGroupArns []string `json:"security_group_arns"`
}

type AwsDatasyncLocationEfsSpec struct {
	Subdirectory     string                       `json:"subdirectory"`
	Tags             map[string]string            `json:"tags"`
	Uri              string                       `json:"uri"`
	Arn              string                       `json:"arn"`
	Ec2Config        []AwsDatasyncLocationEfsSpec `json:"ec2_config"`
	EfsFileSystemArn string                       `json:"efs_file_system_arn"`
}

type AwsDatasyncLocationEfsStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDatasyncLocationEfsList is a list of AwsDatasyncLocationEfss
type AwsDatasyncLocationEfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDatasyncLocationEfs CRD objects
	Items []AwsDatasyncLocationEfs `json:"items,omitempty"`
}
