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

type AwsSnsPlatformApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnsPlatformApplicationSpec   `json:"spec,omitempty"`
	Status            AwsSnsPlatformApplicationStatus `json:"status,omitempty"`
}

type AwsSnsPlatformApplicationSpec struct {
	FailureFeedbackRoleArn       string `json:"failure_feedback_role_arn"`
	PlatformCredential           string `json:"platform_credential"`
	EventDeliveryFailureTopicArn string `json:"event_delivery_failure_topic_arn"`
	EventEndpointCreatedTopicArn string `json:"event_endpoint_created_topic_arn"`
	EventEndpointDeletedTopicArn string `json:"event_endpoint_deleted_topic_arn"`
	EventEndpointUpdatedTopicArn string `json:"event_endpoint_updated_topic_arn"`
	SuccessFeedbackSampleRate    string `json:"success_feedback_sample_rate"`
	Name                         string `json:"name"`
	Platform                     string `json:"platform"`
	Arn                          string `json:"arn"`
	PlatformPrincipal            string `json:"platform_principal"`
	SuccessFeedbackRoleArn       string `json:"success_feedback_role_arn"`
}



type AwsSnsPlatformApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSnsPlatformApplicationList is a list of AwsSnsPlatformApplications
type AwsSnsPlatformApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnsPlatformApplication CRD objects
	Items []AwsSnsPlatformApplication `json:"items,omitempty"`
}