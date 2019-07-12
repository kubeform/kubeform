package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53QueryLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53QueryLogSpec   `json:"spec,omitempty"`
	Status            AwsRoute53QueryLogStatus `json:"status,omitempty"`
}

type AwsRoute53QueryLogSpec struct {
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	ZoneId                string `json:"zone_id"`
}

type AwsRoute53QueryLogStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53QueryLogList is a list of AwsRoute53QueryLogs
type AwsRoute53QueryLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53QueryLog CRD objects
	Items []AwsRoute53QueryLog `json:"items,omitempty"`
}
