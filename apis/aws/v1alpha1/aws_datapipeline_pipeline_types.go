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

type AwsDatapipelinePipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDatapipelinePipelineSpec   `json:"spec,omitempty"`
	Status            AwsDatapipelinePipelineStatus `json:"status,omitempty"`
}

type AwsDatapipelinePipelineSpec struct {
	Description string            `json:"description"`
	Tags        map[string]string `json:"tags"`
	Name        string            `json:"name"`
}



type AwsDatapipelinePipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDatapipelinePipelineList is a list of AwsDatapipelinePipelines
type AwsDatapipelinePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDatapipelinePipeline CRD objects
	Items []AwsDatapipelinePipeline `json:"items,omitempty"`
}