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

type AppsyncResolver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncResolverSpec   `json:"spec,omitempty"`
	Status            AppsyncResolverStatus `json:"status,omitempty"`
}

type AppsyncResolverSpecPipelineConfig struct {
	// +optional
	Functions []string `json:"functions,omitempty"`
}

type AppsyncResolverSpec struct {
	ApiId string `json:"api_id"`
	// +optional
	DataSource string `json:"data_source,omitempty"`
	Field      string `json:"field"`
	// +optional
	Kind string `json:"kind,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PipelineConfig   *[]AppsyncResolverSpec `json:"pipeline_config,omitempty"`
	RequestTemplate  string                 `json:"request_template"`
	ResponseTemplate string                 `json:"response_template"`
	Type             string                 `json:"type"`
}

type AppsyncResolverStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
