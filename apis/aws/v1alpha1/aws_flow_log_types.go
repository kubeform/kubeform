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

type AwsFlowLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsFlowLogSpec   `json:"spec,omitempty"`
	Status            AwsFlowLogStatus `json:"status,omitempty"`
}

type AwsFlowLogSpec struct {
	VpcId              string `json:"vpc_id"`
	SubnetId           string `json:"subnet_id"`
	EniId              string `json:"eni_id"`
	TrafficType        string `json:"traffic_type"`
	IamRoleArn         string `json:"iam_role_arn"`
	LogDestination     string `json:"log_destination"`
	LogDestinationType string `json:"log_destination_type"`
	LogGroupName       string `json:"log_group_name"`
}



type AwsFlowLogStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsFlowLogList is a list of AwsFlowLogs
type AwsFlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsFlowLog CRD objects
	Items []AwsFlowLog `json:"items,omitempty"`
}