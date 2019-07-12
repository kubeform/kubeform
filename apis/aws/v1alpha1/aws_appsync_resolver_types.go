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

type AwsAppsyncResolver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncResolverSpec   `json:"spec,omitempty"`
	Status            AwsAppsyncResolverStatus `json:"status,omitempty"`
}

type AwsAppsyncResolverSpecPipelineConfig struct {
	Functions []string `json:"functions"`
}

type AwsAppsyncResolverSpec struct {
	Type             string                   `json:"type"`
	DataSource       string                   `json:"data_source"`
	ResponseTemplate string                   `json:"response_template"`
	ApiId            string                   `json:"api_id"`
	Field            string                   `json:"field"`
	RequestTemplate  string                   `json:"request_template"`
	Kind             string                   `json:"kind"`
	PipelineConfig   []AwsAppsyncResolverSpec `json:"pipeline_config"`
	Arn              string                   `json:"arn"`
}

type AwsAppsyncResolverStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppsyncResolverList is a list of AwsAppsyncResolvers
type AwsAppsyncResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncResolver CRD objects
	Items []AwsAppsyncResolver `json:"items,omitempty"`
}
