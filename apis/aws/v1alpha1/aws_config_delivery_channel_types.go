package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigDeliveryChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsConfigDeliveryChannelSpec   `json:"spec,omitempty"`
	Status            AwsConfigDeliveryChannelStatus `json:"status,omitempty"`
}

type AwsConfigDeliveryChannelSpecSnapshotDeliveryProperties struct {
	DeliveryFrequency string `json:"delivery_frequency"`
}

type AwsConfigDeliveryChannelSpec struct {
	Name                       string                         `json:"name"`
	S3BucketName               string                         `json:"s3_bucket_name"`
	S3KeyPrefix                string                         `json:"s3_key_prefix"`
	SnsTopicArn                string                         `json:"sns_topic_arn"`
	SnapshotDeliveryProperties []AwsConfigDeliveryChannelSpec `json:"snapshot_delivery_properties"`
}

type AwsConfigDeliveryChannelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigDeliveryChannelList is a list of AwsConfigDeliveryChannels
type AwsConfigDeliveryChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsConfigDeliveryChannel CRD objects
	Items []AwsConfigDeliveryChannel `json:"items,omitempty"`
}
