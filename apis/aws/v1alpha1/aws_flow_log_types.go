package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsFlowLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsFlowLogSpec   `json:"spec,omitempty"`
	Status            AwsFlowLogStatus `json:"status,omitempty"`
}

type AwsFlowLogSpec struct {
	EniId              string `json:"eni_id"`
	TrafficType        string `json:"traffic_type"`
	IamRoleArn         string `json:"iam_role_arn"`
	LogDestination     string `json:"log_destination"`
	LogDestinationType string `json:"log_destination_type"`
	LogGroupName       string `json:"log_group_name"`
	VpcId              string `json:"vpc_id"`
	SubnetId           string `json:"subnet_id"`
}



type AwsFlowLogStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsFlowLogList is a list of AwsFlowLogs
type AwsFlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsFlowLog CRD objects
	Items []AwsFlowLog `json:"items,omitempty"`
}