package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	Events []string `json:"events" tf:"events"`
	// +optional
	FilterPrefix string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`
	// +optional
	FilterSuffix string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	LambdaFunctionArn string `json:"lambdaFunctionArn,omitempty" tf:"lambda_function_arn,omitempty"`
}

type S3BucketNotificationSpecQueue struct {
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events" tf:"events"`
	// +optional
	FilterPrefix string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`
	// +optional
	FilterSuffix string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`
	// +optional
	ID       string `json:"ID,omitempty" tf:"id,omitempty"`
	QueueArn string `json:"queueArn" tf:"queue_arn"`
}

type S3BucketNotificationSpecTopic struct {
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events" tf:"events"`
	// +optional
	FilterPrefix string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`
	// +optional
	FilterSuffix string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`
	// +optional
	ID       string `json:"ID,omitempty" tf:"id,omitempty"`
	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

type S3BucketNotificationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	LambdaFunction []S3BucketNotificationSpecLambdaFunction `json:"lambdaFunction,omitempty" tf:"lambda_function,omitempty"`
	// +optional
	Queue []S3BucketNotificationSpecQueue `json:"queue,omitempty" tf:"queue,omitempty"`
	// +optional
	Topic []S3BucketNotificationSpecTopic `json:"topic,omitempty" tf:"topic,omitempty"`
}

type S3BucketNotificationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *S3BucketNotificationSpec `json:"state,omitempty"`
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
