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

type AwsS3BucketNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketNotificationSpec   `json:"spec,omitempty"`
	Status            AwsS3BucketNotificationStatus `json:"status,omitempty"`
}

type AwsS3BucketNotificationSpecTopic struct {
	Events       []string `json:"events"`
	Id           string   `json:"id"`
	FilterPrefix string   `json:"filter_prefix"`
	FilterSuffix string   `json:"filter_suffix"`
	TopicArn     string   `json:"topic_arn"`
}

type AwsS3BucketNotificationSpecQueue struct {
	FilterPrefix string   `json:"filter_prefix"`
	FilterSuffix string   `json:"filter_suffix"`
	QueueArn     string   `json:"queue_arn"`
	Events       []string `json:"events"`
	Id           string   `json:"id"`
}

type AwsS3BucketNotificationSpecLambdaFunction struct {
	Id                string   `json:"id"`
	FilterPrefix      string   `json:"filter_prefix"`
	FilterSuffix      string   `json:"filter_suffix"`
	LambdaFunctionArn string   `json:"lambda_function_arn"`
	Events            []string `json:"events"`
}

type AwsS3BucketNotificationSpec struct {
	Bucket         string                        `json:"bucket"`
	Topic          []AwsS3BucketNotificationSpec `json:"topic"`
	Queue          []AwsS3BucketNotificationSpec `json:"queue"`
	LambdaFunction []AwsS3BucketNotificationSpec `json:"lambda_function"`
}



type AwsS3BucketNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsS3BucketNotificationList is a list of AwsS3BucketNotifications
type AwsS3BucketNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketNotification CRD objects
	Items []AwsS3BucketNotification `json:"items,omitempty"`
}