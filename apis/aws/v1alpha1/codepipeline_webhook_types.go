package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type CodepipelineWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodepipelineWebhookSpec   `json:"spec,omitempty"`
	Status            CodepipelineWebhookStatus `json:"status,omitempty"`
}

type CodepipelineWebhookSpecAuthenticationConfiguration struct {
	// +optional
	AllowedIPRange string `json:"allowedIPRange,omitempty" tf:"allowed_ip_range,omitempty"`
	// +optional
	SecretToken string `json:"-" sensitive:"true" tf:"secret_token,omitempty"`
}

type CodepipelineWebhookSpecFilter struct {
	JsonPath    string `json:"jsonPath" tf:"json_path"`
	MatchEquals string `json:"matchEquals" tf:"match_equals"`
}

type CodepipelineWebhookSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Authentication string `json:"authentication" tf:"authentication"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	AuthenticationConfiguration []CodepipelineWebhookSpecAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`
	// +kubebuilder:validation:MinItems=1
	Filter         []CodepipelineWebhookSpecFilter `json:"filter" tf:"filter"`
	Name           string                          `json:"name" tf:"name"`
	TargetAction   string                          `json:"targetAction" tf:"target_action"`
	TargetPipeline string                          `json:"targetPipeline" tf:"target_pipeline"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
}

type CodepipelineWebhookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodepipelineWebhookSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
