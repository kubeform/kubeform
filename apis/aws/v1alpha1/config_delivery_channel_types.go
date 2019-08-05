package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	DeliveryFrequency string `json:"deliveryFrequency,omitempty" tf:"delivery_frequency,omitempty"`
}

type ConfigDeliveryChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Name         string `json:"name,omitempty" tf:"name,omitempty"`
	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`
	// +optional
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotDeliveryProperties []ConfigDeliveryChannelSpecSnapshotDeliveryProperties `json:"snapshotDeliveryProperties,omitempty" tf:"snapshot_delivery_properties,omitempty"`
	// +optional
	SnsTopicArn string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`
}

type ConfigDeliveryChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigDeliveryChannelSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
