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

type ConfigDeliveryChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigDeliveryChannelSpec   `json:"spec,omitempty"`
	Status            ConfigDeliveryChannelStatus `json:"status,omitempty"`
}

type ConfigDeliveryChannelSpecSnapshotDeliveryProperties struct {
	// +optional
	DeliveryFrequency string `json:"delivery_frequency,omitempty"`
}

type ConfigDeliveryChannelSpec struct {
	// +optional
	Name         string `json:"name,omitempty"`
	S3BucketName string `json:"s3_bucket_name"`
	// +optional
	S3KeyPrefix string `json:"s3_key_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotDeliveryProperties *[]ConfigDeliveryChannelSpec `json:"snapshot_delivery_properties,omitempty"`
	// +optional
	SnsTopicArn string `json:"sns_topic_arn,omitempty"`
}

type ConfigDeliveryChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigDeliveryChannelList is a list of ConfigDeliveryChannels
type ConfigDeliveryChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigDeliveryChannel CRD objects
	Items []ConfigDeliveryChannel `json:"items,omitempty"`
}
