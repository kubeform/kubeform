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

type S3BucketNotification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketNotificationSpec   `json:"spec,omitempty"`
	Status            S3BucketNotificationStatus `json:"status,omitempty"`
}

type S3BucketNotificationSpecLambdaFunction struct {
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events"`
	// +optional
	FilterPrefix string `json:"filter_prefix,omitempty"`
	// +optional
	FilterSuffix string `json:"filter_suffix,omitempty"`
	// +optional
	LambdaFunctionArn string `json:"lambda_function_arn,omitempty"`
}

type S3BucketNotificationSpecQueue struct {
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events"`
	// +optional
	FilterPrefix string `json:"filter_prefix,omitempty"`
	// +optional
	FilterSuffix string `json:"filter_suffix,omitempty"`
	QueueArn     string `json:"queue_arn"`
}

type S3BucketNotificationSpecTopic struct {
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events"`
	// +optional
	FilterPrefix string `json:"filter_prefix,omitempty"`
	// +optional
	FilterSuffix string `json:"filter_suffix,omitempty"`
	TopicArn     string `json:"topic_arn"`
}

type S3BucketNotificationSpec struct {
	Bucket string `json:"bucket"`
	// +optional
	LambdaFunction *[]S3BucketNotificationSpec `json:"lambda_function,omitempty"`
	// +optional
	Queue *[]S3BucketNotificationSpec `json:"queue,omitempty"`
	// +optional
	Topic *[]S3BucketNotificationSpec `json:"topic,omitempty"`
}

type S3BucketNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3BucketNotificationList is a list of S3BucketNotifications
type S3BucketNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3BucketNotification CRD objects
	Items []S3BucketNotification `json:"items,omitempty"`
}
