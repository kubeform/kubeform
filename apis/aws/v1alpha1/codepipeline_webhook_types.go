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

type CodepipelineWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodepipelineWebhookSpec   `json:"spec,omitempty"`
	Status            CodepipelineWebhookStatus `json:"status,omitempty"`
}

type CodepipelineWebhookSpecAuthenticationConfiguration struct {
	// +optional
	AllowedIpRange string `json:"allowed_ip_range,omitempty"`
	// +optional
	SecretToken string `json:"secret_token,omitempty"`
}

type CodepipelineWebhookSpecFilter struct {
	JsonPath    string `json:"json_path"`
	MatchEquals string `json:"match_equals"`
}

type CodepipelineWebhookSpec struct {
	Authentication string `json:"authentication"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	AuthenticationConfiguration *[]CodepipelineWebhookSpec `json:"authentication_configuration,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Filter []CodepipelineWebhookSpec `json:"filter"`
	Name   string                    `json:"name"`
	// +optional
	Tags           map[string]string `json:"tags,omitempty"`
	TargetAction   string            `json:"target_action"`
	TargetPipeline string            `json:"target_pipeline"`
}

type CodepipelineWebhookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodepipelineWebhookList is a list of CodepipelineWebhooks
type CodepipelineWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodepipelineWebhook CRD objects
	Items []CodepipelineWebhook `json:"items,omitempty"`
}
