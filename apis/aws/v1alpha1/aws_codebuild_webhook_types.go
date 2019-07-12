package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodebuildWebhookSpec   `json:"spec,omitempty"`
	Status            AwsCodebuildWebhookStatus `json:"status,omitempty"`
}

type AwsCodebuildWebhookSpec struct {
	ProjectName  string `json:"project_name"`
	BranchFilter string `json:"branch_filter"`
	PayloadUrl   string `json:"payload_url"`
	Secret       string `json:"secret"`
	Url          string `json:"url"`
}



type AwsCodebuildWebhookStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildWebhookList is a list of AwsCodebuildWebhooks
type AwsCodebuildWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodebuildWebhook CRD objects
	Items []AwsCodebuildWebhook `json:"items,omitempty"`
}