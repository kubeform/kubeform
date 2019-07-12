package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogGroupSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogGroupStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogGroupSpec struct {
	NamePrefix      string            `json:"name_prefix"`
	RetentionInDays int               `json:"retention_in_days"`
	KmsKeyId        string            `json:"kms_key_id"`
	Arn             string            `json:"arn"`
	Tags            map[string]string `json:"tags"`
	Name            string            `json:"name"`
}



type AwsCloudwatchLogGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogGroupList is a list of AwsCloudwatchLogGroups
type AwsCloudwatchLogGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogGroup CRD objects
	Items []AwsCloudwatchLogGroup `json:"items,omitempty"`
}