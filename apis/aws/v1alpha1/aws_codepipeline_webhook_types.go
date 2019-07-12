package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodepipelineWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodepipelineWebhookSpec   `json:"spec,omitempty"`
	Status            AwsCodepipelineWebhookStatus `json:"status,omitempty"`
}

type AwsCodepipelineWebhookSpecFilter struct {
	JsonPath    string `json:"json_path"`
	MatchEquals string `json:"match_equals"`
}

type AwsCodepipelineWebhookSpecAuthenticationConfiguration struct {
	SecretToken    string `json:"secret_token"`
	AllowedIpRange string `json:"allowed_ip_range"`
}

type AwsCodepipelineWebhookSpec struct {
	Filter                      []AwsCodepipelineWebhookSpec `json:"filter"`
	Name                        string                       `json:"name"`
	Url                         string                       `json:"url"`
	TargetAction                string                       `json:"target_action"`
	TargetPipeline              string                       `json:"target_pipeline"`
	Authentication              string                       `json:"authentication"`
	AuthenticationConfiguration []AwsCodepipelineWebhookSpec `json:"authentication_configuration"`
}



type AwsCodepipelineWebhookStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipelineWebhookList is a list of AwsCodepipelineWebhooks
type AwsCodepipelineWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodepipelineWebhook CRD objects
	Items []AwsCodepipelineWebhook `json:"items,omitempty"`
}