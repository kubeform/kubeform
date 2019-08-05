package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppsyncResolver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncResolverSpec   `json:"spec,omitempty"`
	Status            AppsyncResolverStatus `json:"status,omitempty"`
}

type AppsyncResolverSpecPipelineConfig struct {
	// +optional
	Functions []string `json:"functions,omitempty" tf:"functions,omitempty"`
}

type AppsyncResolverSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApiID string `json:"apiID" tf:"api_id"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	DataSource string `json:"dataSource,omitempty" tf:"data_source,omitempty"`
	Field      string `json:"field" tf:"field"`
	// +optional
	Kind string `json:"kind,omitempty" tf:"kind,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PipelineConfig   []AppsyncResolverSpecPipelineConfig `json:"pipelineConfig,omitempty" tf:"pipeline_config,omitempty"`
	RequestTemplate  string                              `json:"requestTemplate" tf:"request_template"`
	ResponseTemplate string                              `json:"responseTemplate" tf:"response_template"`
	Type             string                              `json:"type" tf:"type"`
}

type AppsyncResolverStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppsyncResolverSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncResolverList is a list of AppsyncResolvers
type AppsyncResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncResolver CRD objects
	Items []AppsyncResolver `json:"items,omitempty"`
}
