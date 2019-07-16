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

type SnsPlatformApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsPlatformApplicationSpec   `json:"spec,omitempty"`
	Status            SnsPlatformApplicationStatus `json:"status,omitempty"`
}

type SnsPlatformApplicationSpec struct {
	// +optional
	EventDeliveryFailureTopicArn string `json:"event_delivery_failure_topic_arn,omitempty"`
	// +optional
	EventEndpointCreatedTopicArn string `json:"event_endpoint_created_topic_arn,omitempty"`
	// +optional
	EventEndpointDeletedTopicArn string `json:"event_endpoint_deleted_topic_arn,omitempty"`
	// +optional
	EventEndpointUpdatedTopicArn string `json:"event_endpoint_updated_topic_arn,omitempty"`
	// +optional
	FailureFeedbackRoleArn string `json:"failure_feedback_role_arn,omitempty"`
	Name                   string `json:"name"`
	Platform               string `json:"platform"`
	PlatformCredential     string `json:"platform_credential"`
	// +optional
	PlatformPrincipal string `json:"platform_principal,omitempty"`
	// +optional
	SuccessFeedbackRoleArn string `json:"success_feedback_role_arn,omitempty"`
	// +optional
	SuccessFeedbackSampleRate string `json:"success_feedback_sample_rate,omitempty"`
}

type SnsPlatformApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsPlatformApplicationList is a list of SnsPlatformApplications
type SnsPlatformApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsPlatformApplication CRD objects
	Items []SnsPlatformApplication `json:"items,omitempty"`
}
