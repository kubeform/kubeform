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

type AwsCodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodebuildWebhookSpec   `json:"spec,omitempty"`
	Status            AwsCodebuildWebhookStatus `json:"status,omitempty"`
}

type AwsCodebuildWebhookSpecFilterGroupFilter struct {
	Type                  string `json:"type"`
	ExcludeMatchedPattern bool   `json:"exclude_matched_pattern"`
	Pattern               string `json:"pattern"`
}

type AwsCodebuildWebhookSpecFilterGroup struct {
	Filter []AwsCodebuildWebhookSpecFilterGroup `json:"filter"`
}

type AwsCodebuildWebhookSpec struct {
	Secret       string                    `json:"secret"`
	Url          string                    `json:"url"`
	ProjectName  string                    `json:"project_name"`
	BranchFilter string                    `json:"branch_filter"`
	FilterGroup  []AwsCodebuildWebhookSpec `json:"filter_group"`
	PayloadUrl   string                    `json:"payload_url"`
}



type AwsCodebuildWebhookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCodebuildWebhookList is a list of AwsCodebuildWebhooks
type AwsCodebuildWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodebuildWebhook CRD objects
	Items []AwsCodebuildWebhook `json:"items,omitempty"`
}